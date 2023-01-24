package internal

import (
	"errors"
	"net"
	"strings"
)

func PrivateIpCheck(ip string) (bool, error) {
	var IP net.IP
	var err error
	if strings.Contains(ip, "/") {
		IP, _, err = net.ParseCIDR(ip)
	} else {
		IP = net.ParseIP(ip)
	}
	if err != nil {
		return false, err
	}
	if IP == nil {
		return false, errors.New("invalid ip")
	}
	_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
	_, private8BitBlock, _ := net.ParseCIDR("127.0.0.0/8")
	_, privateIpv6LoBlock, err := net.ParseCIDR("::1/128")
	if err != nil {
		panic(err)
	}

	switch true {
	case private24BitBlock.Contains(IP):
		return true, nil
	case private20BitBlock.Contains(IP):
		return true, nil
	case private16BitBlock.Contains(IP):
		return true, nil
	case private8BitBlock.Contains(IP):
		return true, nil
	case privateIpv6LoBlock.Contains(IP):
		return true, nil
	}
	return false, nil
}
