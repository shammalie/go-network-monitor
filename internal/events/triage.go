package events

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/shammalie/go-network-monitor/internal"
	"github.com/shammalie/go-network-monitor/internal/state"
	network_capture_v1 "github.com/shammalie/go-network-monitor/pkg/network_capture.v1"
	"github.com/shammalie/network_ip_bearer/pkg/ipapi"
)

type TriageService struct {
	ipQueryQueue *state.Fifo
	filter       *state.Prioritylist
	warmCache    *state.Cache
	db           *internal.Db
}

func NewTriageService(filterList []string) *TriageService {
	var wg sync.WaitGroup
	wg.Add(1)
	t := &TriageService{
		ipQueryQueue: state.NewFifoQueue(),
		filter:       state.NewPrioritylist(filterList),
		warmCache:    state.NewLocalCache(10 * time.Second),
		db:           internal.NewMongoClient(),
	}
	go func() {
		defer wg.Done()
		t.queueProcessor()
	}()

	return t
}

func (t *TriageService) Triage(event *network_capture_v1.NetworkCaptureRequest) {
	sourceIP := event.NetworkLayer.SrcIp
	if found, err := internal.PrivateIpCheck(sourceIP); err != nil || found {
		fmt.Printf("dropping private ip %s\n", sourceIP)
		return
	}
	if t.filter.Lookup(sourceIP) {
		return
	}
	if _, found := t.warmCache.Get(sourceIP); found {
		return
	}
	if _, err := t.db.GetIpDataByIp(sourceIP); err == nil {
		return
	}
	t.addIpToQueue(sourceIP)
}

func (t *TriageService) addIpToQueue(ip string) {
	t.warmCache.Set(ip)
	t.ipQueryQueue.Enqueue(ip)
}

func (t *TriageService) queueProcessor() {
	for range time.NewTicker(100 * time.Millisecond).C {
		if t.ipQueryQueue.GetLen() == 0 {
			continue
		}
		ip := t.ipQueryQueue.Dequeue()
		ipDetail := ipapi.RetryRequest(ip)
		log.Printf("processed ip: %s\n", ip)
		err := t.db.InsertIpDetail(ipDetail)
		if err != nil {
			fmt.Println(err)
		}
		t.warmCache.Remove(ip)
	}
}
