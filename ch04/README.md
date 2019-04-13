# Ch04 Service(go-micro+redis+go-config+etcd)

- go-micro
- go-config
- redis
- etcd (配置中心)

## etcd 配置redis连接参数

```
//set redisConfig
ETCDCTL_API=3 ./etcdctl put /micro/config/redis '{"address":"192.168.1.105","port":6379}'
//get
ETCDCTL_API=3 ./etcdctl --prefix --keys-only=false  get /
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./ch04 --registry=etcdv3   //服务注册:etcdv3
```

Run the client
```
go run cli/main.go
```

## 参考资料
```
//go-config读取etcd配置
https://github.com/micro/go-config/tree/master/source/etcd
//redis
https://blog.csdn.net/wangshubo1989/article/details/75050024
//go-config
https://micro.mu/docs/cn/go-config.html
```