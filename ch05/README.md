# Ch05 Service(go-micro发布与订阅)

Go-micro 给事件驱动架构内置了消息代理（broker）接口。发布与订阅像RPC一样操控生成的protobuf消息。这些消息会自动编/解码并通过代理发送。

Go-micro默认包含点到点的http代理，但是也可以通过go-plugins把这层逻辑替换掉

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./ch05 --registry=etcdv3   //服务注册:etcdv3
```

Run the client
```
go run cli/main.go
```

## 参考资料
```
//go-micro发布与订阅
https://micro.mu/docs/cn/go-micro.html
//go-micro发布与订阅，通过插件还支持kafka、nats、rabbitmq、redis
https://github.com/micro/go-plugins/tree/master/broker
```