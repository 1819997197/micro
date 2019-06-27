# ch10 ELK日志架构

------

## 服务器环境
| IP地址            | 主机名    |  角色                | 所属集群      |
| --------          | -----:    | :----:               | :----:        |
| 192.168.1.101     | filebeat1 |  业务服务器+filebeat |业务服务器集群 |
| ...               | ...       |  业务服务器+filebeat |业务服务器集群 |
| 192.168.1.110     | filebeatN |  业务服务器+filebeat |业务服务器集群 |
| 192.168.1.111     | kafkazk1  |  Kafka+ZooKeeper     |Kafka Broker集群 |
| 192.168.1.112     | kafkazk2  |  Kafka+ZooKeeper     |Kafka Broker集群 |
| 192.168.1.113     | kafkazk3  |  Kafka+ZooKeeper     |Kafka Broker集群 |
| 192.168.1.120     | Logstash1 |  Logstash            |Logstash集群  |
| 192.168.1.121     | Logstash2 |  Logstash            |Logstash集群  |
| 192.168.1.122     | Logstash3 |  Logstash            |Logstash集群  |
| 192.168.1.130     | ES1 |  ES Master、ES DataNode  |ES集群  |
| 192.168.1.131     | ES2 |  ES Master、ES DataNode  |ES集群  |
| 192.168.1.132     | ES3 |  ES Master、ES DataNode  |ES集群  |
| 192.168.1.133     | ES4 |  ES Master、Kibana  |ES集群  |

## ELK架构图

![ELK架构图](https://github.com/1819997197/micro/blob/master/ch11/picture/lxc.jpg)

注：
filebeat部署于各个业务节点上，主要用来收集原始数据；

filebeat将数据传输给消息队列(常见的有kafka、redis等)；

logstash拉取消息队列数据，过滤并分析数据，然后将格式化的数据传递给elasticsearch进行存储;

最后，由Kibana将日志和数据呈现给用户。

具体数据流向见下图：

![ELK数据流向图](https://github.com/1819997197/micro/blob/master/ch11/picture/lxc.jpg)