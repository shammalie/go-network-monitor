package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getIpInformation(obj IpDetail) (*IpDetail, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ipapi.co/%s/json/", obj.Ip), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.8")
	resp, err := client.Do(req)
	if err != nil && resp.StatusCode >= 400 {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &obj); err != nil {
		return nil, err
	}
	if obj.Error {
		return nil, fmt.Errorf(fmt.Sprintf("reason: %s, message:%s", obj.Reason, obj.Message))
	}
	return &obj, nil
}
