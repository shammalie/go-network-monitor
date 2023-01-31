package internal

import (
	"fmt"
	"sync"
	"time"

	network_capture_v1 "github.com/shammalie/go-network-monitor/pkg/network_capture.v1"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventProcessor struct {
	Events      chan *network_capture_v1.NetworkCaptureRequest
	saveEvents  bool
	db          *Db
	warmCache   *Cache
	ipProcessor *IpProcessor
}

type Event struct {
	Id                           primitive.ObjectID `bson:"_id" json:"_id"`
	Ip_id                        primitive.ObjectID `bson:"ip_id" json:"ip_id"`
	NetworkLayerSourceIp         string             `bson:"networkLayerSourceIp" json:"networkLayerSourceIp"`
	NetworkLayerDestinationIp    string             `bson:"NetworkLayerDestinationIp" json:"etworkLayerDestinationIp"`
	NetworkLayerProtocol         string             `bson:"networkLayerProtocol" json:"networkLayerProtocol"`
	TransportLayerSourceIp       string             `bson:"transportLayerSourceIp" json:"transportLayerSourceIp"`
	TransportLayerDestinationIp  string             `bson:"transportLayerDestinationIp" json:"transportLayeDestinationIp"`
	TransportLayerProtocol       string             `bson:"transportLayerProtocol" json:"transportLayerProtocol"`
	ApplicationLayerProtocol     string             `bson:"applicationLayerProtocol" json:"applicationLayerProtocol"`
	ApplicationLayerPayload      []byte             `bson:"applicationLayerPayload" json:"applicationLayerPayload"`
	MetadataCaptureLength        int64              `bson:"metadataCaptureLength" json:"metadataCaptureLength"`
	MetadataOriginalPacketLength int64              `bson:"metadataOriginalPacketLength" json:"metadataOriginalPacketLength"`
	MetadataTimestamp            int64              `bson:"metadataTimestamp" json:"metadataTimestamp"`
	MetadataTruncated            bool               `bson:"metadataTruncated" json:"metadataTruncated"`
}

func NewEventProcessor(db *Db) *EventProcessor {
	var wg sync.WaitGroup
	recEvents := viper.GetBool("RECORD_IP_EVENTS")
	processor := &EventProcessor{
		Events:      make(chan *network_capture_v1.NetworkCaptureRequest),
		saveEvents:  recEvents,
		db:          db,
		warmCache:   NewLocalCache(5 * time.Second),
		ipProcessor: NewIpProcessor(),
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range processor.Events {
			processor.handleEvent(event)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ipDetail := range processor.ipProcessor.outgoing {
			err := db.InsertIpDetail(ipDetail)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("processed %s\n", ipDetail.Ip)
		}
	}()
	fmt.Println("started event processor")
	return processor
}

func (p *EventProcessor) handleEvent(event *network_capture_v1.NetworkCaptureRequest) {
	if p.ipProcessor.isRateLimited() {
		fmt.Println("processor currently rate limited, not accepting new requests")
		return
	}
	processedEvent := p.convertEvent(event)
	srcIp := processedEvent.NetworkLayerSourceIp
	isPrivate, err := PrivateIpCheck(srcIp)
	if err != nil {
		panic(err)
	}
	if isPrivate {
		return
	}
	var existingId primitive.ObjectID
	cacheElement, found := p.warmCache.Get(srcIp)
	if !found {
		d, err := p.db.GetIpDataByIp(srcIp)
		if err != nil {
			id := primitive.NewObjectID()
			go func() {
				p.ipProcessor.incoming <- IpDetail{
					Id:        id,
					Ip:        srcIp,
					FirstSeen: time.Now().UTC().UnixMilli(),
				}
			}()
			p.warmCache.Set(Element{
				id: id,
				ip: srcIp,
			})
			existingId = id
		} else {
			existingId = d.Id
		}
	} else {
		existingId = cacheElement.id
	}
	convEvent := p.convertEvent(event)
	convEvent.Ip_id = existingId
	p.db.InsertIpEvent(convEvent)
}

func (p *EventProcessor) convertEvent(e *network_capture_v1.NetworkCaptureRequest) *Event {
	event := &Event{}
	var err error
	err = event.handleLayer(*e.NetworkLayer)
	if err != nil {
		fmt.Println(err)
	}
	err = event.handleLayer(*e.TransportLayer)
	if err != nil {
		fmt.Println(err)

	}
	err = event.handleLayer(*e.ApplicationLayer)
	if err != nil {
		fmt.Println(err)
	}
	err = event.handleLayer(*e.Metadata)
	if err != nil {
		fmt.Println(err)
	}
	return event
}

func (e *Event) handleLayer(layer interface{}) error {
	switch layer := layer.(type) {
	case network_capture_v1.ApplicationLayer:
		if layer.Protocol != nil {
			e.ApplicationLayerProtocol = layer.Protocol.Name
		}
		e.ApplicationLayerPayload = layer.Payload
	case network_capture_v1.NetworkLayer:
		if layer.Protocol != nil {
			e.NetworkLayerProtocol = layer.Protocol.Name
		}
		e.NetworkLayerSourceIp = layer.SrcIp
		e.NetworkLayerDestinationIp = layer.DstIp
	case network_capture_v1.TransportLayer:
		if layer.Protocol != nil {
			e.TransportLayerProtocol = layer.Protocol.Name
		}
		e.TransportLayerSourceIp = layer.SrcPort
		e.TransportLayerDestinationIp = layer.DstPort
	case network_capture_v1.Metadata:
		e.MetadataCaptureLength = layer.OriginalPacketLength
		e.MetadataOriginalPacketLength = layer.OriginalPacketLength
		e.MetadataTimestamp = layer.Timestamp
		e.MetadataTruncated = layer.Truncated
	default:
		return fmt.Errorf("provided layer is not of the supported types")
	}
	return nil
}
