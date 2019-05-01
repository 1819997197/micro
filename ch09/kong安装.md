# postgresql 安装
Kong 默认使用 postgresql 作为数据库

## 安装 Postgresql
```
// 添加 rpm
sudo yum install -y https://download.postgresql.org/pub/repos/yum/9.5/redhat/rhel-7-x86_64/pgdg-centos95-9.5-2.noarch.rpm
// 安装 postgresql 9.5
sudo  yum install -y postgresql95 postgresql95-server postgresql95-contrib
// 初始化数据库
sudo /usr/pgsql-9.5/bin/postgresql95-setup initdb

// 设置成 centos7 开机启动服务
sudo systemctl enable postgresql-9.5.service
// 启动 postgresql 服务
sudo systemctl start postgresql-9.5.service
// 查看 postgresql 状态
suso systemctl status postgresql-9.5.service
```

## 配置 Postgresql
执行完初始化任务之后，postgresql 会自动创建和生成两个用户和一个数据库：
> * linux 系统用户 postgres: 管理数据库的系统用户;
> * postgresql 用户 postgres: 数据库超级管理员;
> * 数据库 postgres: 用户 postgres 的默认数据库;
> * 密码由于是默认生成的，需要在系统中修改一下;
```
// 修改密码
sudo passwd postgres
```

为了安全以及满足 Kong 初始化的需求，需要在建立一个 postgre 用户 kong 和对应的 linux 用户 kong，并新建数据库 kong
```
// 新建 linux kong 用户
sudo adduser kong

// 使用管理员账号登录 psql 创建用户和数据库
// 切换 postgres 用户
// 切换 postgres 用户后，提示符变成 `-bash-4.2$`
su postgres

// 进入 psql 控制台
psql

// 此时会进入到控制台（系统提示符变为'postgres=#'）
// 先为管理员用户postgres修改密码
\password postgres

// 建立新的数据库用户（和之前建立的系统用户要重名）
create user kong with password '123456';

// 为新用户建立数据库
create database kong owner kong;

// 把新建的数据库权限赋予 kong
grant all privileges on database kong to kong;

// 退出控制台
\q
```
**在 psql 控制台下执行命令，一定记得在命令后添加分号**

登录命令为：

psql -U kong -d kong -h 127.0.0.1 -p 5432

在 work 或者 root 账户下登录 postgresql 数据库会提示权限问题.

认证权限配置文件为 /var/lib/pgsql/9.5/data/pg_hba.conf

常见的四种身份验证为：
> * trust：凡是连接到服务器的，都是可信任的。只需要提供psql用户名，可以没有对应的操作系统同名用户；
> * password 和 md5：对于外部访问，需要提供 psql 用户名和密码。对于本地连接，提供 psql 用户名密码之外，还需要有操作系统访问权。（用操作系统同名用户验证）password 和 md5 的区别就是外部访问时传输的密码是否用 md5 加密；
> * 对于外部访问，从 ident 服务器获得客户端操作系统用户名，然后把操作系统作为数据库用户名进行登录对于本地连接，实际上使用了peer；
> * peer：通过客户端操作系统内核来获取当前系统登录的用户名，并作为psql用户名进行登录。

psql 用户必须有同名的操作系统用户名。并且必须以与 psql 同名用户登录 linux 才可以登录 psql 。想用其他用户（例如 root ）登录 psql，修改本地认证方式为 trust 或者 password 即可。
```
# IPv4 local connections:
host    all             all             127.0.0.1/32            trust
host    all             all             0.0.0.0/0               trust
```

pgsql 默认只能通过本地访问，需要开启远程访问。
修改配置文件 var/lib/pgsql/9.5/data/postgresql.conf ，将 listen_address 设置为 '*'。
```
#------------------------------------------------------------------------------
# CONNECTIONS AND AUTHENTICATION
#------------------------------------------------------------------------------

# - Connection Settings -

listen_addresses = '*'          # what IP address(es) to listen on;
```

