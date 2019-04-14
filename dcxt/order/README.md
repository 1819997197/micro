# order Service

- go-micro
- gorm
- go-config(热加载)

## 数据库
```
CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `price` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into orders(price) values(100);
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./order --registry=etcdv3   //服务注册:etcdv3
```

Run the client
```
go run cli/main.go
```

## 参考资料
```
https://micro.mu/docs/cn/go-config.html //go-config
http://gorm.io/zh_CN/docs/index.html //gorm
```