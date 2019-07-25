package ip

import (
	"net"
)

const (
	// IpLoopback 本地回环ip
	IpLoopback = "127.0.0.1"
)

type Ip struct {
	ipAddrs []string
	macAddrs []string
}

// New 创建Ip实例
func New() (ip *Ip, err error) {
	ip = &Ip{}
	ip.ipAddrs, err = GetIps()
	if err != nil {
		return nil, err
	}
	ip.macAddrs, err = GetMacAddrs()
	if err != nil {
		return nil, err
	}
	return
}

// IpAddr 获取本机所有Ip地址
func (ip *Ip) IpAddrs() []string {
	return ip.ipAddrs
}

// IpAddr 获取本机Ip地址
func (ip *Ip) IpAddr() string {
	l := len(ip.ipAddrs)
	if l > 1 {
		return ip.ipAddrs[1]
	}
	if l > 0 {
		return ip.ipAddrs[0]
	}
	return ""
}


// MacAddrs 获取本机所有mac地址
func (ip *Ip) MacAddrs() []string {
	return ip.macAddrs
}

// GetMacAddr 获取本机mac地址
func (ip *Ip) MacAddr() string {
	l := len(ip.macAddrs)
	if l > 1 {
		return ip.macAddrs[1]
	}
	if l > 0 {
		return ip.macAddrs[0]
	}
	return ""
}

// 获取本机所有ip
func GetIps() (ips[]string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := addr.(*net.IPNet); ok {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return
}

// GetMacAddrs 获取本机所有mac地址
func GetMacAddrs() (macAddrs []string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, inter := range interfaces {
		if s := inter.HardwareAddr.String(); s != "" {
			macAddrs = append(macAddrs, s)
		}
	}
	return
}