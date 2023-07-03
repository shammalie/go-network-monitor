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
	ipQueryQueue  *state.PriorityQueue
	filter        *state.Prioritylist
	warmCache     *state.Cache
	coolDownCache *state.Cache
	db            *internal.Db
}

func NewTriageService(filterList []string) *TriageService {
	var wg sync.WaitGroup
	wg.Add(1)
	t := &TriageService{
		ipQueryQueue:  &state.PriorityQueue{},
		filter:        state.NewPrioritylist(filterList),
		warmCache:     state.NewLocalCache(10 * time.Second),
		coolDownCache: state.NewLocalCache(1 * time.Minute),
		db:            internal.NewMongoClient(),
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
	if _, found := t.coolDownCache.Get(sourceIP); found {
		return
	}
	if _, found := t.warmCache.Get(sourceIP); found {
		return
	}
	if ipData, err := t.db.GetIpDataByIp(sourceIP); err == nil {
		timeLastMonth := time.Now().Add((-24 * time.Hour) * 30).UnixMilli()
		lastSeenEpoch := ipData.LastSeen
		// Consider firstSeen when lastSeen has not been set.
		if lastSeenEpoch == 0 {
			if ipData.FirstSeen < timeLastMonth {
				t.addIpToQueue(ipData.Ip, 1)
			}
		} else {
			if lastSeenEpoch < timeLastMonth {
				t.addIpToQueue(ipData.Ip, 1)
			}
		}
		t.coolDownCache.Set(sourceIP)
		return
	}
	// new ip's have higher priority for lookup.
	t.addIpToQueue(sourceIP, 10000)
}

func (t *TriageService) addIpToQueue(ip string, priority int) {
	t.warmCache.Set(ip)
	t.ipQueryQueue.PushString(ip, priority)
}

func (t *TriageService) queueProcessor() {
	for range time.NewTicker(100 * time.Millisecond).C {
		if t.ipQueryQueue.Len() == 0 {
			continue
		}
		item := t.ipQueryQueue.Pop().(*state.Item)
		ip := item.GetValue()
		ipDetail := ipapi.RetryRequest(ip)
		if ipDetail == nil {
			continue
		}
		log.Printf("processed ip: %s\n", ip)
		var err error
		if item.GetPriority() == 1 {
			err = t.db.UpdateIpDetail(ipDetail)
		} else {
			err = t.db.InsertIpDetail(ipDetail)
		}
		if err != nil {
			fmt.Println(err)
		}
		t.warmCache.Remove(ip)
	}
}
