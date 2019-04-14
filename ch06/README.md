# Ch06 Service(开发与调试)

本地开发，用到别的服务，但是别的服务在本地环境又没有部署，则可以指定ip+端口或者mock

1.启动服务时，指定端口。
```
eg:go run srv/main.go --server_address=127.0.0.1:8089
```

或者直接去etcd里面查看服务的ip以及端口，具体使用说明见https://github.com/1819997197/micro/tree/master/ch00

2.调用服务时，指明服务的ip+端口。
```
eg:c.Call(context.Background(), req, nil, WithAddress(fmt.Sprintf("%s:%d", address, port)))
```


## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./ch06 --server_address=127.0.0.1:8089 --registry=etcdv3
```

Run the client
```
go run cli/main.go
```

## 参考资料
```
//gomock
https://github.com/golang/mock
```