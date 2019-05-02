# ch08

- zipkin
- elasticsearch
- gin
- go-micro
- etcd

## 准备
> * 主机 192.168.1.104 centos7.X(root、will普通用户)
> * 主机 192.168.1.101 centos6.X
> * 104主机安装elasticsearch、elasticsearch-head
> * 101主机安装etcd、zipkin、go环境

## 安装JDK
```
yum install java-1.8.0-openjdk* -y
// 安装完成之后，查看安装是否成功
java -version
```

## 安装 zipkin
```
// 新建目录
mkdir -p /usr/local/zipkin && cd "$_"
// 下载zipkin
wget -O zipkin.jar 'https://search.maven.org/remote_content?g=io.zipkin.java&a=zipkin-server&v=LATEST&c=exec'
// 启动
java -jar zipkin.jar
```

## 安装 elasticsearch
```
// 选择一个下载目录
wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-6.2.4.tar.gz
// 解压
tar -zxvf elasticsearch-6.2.4.tar.gz
// 移动到/usr/local/elasticsearch-6.2.4目录
mv elasticsearch-6.2.4 /usr/local/
```

因为安全问题elasticsearch不让用root用户直接运行，需要新建用户es:es(chown -R es:es /usr/local/elasticsearch-6.2.4/)

默认情况下，Elasticsearch 只允许本机访问，如果需要远程访问，可以修改 Elasticsearch 安装目录中的config/elasticsearch.yml文件
```
// 去掉network.host的注释，将它的值改成0.0.0.0，让任何人都可以访问，然后重新启动 Elasticsearch
// "network.host:"和"0.0.0.0"中间有个空格，不能忽略，不然启动会报错。线上服务不要这样设置，要设成具体的 IP
51 # ---------------------------------- Network -----------------------------------
52 #
53 # Set the bind address to a specific IP (IPv4 or IPv6):
54 #
55 network.host: 0.0.0.0
56 #
```

## 启动
```
[will@will bin]$ /usr/local/elasticsearch-6.2.4/bin/elasticsearch
......
ERROR: [3] bootstrap checks failed
[1]: max file descriptors [4096] for elasticsearch process is too low, increase to at least [65536]
[2]: max number of threads [3871] for user [will] is too low, increase to at least [4096]
[3]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
[2019-05-01T22:17:12,344][INFO ][o.e.n.Node               ] [80krV8-] stopping ...
[2019-05-01T22:17:12,377][INFO ][o.e.n.Node               ] [80krV8-] stopped
[2019-05-01T22:17:12,377][INFO ][o.e.n.Node               ] [80krV8-] closing ...
[2019-05-01T22:17:12,396][INFO ][o.e.n.Node               ] [80krV8-] closed
```

启动遇到的问题:

1、针对错误[1]、[2]，可以采取如下方式:

修改/etc/security/limits.conf配置文件:
```
[root@will config]# vi /etc/security/limits.conf

// 添加如下配置项：
60 *       soft    nofile  65536
61 *       hard    nofile  65536
62 *       soft    nproc   4096
63 *       hard    nproc   4096
64
65 # End of file
```

修改/etc/security/limits.d/20-nproc.conf(centos6.x文件名不一样)配置文件:
```
vi /etc/security/limits.d/20-nproc.conf

// 修改如下配置:
4
5 *          soft    nproc     4096
6 root       soft    nproc     unlimited
```

修改完成后，重新登录elk账户，查看设置是否生效:
```
[will@will ~]$ ulimit -u
4096
[will@will ~]$ ulimit -n
65536
```

2、针对错误[3]，可以采取如下方式:
```
[root@will will]# vi /etc/sysctl.conf

// 增加如下配置:
10 # For more information, see sysctl.conf(5) and sysctl.d(5).
11 vm.max_map_count=262144

// 执行命令sysctl -p生效
[root@will will]# sysctl -p
vm.max_map_count = 262144
```

重新启动：
```
[will@will ~]$ /usr/local/elasticsearch-6.2.4/bin/elasticsearch
......
[2019-05-01T23:15:02,027][INFO ][o.e.n.Node               ] [80krV8-] started
[2019-05-01T23:15:02,041][INFO ][o.e.g.GatewayService     ] [80krV8-] recovered [0] indices into cluster_state
```

打开另一个终端进行测试:
```
[will@will ~]$ curl 'http://localhost:9200/?pretty'
{
  "name" : "80krV8-",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "-HDLdNecQBKZJ5UDHqqITA",
  "version" : {
    "number" : "6.2.4",
    "build_hash" : "ccec39f",
    "build_date" : "2018-04-12T20:37:28.497551Z",
    "build_snapshot" : false,
    "lucene_version" : "7.2.1",
    "minimum_wire_compatibility_version" : "5.6.0",
    "minimum_index_compatibility_version" : "5.0.0"
  },
  "tagline" : "You Know, for Search"
}
```

## elasticsearch-head插件(略)
```
// 参考文档
https://segmentfault.com/a/1190000014347757
```

## 启动zipkin(storage:elasticsearch)
```
STORAGE_TYPE=elasticsearch ES_HOSTS={es安装服务器IP}:9200 java -jar zipkin.jar
```

**修改配置，开放相关端口**



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

Run payservice
```
make build
./pay  --registry=etcdv3
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

Zipkin
```
http://127.0.0.1:9411/zipkin
```

运行结果:


查看es:

