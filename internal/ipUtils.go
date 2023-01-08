package internal

import (
	"errors"
	"net"
	"strings"
)

func IsPrivateIP(ip string) (bool, error) {
	var err error
	private := false
	var IP net.IP
	if strings.Contains(ip, "/") {
		IP, _, err = net.ParseCIDR(ip)
	} else {
		IP = net.ParseIP(ip)
	}
	if IP == nil {
		err = errors.New("invalid ip")
	} else {
		_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
		_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
		_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
		_, private8BitBlock, _ := net.ParseCIDR("127.0.0.0/8")
		_, privateIpv6LoBlock, err := net.ParseCIDR("::1/128")
		if err != nil {
			panic(err)
		}

		private = private24BitBlock.Contains(IP) || private20BitBlock.Contains(IP) || private16BitBlock.Contains(IP) || private8BitBlock.Contains(IP) || privateIpv6LoBlock.Contains(IP)
	}
	return private, err
}
