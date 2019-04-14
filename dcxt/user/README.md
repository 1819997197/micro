# user Service

go.micro.srv.user

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user --registry=etcdv3   //服务注册:etcdv3
```

Run the client
```
go run cli/main.go
```