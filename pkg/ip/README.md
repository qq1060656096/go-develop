## Ip包
> 提供快速获取ip地址和mac地址方法

```go
ip, err := ip.New()
if err != nil {
    panic(err)
}
// 获取本机所有ip地址
ip.IpAddrs()
// 获取本机ip地址
ip.IpAddr()
// 获取本机所有mac地址
ip.MacAddrs()
// 获取本机mac地址
ip.MacAddr()
```