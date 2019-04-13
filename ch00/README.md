#### etcd集群的搭建与使用
1.下载(https://github.com/coreos/etcd/releases)

2.解压缩，将两个bin文件etcd、etcdctl添加到系统环境/usr/local/bin中

3.单机启动
```
cd /usr/local/bin
./ etcd
```

4.etcdctl客户端：默认版本2，V2版本基本操作
```
./etcdctl --version //查看api的版本是v2
./etcdctl set key value  //创建一个节点，并给他一个value
./etcdctl get key   //获取一个节点的value
```

5.etcdctl客户端，改用版本3：
```
export ETCDCTL_API=3
./etcdctl version //查看版本(v3)
./etcdctl put key value //添加/更新

```
6.查看所有的key-value
```
//v2
./etcdctl ls -p /  //列出目录下所有目录或者节点
curl -s http://127.0.0.1:2479/v2/keys/?recursive=true

//v3
ETCDCTL_API=3 ./etcdctl --prefix --keys-only=false  get /
ETCDCTL_API=3 ./etcdctl --prefix --keys-only=false  get /p //查看拥有/p前缀的keys-values
```

7.etcd 中，key 可以有 TTL 属性，超过这个时间会被自动删除

8.etcd 除了提供命令行工具之外，还对外提供 HTTP API服务。方便测试，也方便集成到各种语言中

9.集群搭建
```
//由于条件受限，用三个端口2380,2381,2382来模拟集群(这三个是成员之间通信),2379,2389,2399是给客户端连接的。
//带advertise参数是广播参数: 如--listen-client-urls和--advertise-client-urls, 前者是Etcd端监听客户端的url,后者是Etcd客户端请求的url, 两者端口是相同的, 只不过后者一般为公网IP, 暴露给外部使用
./etcd --name my-etcd-1  \
--listen-client-urls http://0.0.0.0:2379 \
--advertise-client-urls http://127.0.0.1:2379 \
--listen-peer-urls http://0.0.0.0:2380 \
--initial-advertise-peer-urls http://127.0.0.1:2380  \
--initial-cluster-token etcd-cluster-test \
--initial-cluster-state new \
--initial-cluster my-etcd-1=http://127.0.0.1:2380,my-etcd-2=http://127.0.0.1:2381,my-etcd-3=http://127.0.0.1:2382

./etcd --name my-etcd-2  \
--listen-client-urls http://0.0.0.0:2389 \
--advertise-client-urls http://127.0.0.1:2389 \
--listen-peer-urls http://0.0.0.0:2381 \
--initial-advertise-peer-urls http://127.0.0.1:2381  \
--initial-cluster-token etcd-cluster-test \
--initial-cluster-state new \
--initial-cluster my-etcd-1=http://127.0.0.1:2380,my-etcd-2=http://127.0.0.1:2381,my-etcd-3=http://127.0.0.1:2382

./etcd --name my-etcd-3
--listen-client-urls http://0.0.0.0:2399
--advertise-client-urls http://127.0.0.1:2399
--listen-peer-urls http://0.0.0.0:2382
--initial-advertise-peer-urls http://127.0.0.1:2382
--initial-cluster-token etcd-cluster-test
--initial-cluster-state new
--initial-cluster my-etcd-1=http://127.0.0.1:2380,my-etcd-2=http://127.0.0.1:2381,my-etcd-3=http://127.0.0.1:2382
```

10.查看集群成员
```
etcdctl member list
```

11.使用时需要指定endpoints(默认本地端口2379), 集群时数据会迅速同步:
```
ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2389 put /foo will
ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 get /foo
ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 --prefix --keys-only=false  get /  //查看:2379所有key-value
```

12.参考资料
```
http://www.iigrowing.cn/etcd_shi_yong_ru_men.html
https://blog.csdn.net/varyall/article/details/79128181
https://blog.csdn.net/u010511236/article/details/52386229
https://segmentfault.com/a/1190000008672912
```
