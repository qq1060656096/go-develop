package ip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetIps 测试获取本机所有ip地址
func TestGetIps(t *testing.T) {
	ipAddrs, err := GetIps()
	assert.Equal(t, nil, err)
	assert.LessOrEqual(t, 1, len(ipAddrs))
}
// TestGetMacAddrs 测试获取本机所有Mac地址
func TestGetMacAddrs(t *testing.T) {
	macAddrs, err := GetMacAddrs()
	assert.Equal(t, nil, err)
	assert.LessOrEqual(t, 1, len(macAddrs))
}

func TestNew(t *testing.T) {
	ip, err := New()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, ip)
}
func TestIp_IpAddrs(t *testing.T) {
	ip, err := New()
	assert.Equal(t, nil, err)
	assert.LessOrEqual(t, 1, len(ip.IpAddrs()))
}

func TestIp_IpAddr(t *testing.T) {
	ip, err := New()
	assert.Equal(t, nil, err)
	assert.IsType(t, "127.0.0.1", ip.IpAddr())
	assert.LessOrEqual(t, 1, len(ip.IpAddr()))
}

func TestIp_MacAddrs(t *testing.T) {
	ip, err := New()
	assert.Equal(t, nil, err)
	assert.LessOrEqual(t, 1, len(ip.MacAddrs()))
}

func TestIp_MacAddr(t *testing.T) {
	ip, err := New()
	assert.Equal(t, nil, err)
	assert.IsType(t, "ac:de:48:00:11:22", ip.MacAddr())
	assert.LessOrEqual(t, 1, len(ip.MacAddr()))
}
