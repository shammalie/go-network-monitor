package internal

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	requestCountLimit = 20
)

type IpProcessor struct {
	failureCount   int
	counter        float64
	StopProcessing chan bool
	db             *Db
	cache          *RedisClient
	events         chan ipEvent
}

type IpDetail struct {
	Id                 primitive.ObjectID `bson:"_id" json:"_id"`
	Ip                 string             `bson:"ip" json:"ip"`
	City               string             `bson:"city" json:"city"`
	Region             string             `bson:"region" json:"region"`
	RegionCode         string             `bson:"region_code" json:"region_code"`
	CountryCode        string             `bson:"country_code" json:"country_code"`
	CountryCodeIso3    string             `bson:"country_code_iso3" json:"country_code_iso3"`
	CountryName        string             `bson:"country_name" json:"country_name"`
	CountryCapital     string             `bson:"country_capital" json:"country_capital"`
	CountryTld         string             `bson:"country_tld" json:"country_tld"`
	ContinentCode      string             `bson:"continent_code" json:"continent_code"`
	InEu               bool               `bson:"in_eu" json:"in_eu"`
	Postal             string             `bson:"postal" json:"postal"`
	Latitude           float64            `bson:"latitude" json:"latitude"`
	Longitude          float64            `bson:"longitude" json:"longitude"`
	Timezone           string             `bson:"timezone" json:"timezone"`
	UtcOffset          string             `bson:"utc_offset" json:"utc_offset"`
	CountryCallingCode string             `bson:"country_calling_code" json:"country_calling_code"`
	Currency           string             `bson:"currency" json:"currency"`
	CurrencyName       string             `bson:"currency_name" json:"currency_name"`
	Languages          string             `bson:"languages" json:"languages"`
	Asn                string             `bson:"asn" json:"asn"`
	Org                string             `bson:"org" json:"org"`
	FirstSeen          int64              `bson:"first_seen" json:"first_seen"`
	Error              *bool
	Reason             *string
	Message            *string
}

func NewIpProcessor(db *Db) *IpProcessor {
	var wg sync.WaitGroup
	processor := &IpProcessor{
		failureCount:   0,
		counter:        1,
		StopProcessing: make(chan bool),
		db:             db,
		cache:          NewRedisClient(5),
		events:         make(chan ipEvent, 1000),
	}
	entries, err := processor.cache.loadCacheEntries()
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		processor.events <- e
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		processor.processorLoop()
	}()
	fmt.Println("started IP lookup processor")
	return processor
}

func (p *IpProcessor) processorLoop() {
	defer close(p.StopProcessing)
	for {
		select {
		case event := <-p.events:
			response := p.processRequest(event.Ip)
			if response == nil {
				p.cache.Del(event.Ip)
				return
			}
			if p.failureCount != 0 {
				p.failureCount = 0
			}
			timeInterval := math.Floor(float64(p.counter) / 2)
			if timeInterval <= 0 || timeInterval > 32 {
				p.counter = 1
			} else {
				p.counter = timeInterval
			}
			response.FirstSeen = event.Timestamp
			response.Id = event.Id
			err := p.db.InsertIpDetail(response)
			if err != nil {
				log.Fatal(err)
			}
			p.cache.Del(response.Ip)
			fmt.Printf("processed new ip %s\n", response.Ip)
		case <-p.StopProcessing:
			return
		}
	}
}

func (p *IpProcessor) processRequest(ip string) *IpDetail {
	for {
		t := time.NewTicker(time.Duration(p.counter) * time.Second)
		for range t.C {
			response, err := getIpInformation(ip)
			if err != nil {
				if strings.Contains(err.Error(), "Reserved IP Address") {
					return nil
				}
				p.failureCount += 1
				fmt.Printf("response error: %s, failure count %d, tick count: %f, queue count: %d, ip: %s\n",
					err,
					p.failureCount,
					p.counter,
					len(p.events),
					ip)
				if p.failureCount > requestCountLimit && p.counter != 3600 {
					fmt.Println("extending tick period")
					p.counter = 3600
				} else if p.counter < 32 {
					p.counter += p.counter
				}
				break
			}
			return response
		}
	}
}
