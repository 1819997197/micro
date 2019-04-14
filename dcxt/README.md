# dcxt 点餐系统

- go-micro
- go-config
- gorm
- gin
- redis
- etcd

![Image text](https://github.com/1819997197/micro/tree/master/dcxt/dcxt.png)

gin作为bff层，主要用它的路由功能，对外提供http接口

go-micro作为一个成熟的微服务框架提供服务，外部不可访问

etcd注册中心: 微服务的注册与发现; 配置中心: 配置服务的公共配置信息


## Usage

Run order service
```
make build
./order  --registry=etcdv3
```

Run user service
```
make build
./user  --registry=etcdv3
```

Run api
```
make build
./api
```

Run
```
//curl或者postman
curl http://127.0.0.1:8080/
```


## 测试服务负载均衡
1.修改打印received_order的地方(区分三个服务)，运行三个order服务

2.运行user服务、api

3.不断运行 curl http://127.0.0.1:8080/

4.观察各个服务端接收请求数？
order1:
received_order_1:  1
received_order_1:  1

order2：
received_order_2:  1

order3：
received_order_3:  1

5.把服务端order3退出进程，观察各个服务请求接收数?
order1:
received_order_1:  1
received_order_1:  1
received_order_1:  1
received_order_1:  1
received_order_1:  1
received_order_1:  1

order2:
received_order_2:  1
received_order_2:  1
received_order_2:  1
received_order_2:  1
received_order_2:  1

6.服务可伸缩