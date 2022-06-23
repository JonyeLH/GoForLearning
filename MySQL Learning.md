# MySQL学习





## MySql基本学习

### 一、环境部署
1、官网下载Windows、Linux安装包

2、解压文件：tar -zxvfmysql-8.0.28-linux-glibc2.12-x86_64.tar

3、检查是否有mysql用户组和mysql用户,没有就添加有就忽略： groups mysql

添加用户组和用户： groupadd mysql && useradd -r -g mysql mysql

4、创建文件目录： mkdir /usr/local/mysql

授权：chown mysql:mysql -R /data/mysql

5、修改配置文件：vim /etc/my.cof 
若是文件不存在则新建

6、编辑my.cof文件：
[mysqld]
bind-address=0.0.0.0
port=3306
user=mysql
basedir=/usr/local/mysql
datadir=/data/mysql
socket=/tmp/mysql.sock
log-error=/data/mysql/mysql.err
pid-file=/data/mysql/mysql.pid
character_set_server=utf8mb4
symbolic-links=0
explicit_defaults_for_timestamp=true

7、初始化：
将解压后的文件复制到 /usr/local/mysql

cd /usr/local/mysql/bin/

./mysqld --defaults-file=/etc/my.cnf --basedir=/usr/local/mysql/ --datadir=/data/mysql/ --user=mysql --initialize


8、启动mysql：
cp /usr/local/mysql/support-files/mysql.server /etc/init.d/mysql

service mysql start

9、修改密码：
1)配置文件添加跳过密码检测模式： vim /etc/my.conf  

添加skip-grant-tables   保存退出

2)重启服务：service mysql restart

3)登录用户root：mysql -u root -p

4)刷新权限：
use mysql;

update user set host='%' where user='root';

5)修改密码：
ALTER USER "root"@"%" IDENTIFIED  BY "1234";

FLUSH PRIVILEGES; 　　 　　 #刷新

6)退出mysql服务：quit

把/etc/my.cnf免密删掉。

7)重启服务

service mysql restart

8)登陆mysql

/usr/local/mysql/bin/mysql -u root -p 

9)远程连接mysql

按照MySQL8.0前后版本，不同方法刷远程连接权限：

8.0之前版本：


8.0之后版本：

注意：若MySQL部署在云服务上，要在防火墙添加规则。将MySQL端口添加即可。

### 二、基本语法
1、创建数据库：

create database <数据库名>;

2、为用户分配操作权限：

GRANT SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,ALTER ON 数据库名.* TO 数据库名@localhost IDENTIFIED BY '密码';

3、显示数据库：
show databases;

4、删除数据库、用户、数据表：

drop database <数据库名>;   

drop user 'user'@'host';

show columns from 表名;

5、使用数据库：

use <数据库名>;

6、显示当前选择（连接）的数据库：

select database();

select user,host from user;

7、创建数据表:

create table <表名>(<字段名1> <类型1> [,..<字段名n> <类型n>]);

mysql> create table MyClass(
> id int(4) not null primary key auto_increment,
> name char(20) not null,
> sex int(4) not null default '0',
> degree double(16,2));


8、获取数据表结构:

desc 表名;

show columns from 表名;


9、表中插入数据：

insert into <表名> [(<字段名1>[,..<字段名n > ])] values ( 值1 )[, ( 值n )];


10、查询表中的数据：

select <字段1, 字段2, ...> from < 表名 > where < 表达式 >;

select * from MyClass order by id limit 0,2;


11、删除表中的数据

delete from 表名 where 表达式;

delete from MyClass where id=1;


12、修改表中的数据

update 表名 set 字段=新值,… where 条件;

update MyClass set name='Mary' where id=1;


13、增加表的字段：（待补充）

alter table 表名 add字段 类型 其他;

alter table MyClass add passtest int(4) default '0';


14、修改表名：（要有权限）

rename table 原表名 to 新表名;

rename table MyClass to YouClass;


15、备份数据库






## MySQL进阶学习

### 一、常用函数
- 查询两个日期之间的**天数**
    
    DATEDIFF(date1, date2)
 
    select DATEDIFF('2008-12-30','2008-12-29') AS DiffDate    #返回结果赋给DiffDate

    DateDiff(NOW(),start_time)=0 #可以视为当天


### 二、日志系统



### 三、事务隔离








## MySQL实战

### 一、问题分析