执行 sudo systemctl restart postgresql-9.5.service 重启 postgresql。


## 安装 kong
```
yum install -y https://kong.bintray.com/kong-community-edition-rpm/centos/7/kong-community-edition-0.13.1.el7.noarch.rpm

// 默认配置文件路径: /etc/kong/kong.conf.default
sudo cp /etc/kong/kong.conf.default /etc/kong/kong.conf
```

将之前安装配置好的 postgresql 信息填入 kong 配置文件中：
sudo vi /etc/kong/kong.conf
```
#------------------------------------------------------------------------------
# DATASTORE
#------------------------------------------------------------------------------

# Kong will store all of its data (such as APIs, consumers and plugins) in
# either Cassandra or PostgreSQL.
#
# All Kong nodes belonging to the same cluster must connect themselves to the
# same database.

database = postgres              # Determines which of PostgreSQL or Cassandra
                                 # this node will use as its datastore.
                                 # Accepted values are `postgres` and
                                 # `cassandra`.

pg_host = 127.0.0.1             # The PostgreSQL host to connect to.
pg_port = 5432                  # The port to connect to.
pg_user = kong                  # The username to authenticate if required.
pg_password = 123456            # The password to authenticate if required.
pg_database = kong              # The database name to connect to.

ssl = off                       # 如果不希望开放 8443 的 ssl 访问可关闭
```

现在，我们来启动Kong，因为是第一次启动，所以需要先运行迁移命令，以初始化数据库
```
// -v、-vv打印debug信息
kong migrations list -c /etc/kong/kong.conf -v

kong migrations up -c /etc/kong/kong.conf -v

kong start -c /etc/kong/kong.conf --vv
```

## 验证
```
[root@will will]# netstat -ntlp
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      3046/sshd
tcp        0      0 0.0.0.0:5432            0.0.0.0:*               LISTEN      3375/postgres
tcp        0      0 127.0.0.1:25            0.0.0.0:*               LISTEN      3261/master
tcp        0      0 0.0.0.0:8443            0.0.0.0:*               LISTEN      3449/nginx: master
tcp        0      0 127.0.0.1:8444          0.0.0.0:*               LISTEN      3449/nginx: master
tcp        0      0 0.0.0.0:8000            0.0.0.0:*               LISTEN      3449/nginx: master
tcp        0      0 127.0.0.1:8001          0.0.0.0:*               LISTEN      3449/nginx: master
tcp6       0      0 :::22                   :::*                    LISTEN      3046/sshd
tcp6       0      0 :::5432                 :::*                    LISTEN      3375/postgres
tcp6       0      0 ::1:25                  :::*                    LISTEN      3261/master
// 启动进程的名称均为nginx
```

Kong的默认运行目录在/usr/local/kong
```
[root@will will]# cd /usr/local/kong
[root@will kong]# ll
总用量 12
drwx------. 2 nobody root    6 4月  27 21:47 client_body_temp
drwx------. 2 nobody root    6 4月  27 21:47 fastcgi_temp
drwxr-xr-x. 2 root   root   65 4月  27 21:44 logs
-rw-r--r--. 1 root   root  219 5月   1 09:15 nginx.conf
-rw-r--r--. 1 root   root 4700 5月   1 09:15 nginx-kong.conf
drwxr-xr-x. 2 root   root   23 5月   1 09:15 pids
drwx------. 2 nobody root    6 4月  27 21:47 proxy_temp
drwx------. 2 nobody root    6 4月  27 21:47 scgi_temp
drwxr-xr-x. 2 root   root  114 4月  27 21:44 ssl
drwx------. 2 nobody root    6 4月  27 21:47 uwsgi_temp
// 这个目录和nginx的工作目录非常像
```

## 参考资料
```
https://docs.konghq.com/install/centos/?_ga=2.200307758.1639498594.1556679153-1351509423.1556679153

https://www.jianshu.com/p/a68e45bcadb6/

https://www.jianshu.com/p/b31990c5fb6e

https://linuxops.org/blog/kong/install.html
```

