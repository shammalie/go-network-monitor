package internal

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	requestCountLimit = 8
	openStatus        = "OPEN"
	rateLimitedStatus = "RATE_LIMITED"
)

var rateLimitTimeMs = 3600000

type IpProcessor struct {
	status   string
	incoming chan IpDetail
	outgoing chan *IpDetail
	mu       sync.Mutex
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

func NewIpProcessor() *IpProcessor {
	var wg sync.WaitGroup
	processor := &IpProcessor{
		status:   "OPEN",
		incoming: make(chan IpDetail),
		outgoing: make(chan *IpDetail),
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ip := range processor.incoming {
			processor.ipProcessor(ip)
		}
	}()
	return processor
}

func (p *IpProcessor) isRateLimited() bool {
	defer p.mu.Unlock()
	p.mu.Lock()
	return p.status == rateLimitedStatus
}

func (p *IpProcessor) setStatus(status string) error {
	if strings.EqualFold(status, openStatus) || strings.EqualFold(status, rateLimitedStatus) {
		return fmt.Errorf("provided state is not valid")
	}
	defer p.mu.Unlock()
	p.mu.Lock()
	p.status = status
	return nil
}

func (p *IpProcessor) ipProcessor(event IpDetail) {
	var failCount int
	counterMs := 500
	for {
		t := time.NewTicker(time.Duration(counterMs) * time.Millisecond)
		for range t.C {
			response, err := getIpInformation(event)
			if err != nil {
				if strings.Contains(err.Error(), "Reserved IP Address") {
					return
				}
				failCount += 1
				fmt.Printf("response error: %s, failure count %d, tick count ms: %d, ip: %s\n",
					err,
					failCount,
					counterMs,
					event.Ip)
				if !p.isRateLimited() {
					p.setStatus(rateLimitedStatus)
				}
				if failCount >= requestCountLimit && counterMs != rateLimitTimeMs {
					fmt.Println("extending tick period")
					counterMs = rateLimitTimeMs
				} else if failCount < requestCountLimit {
					counterMs += counterMs
				}
				break
			}
			if p.isRateLimited() {
				p.setStatus(openStatus)
			}
			p.outgoing <- response
			return
		}
		p.mu.Unlock()
	}
}
