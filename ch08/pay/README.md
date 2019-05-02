# user Service

go.micro.srv.pay

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./pay --registry=etcdv3   //服务注册:etcdv3
```

Run the client
```
go run cli/main.go
```