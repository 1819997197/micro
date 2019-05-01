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



