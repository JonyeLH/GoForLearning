# Redis学习与实战



## 一、基础篇
### 1、下载安装
- redis一般下载至/usr/local/src    安装至 /usr/local/redis
- 下载及安装：

    
    下载
    wget http://download.redis.io/releases/redis-5.0.5.tar.gz

    解压
    tar -xzvf redis-5.0.5.tar.gz

    编译和安装
    cd redis-5.0.5
    make PREFIX=/usr/local/redis install

    配置文件及客户端可执行移动
    cp redis.conf /usr/local/redis/bin/
    cp /usr/local/redis/bin/redis-cli /usr/local/bin

    启动
    cd /usr/local/redis/bin
    ./redis-server /usr/local/redis/bin/redis.conf

### 2、redis.conf配置文件
- vim redis.conf
- 设定关键配置

    后台运行redis-server： daemonize no  修改为 yes

    绑定host： 注释默认ip  ----  #bind 127.0.0.1
              设定网卡host  ----  bind xxx.xx.xx.xx   （ifconfig查询网卡host）

    修改端口port：  port  6301   （vim redis.conf /port）

- 生效修改

    停止原始redis-server进程: ps -ef |grep redis      kill -9 id

    重启redis-server：./redis-sever  ./redis.conf

    重启redis-cli:  ./redis-cli -h ip -p port -a password
  
    客户端关闭：shutdown   

    退出redis：exit


### 3、安装后可执行文件说明
  redis-server ：启动 redis 服务
  
  redis-cli ：进入 redis 命令客户端
  
  redis-benchmark ： 性能测试的工具
  
  redis-check-aof ： aof 文件进行检查的工具
  
  redis-check-dump ： rdb 文件进行检查的工具
  
  redis-sentinel ： 启动哨兵监控服务``



### 4、redis基本数据类型
- redis支持：string、list、hash、set、zset（有序集合）


- string操作：将string类型的key-value对存入redis

    set key value

    get key


- list操作：将string存入list中

    lpush list string

    rpop list string

    lrange list 0 10


- hash操作：将string类型的key-value对存入NAME中，适用于存储对象

    HMSET NAME key value key value key value

    HGETALL NAME


- set操作：将string类型的value存入NAME的set

    sadd NAME value

    snumbers NAME


- zset操作：将string类型的value存入NAME的有序zset，每个元素都会关联一个 double 类型的分数，通过分数来为集合中的成员进行从小到大的排序。

  - zset 的成员是唯一的,但分数(score)却可以重复

    zadd NAME score member

    ZRANGEBYSCORE NAME 0 10

### 5、redis命令

- 数据db

    select db ： 选择db


- key
    
    del key ： 删除键

    exists key ： 查询是否存在键

    EXPIRE key seconds ： 设定键的生存时间

    MOVE key db ： 将当前数据库的 key 移动到给定的数据库 db 当中



- string
  
  set key value

  get key

  SETEX key seconds value ： 将值 value 关联到 key ，并将 key 的过期时间设为 seconds


- hash

  HDEL key field2 [field2] ： 删除一个或多个哈希表字段

  HEXISTS key field  ： 查看哈希表 key 中，指定的字段是否存在

  HGET key field ： 获取存储在哈希表中指定字段的值

  GETALL key  ： 获取在哈希表中指定 key 的所有字段和值

  HKEYS key  ： 获取所有哈希表中的 key 字段

  HVALS key  ： 获取哈希表中所有 value 值


- list

  BLPOP LIST1 LIST2 .. LISTN TIMEOUT ： 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时（返回nil）或发现可弹出元素为止

  BRPOP LIST1 LIST2 .. LISTN TIMEOUT ： 移出并获取列表最后一个元素， 如果列表没有元素会阻塞列表直到等待超时（返回nil）或发现可弹出元素为止

  LPOP key  ： 移出并获取列表的第一个元素

  RPOP key  ： 移除并获取列表最后一个元素

  LPUSH key value1 [value2] ： 将一个或多个值插入到列表头部

  RPUSH key value1 [value2] ： 在列表最后中添加一个或多个值

  LRANGE key start stop  ： 获取列表指定范围内的元素

  LLEN key  ： 获取列表长度


- set

  SADD key member1 [member2] ： 向集合添加一个或多个成员

  SCARD key / SMEMBERS key ： 获取集合的成员数

  SISMEMBER key member  ： 判断 member 元素是否是集合 key 的成员

  SPOP key  : 移除并返回集合中的一个**随机**元素

  

- zset

  ZADD key score1 member1 [score2 member2] : 向有序集合添加一个或多个成员，或者更新已存在成员的分数

  ZCARD key  : 获取有序集合的成员数

  ZCOUNT key min max  : 计算在有序集合中指定区间分数的**成员数目**

  ZRANK key member  ： 返回有序集合中指定成员的索引

  ZREM key member [member ...] ： 移除有序集合中的一个或多个成员

  ZSCORE key member ： 返回有序集中，成员的分数值


## 二、进阶篇
### 1、Redis持久化
- Redis是运行在内存的键值数据库，支持将数据写落磁盘，实现断电重启后的数据恢复    
- Redis通过AOF日志、RDB快照完成持久化

#### 1）AOF持久化
##### 1-理解
- AOF(Append only file)：redis以写后的方式增量保存每个写操作命令到日志文本，具体是通过append追加到AOF缓冲区。写后记录日志优势：1、恢复数据时保证每个命令都是正确的；2、不会阻塞当前进程的操作
- AOF保存命令操作可能在宕机的过程，从而导致数据落盘存在异常服务重启后数据缺失。解决方法，AOF采用三种数据写回策略：always(同步写数据到磁盘)、everysec(每秒写数据到磁盘)、no(数据交由操作系统写入磁盘)。三种策略redis性能由低到高
- AOF是以文本的形式保存每条写操作命令，因而AOF日志文件将会十分庞大，服务重启时耗时严重。解决方法，当AOF文件大小超过一个阈值，启动AOF重写日志：利用AOF是以增量形式保存每条命令，可将多个写操作合一的形式重写一份AOF日志文件，实现优化
- AOF重写是在redis的AOF日志主进程fork一个子进程进行，所以不会阻塞当前进程的操作。

##### 2-动手
- 配置文件设置AOF 
  - 开启 AOF 持久化： redis.conf 文件中设置持久化模式 appendonly no 默认是不开启AOF，使用RDB做持久化
  - AOF 写回策略： redis.conf 文件中设置写回三种策略 appendfsync always everysec no 默认采用everysec
    - AOF 重写： redis.conf 文件中设置重写机制 
      no-appendfsync-on-rewrite no 该设置是将AOF文件数据写入磁盘，重写发生时，阻塞进程； yes  是将AOF文件写入内存，宕机时数据丢失程度较大
    - 触发重写条件：   
      auto-aof-rewrite-percentage 100 当AOF文件写达100%时
    
      auto-aof-rewrite-min-size 64mb  当AOF重写文件达到64Mb时
    

#### 2）RDB持久化
##### 1-理解
- RDB(Redis DataBase):redis将某一时刻数据以二进制文件的形式保存，实现持久化
- 







### 2、主从库一致







## 三、实战篇



