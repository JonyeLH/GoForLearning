package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisSentinelConn *redis.Client

/*
connect_type: sentinel
sentinel_master_name: master6301
conn: AcsdRedis01Host:26301;AcsdRedis02Host:26301;AcsdRedis03Host:26301
password: cY0_Omr0Vhds
max_idle: 200
#默认的db 存放引擎相关信息
default_db_num: 6
*/
//
const (
	EMPSTR         = ""
	INVALIDTTL     = -2
	LEN            = 0
	RedisKeyExpire = 1 * time.Hour
)

// redis key
const (
	RedisKeyAppId   = "cs:app:appId:%s"
	RedisKeyTaskPid = "cs:task:pid:%s"
	RedisKeyStopPid = "cs:stop:pid:%s"
)

var ctx = context.Background()

func InitRedis(master_name, passwd string, db int, sentinel_addrs []string) bool {
	if redisSentinelConn == nil {
		redisSentinelConn = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    master_name,
			SentinelAddrs: sentinel_addrs,
			Password:      passwd,
			DB:            db,
		})
	}
	return true
}
func GetRedisClient() *redis.Client {
	if redisSentinelConn == nil {
		return nil
	}
	return redisSentinelConn
}

//Redis SET 命令用于设置给定 key 的值。如果 key 已经存储其他值， SET 就覆写旧值，且无视类型。
//Redis `SET key value [expiration]` command.
//https://redis.io/commands/set
func Set(key, value string) (err error) {
	// Zero expiration means the key has no expiration time.
	if err := redisSentinelConn.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

//Redis Setex 命令为指定的 key 设置值及其过期时间。
//如果 key 已经存在， SETEX 命令将会替换旧的值。
//https://redis.io/commands/setex
func SetEx(key, value string, expire time.Duration) error {
	if err := redisSentinelConn.SetEX(ctx, key, value, expire).Err(); err != nil {
		return err
	}
	return nil
}

//Redis TTL 命令以秒为单位返回 key 的剩余过期时间
//当 key 不存在时，返回 0
//当 key 存在但没有设置剩余生存时间时，返回 0
//否则，以秒为单位，返回 key 的剩余生存时间
func TTL(key string) (int, error) {
	//Redis `TTL KEY_NAME` command.
	duration, err := redisSentinelConn.TTL(ctx, key).Result()
	if err != nil {
		return INVALIDTTL, err
	}
	value := int(time.Duration.Seconds(duration))
	return value, nil
}

//Redis Get 命令用于获取指定 key 的值。
//如果 key 不存在，返回 nil 。如果key 储存的值不是字符串类型，返回一个错误
//https://redis.io/commands/get
func Get(key string) (string, error) {
	//Redis `GET key` command.
	value, err := redisSentinelConn.Get(ctx, key).Result()
	if err != nil {
		return EMPSTR, err
	}
	return value, nil
}

//Redis DEL 命令用于删除已存在的键。不存在的 key 会被忽略。
//https://redis.io/commands/del
func Del(key string) error {
	//Redis `DEL key` command.
	if err := redisSentinelConn.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

//Redis Lpush 命令将一个或多个值插入到列表头部。
//如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作。
//当 key 存在但不是列表类型时，返回一个错误。
//https://redis.io/commands/lpush
func LPush(key, value string) error {
	if err := redisSentinelConn.LPush(ctx, key, value).Err(); err != nil {
		return err
	}
	return nil
}

//Redis Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)
//如果列表不存在，一个空列表会被创建并执行 RPUSH 操作。
//当列表存在但不是列表类型时，返回一个错误
//https://redis.io/commands/rpush
func RPush(key, value string) error {
	if err := redisSentinelConn.RPush(ctx, key, value).Err(); err != nil {
		return err
	}
	return nil
}

//Redis Lpop 命令用于移除并返回列表的第一个元素。
//https://redis.io/commands/lpop
func Lpop(key string) (string, error) {
	value, err := redisSentinelConn.LPop(ctx, key).Result()
	if err != nil {
		return EMPSTR, err
	}
	return value, err
}

//Redis Llen 命令用于返回列表的长度。
//如果列表 key 不存在，则 key 被解释为一个空列表，返回 0 。
//如果 key 不是列表类型，返回一个错误
//https://redis.io/commands/llen
func Llen(key string) (int64, error) {
	value, err := redisSentinelConn.LLen(ctx, key).Result()
	if err != nil {
		return LEN, err
	}
	return value, nil
}

//Redis Lrange 返回列表中指定区间内的元素，区间以偏移量 start 和 end 指定。
//其中 0 表示列表的第一个元素， 1 表示列表的第二个元素，以此类推
//你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推
//https://redis.io/commands/lrange
func Lrange(key string, start, end int64) ([]string, error) {
	values, err := redisSentinelConn.LRange(ctx, key, start, end).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}
