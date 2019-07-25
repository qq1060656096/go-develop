# go-develop
go 开发


### 目录结构
```php
go-develop  go开发目录
├─database      数据库目录
│  ├─mysql          MySQL
│  │  ├─indexes         MySQL索引
│  │  │  ├─README.md        MySQL索引理解
├─pkg           数据库目录
│  ├─ip             ip地址库(用于获取本机ip和mac地址)
├─rfc           rfc目录
│  ├─4122           rfc-4122: uuid生成
│  │  ├─uuid            uuid生成
├─tools         工具目录
├─.gitignore    .git忽略文件
├─LICENSE.txt   授权协议文件
├─README.md     README文件
```

### go常用库(pkg)
* **已定稿** [ip包](https://github.com/qq1060656096/go-develop/tree/master/pkg/ip)
> 获取本机所有ip和mac地址


### golang开发的cli工具
* **已定稿** [批量序列生成器](https://github.com/qq1060656096/batch-generate-sequence)
> 根据模板批量生成序列支持csv和excel
