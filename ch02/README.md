# Ch02 Service

This is the Ch02 service

Generated with

```
micro new micro/ch02 --namespace=go.micro --type=srv --plugin=registry=etcdv3
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.ch02
- Type: srv
- Alias: ch02

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./ch02-srv  //默认情况下，服务注册:mdns
./ch02-srv --registry=etcdv3   //服务注册:etcdv3
./ch02-srv --registry=consul   //服务注册:consul
```

Build a docker image
```
make docker
```

## micro new [service]深入了解
```
https://micro.mu/docs/cn/new.html
```
