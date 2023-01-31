package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func verifyIPEvent(event IpDetail) error {
	if event.Ip == "" {
		return fmt.Errorf("ip is empty")
	}
	if event.Id.IsZero() {
		return fmt.Errorf("primitive id is zero")
	}
	if event.FirstSeen == 0 {
		return fmt.Errorf("no timestamp provided for event")
	}
	return nil
}

func getIpInformation(event IpDetail) (*IpDetail, error) {
	if err := verifyIPEvent(event); err != nil {
		return nil, err
	}
	var obj *IpDetail
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ipapi.co/%s/json/", event.Ip), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.19.4")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &obj); err != nil {
		return nil, err
	}
	if obj.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf("reason: %s, message:%s", *obj.Reason, *obj.Message))
	}
	obj.Id = event.Id
	obj.FirstSeen = event.FirstSeen
	return obj, nil
}
