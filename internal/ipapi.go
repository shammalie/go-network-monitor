package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIpInformation(ip string) (*IpDetail, error) {
	var obj *IpDetail
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ipapi.co/%s/json/", ip), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.19.4")
	resp, err := client.Do(req)
	if err != nil && resp.StatusCode >= 400 {
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
	return obj, nil
}
