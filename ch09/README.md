# ch09

- gin
- kong

## 准备工作
> * 两台主机 192.168.1.101、192.168.1.104
> * 104主机安装kong(详情见kong安装.md)
> * 101主机对外提供服务(api)
> * 域名配置192.168.1.104主机，后端服务走内网
> * 统一api入口，所有流量经192.168.1.104主机(kong可配置集群)

启动kong以及api服务:
```
// 启动 postgresql 服务
systemctl start postgresql-9.5.service
// 启动kong
kong start -c /etc/kong/kong.conf
// 启动api
cd api
go build -o api
./api
```

原始URL: http://192.168.1.101:8080/person?id=1

## kong

 kong v0.13.x之前的版本是通过这个接口来管理用户接入的API，但是v0.13.x版本之后，官方不建议使用API来管理用户接口，而是用Service和Route模块来替代，管理的更精细。

## API模块管理
KONG API模块管理的是接入kong的上游API，每个接入的api必须至少指定hosts/uris/methods其中一个参数，kong将会代理所有指定upstream url的请求。

### 新接入一个api
```
[root@will kong]# curl -i -X POST \
>   --url  http://localhost:8001/apis/ \
>   --data 'name=person-api' \
>   --data 'hosts=192.168.1.101' \
>   --data 'uris=/person' \
>   --data 'retries=3' \
>   --data 'upstream_connect_timeout=60000' \
>   --data 'upstream_url=http://192.168.1.101:8080/person'
HTTP/1.1 201 Created
Date: Wed, 01 May 2019 03:29:45 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.13.1

{"created_at":1556710185813,"strip_uri":true,"id":"62242892-ed31-4dd8-827a-d49e44312cb1","hosts":["192.168.1.101"],"name":"person-api","http_if_terminated":false,"preserve_host":false,"upstream_url":"http:\/\/192.168.1.101:8080\/person","uris":["\/person"],"upstream_connect_timeout":60000,"upstream_send_timeout":60000,"upstream_read_timeout":60000,"retries":3,"https_only":false}

```

大致意思就是:这个API注册的名字叫person-api。它被挂载在网关的/person路径下，上游转发到http://localhost:8000去处理，转发的时候把前面的/person前缀给去掉。

name/id: 接入API到唯一标识符

新增好API后则可以通过kong代理来访问代理的服务
```
# 原接口
curl http://192.168.1.101:8080/person?id=1

#通过kong代理访问
curl -i -X GET \
  --header 'host:192.168.1.101' \
  --url  http://localhost:8000/person?id=1
//如果后端服务挂掉，kong会返回502 Bad Gateway
```

字段解析：

![Image text](https://github.com/1819997197/micro/blob/master/ch07/ch07.png)

### 根据name或id获取一个API
```
curl -i -X GET \
  --url  http://localhost:8001/apis/{name}
或
  curl -i -X GET \
  --url  http://localhost:8001/apis/{id}
```

### 查询所有接入到API列表
```
curl -i -X GET   --url  http://localhost:8001/apis/
```

### 根据name或id更新一个API
```
curl -i -X PATCH \
  --url  http://localhost:8001/apis/person-api \
  --data 'name=person-api-1' \
  --data 'retries=6'
或
  curl -i -X PATCH \
  --url  'http://localhost:8001/apis/62242892-ed31-4dd8-827a-d49e44312cb1' \
  --data 'name=person-api-1' \
  --data 'retries=6'
```

### 更新或新增一个API
```
#更新一个存在的API
curl -i -X PUT \
  --url  http://localhost:8001/apis/ \
  --data 'id=b93fcbe7-5dba-4888-bf8c-f4c8f798b53a' \
  --data 'hosts=192.168.1.101' \
  --data 'uris=/person' \
  --data 'upstream_url=http://192.168.1.101:8080/person'

#更新一个不存在的API
curl -i -X PUT \
  --url  http://localhost:8001/apis/ \
  --data 'name=test' \
  --data 'hosts=192.168.1.101' \
  --data 'uris=/person' \
  --data 'upstream_url=http://192.168.1.101:8080/person'

#新增
  curl -i -X PUT \
  --url  http://localhost:8001/apis/ \
  --data 'hosts=192.168.1.101' \
  --data 'upstream_url=http://192.168.1.101:8080/person'
```
如果request body中包含已有的API的主键（name or id），则会根据name or id 执行更新操作，同PATCH /apis/{name or id}；

如果指定了name or id但是没有查询到该记录，则返回404 NOT FOUND；

如果没有指定主键，则会新增一个API，同POST /apis/。

### 根据name或id删除一个API
```
curl -i -X DELETE \
  --url  http://localhost:8001/apis/{name}
或
  curl -i -X DELETE \
  --url  http://localhost:8001/apis/{id}
```

## Service和Route模块管理
(略)具体详情查看官网或者https://www.cnblogs.com/zhoujie/tag/kong/