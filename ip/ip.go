package ip

import (
	"encoding/binary"
	"net"
)

func GetCIDRIpList(cidr string) ([]net.IP, error) {
	var (
		index  = 0
		ipList []net.IP
	)
	_, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)
	finish := (start & mask) | (mask ^ 0xffffffff)
	ipList = make([]net.IP, finish-start+1)
	for i := start; i <= finish; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		ipList[index] = ip
		index++
	}
	return ipList, nil
}

func CheckIsKeepNetwork(ip string) bool {
	var ipNetList = []string{
		"0.0.0.0/8",
		"10.0.0.0/8",
		"100.64.0.0/10",
		"127.0.0.0/8",
		"169.254.0.0/16",
		"172.16.0.0/12",
		"192.0.0.0/24",
		"192.0.2.0/24",
		"192.88.99.0/24",
		"192.168.0.0/16",
		"198.18.0.0/15",
		"198.51.100.0/24",
		"203.0.113.0/24",
		"224.0.0.0/4",
		"240.0.0.0/4",
		"255.255.255.255/32"}
	if ip == "" {
		return true
	}
	for _, item := range ipNetList {
		_, ipNet, _ := net.ParseCIDR(item)
		if ipNet.Contains(net.ParseIP(ip)) {
			return true
		}
	}
	return false
}
