//  bm, err := caches.NewCache("engine_config", `{"conn":"127.0.0.1:11211"}`)
//
package redis

import (
	caches "MyGo_middleware/common/Cache"
	"MyGo_middleware/logs"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"

	"strings"
)

func init() {

}

var (
	// DefaultKey the collection name of engine_config for cache adapter.
	DefaultKey = "cacheRedis"
)

// Cache is Redis cache adapter.
type Cache struct {
	p                  *redis.Pool // engine_config connection pool
	connectType        string
	sentinelMasterName string
	conninfo           string
	dbNum              int
	key                string
	password           string
	maxIdle            int
}

// NewRedisCache create new engine_config cache with default collection name.
func NewRedisCache() caches.Cache {
	return &Cache{key: DefaultKey}
}

// actually do the engine_config cmds, args[0] must be the key name.
func (rc *Cache) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}
	args[0] = rc.associate(args[0])
	c := rc.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

// associate with config key.
func (rc *Cache) associate(originKey interface{}) string {
	return fmt.Sprintf("%s%s", rc.key, originKey)
}

// Get cache from engine_config.
func (rc *Cache) Get(key string) (string, error) {
	v, err := rc.do("GET", key)

	return redis.String(v, err)
}

// Get cache from engine_config.
func (rc *Cache) Expire(key string, expire int) error {
	_, err := rc.do("EXPIRE", key, expire)
	return err
}

// HGet cache from engine_config.
func (rc *Cache) HGet(hashKey string, key string) (string, error) {
	v, err := rc.do("HGET", hashKey, key)
	return redis.String(v, err)
}

// HSet cache from engine_config.
func (rc *Cache) HSet(hashKey string, key string, value string) error {
	_, err := rc.do("HSET", hashKey, key, value)
	return err
}

// HGet cache from engine_config.
func (rc *Cache) HGetAll(hashKey string) ([]string, error) {
	c := rc.p.Get()
	defer c.Close()
	values, err := redis.Values(c.Do("HGETALL", hashKey))
	return redis.Strings(values, err)
}

// GetMulti get cache from engine_config.
func (rc *Cache) GetMulti(keys []string) ([]string, error) {
	c := rc.p.Get()
	defer c.Close()
	var args []interface{}
	for _, key := range keys {
		args = append(args, rc.associate(key))
	}
	values, err := redis.Values(c.Do("MGET", args...))
	return redis.Strings(values, err)
}

// Put put cache to engine_config.
func (rc *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	_, err := rc.do("SETEX", key, int64(timeout/time.Second), val)
	return err
}

func (rc *Cache) Set(key string, val interface{}) error {
	_, err := rc.do("SET", key, val)
	return err
}

// Put put cache to engine_config.
func (rc *Cache) HPut(hashKey string, key string, val interface{}, timeout time.Duration) error {
	_, err := rc.do("HSETEX", hashKey, key, int64(timeout/time.Second), val)
	return err
}

// Delete delete cache in engine_config.
func (rc *Cache) Delete(key string) error {
	_, err := rc.do("DEL", key)
	return err
}

// Delete delete cache in engine_config.
func (rc *Cache) HDelete(hashKey string, key string) error {
	_, err := rc.do("HDEL", hashKey, key)
	return err
}

// IsExist check cache's existence in engine_config.
func (rc *Cache) IsExist(key string) bool {
	v, err := redis.Bool(rc.do("EXISTS", key))
	if err != nil {
		return false
	}
	return v
}

// IsExist check cache's existence in engine_config.
func (rc *Cache) HIsExist(hashKey string, key string) bool {
	v, err := redis.Bool(rc.do("HEXISTS", hashKey, key))
	if err != nil {
		return false
	}
	return v
}

// Incr increase counter in engine_config.
func (rc *Cache) Incr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, 1))
	return err
}

// Incr increase counter in engine_config.
func (rc *Cache) HIncr(hashKey string, key string) error {
	_, err := redis.Bool(rc.do("HINCRBY", hashKey, key, 1))
	return err
}

// Incr increase counter in engine_config.
func (rc *Cache) HInDecrBatch(hashKey string, key string, count int64) error {
	_, err := redis.Bool(rc.do("HINCRBY", hashKey, key, count))
	return err
}

// Decr decrease counter in engine_config.
func (rc *Cache) Decr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, -1))
	return err
}

// Decr decrease counter in engine_config.
func (rc *Cache) HDecr(hashKey string, key string) error {
	_, err := redis.Bool(rc.do("HINCRBY", hashKey, key, -1))
	return err
}

// ClearAll clean all cache in engine_config. delete this engine_config collection.
func (rc *Cache) ClearAll() error {
	c := rc.p.Get()
	defer c.Close()
	cachedKeys, err := redis.Strings(c.Do("KEYS", rc.key+":*"))
	if err != nil {
		return err
	}
	for _, str := range cachedKeys {
		if _, err = c.Do("DEL", str); err != nil {
			return err
		}
	}
	return err
}

// StartAndGC start engine_config cache adapter.
// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in engine_config are stored forever,
// so no gc operation.
func (rc *Cache) StartAndGC(config string) error {
	var cf map[string]string
	json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}
	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}

	// Format engine_config://<password>@<host>:<port>
	cf["conn"] = strings.Replace(cf["conn"], "engine_config://", "", 1)
	if i := strings.Index(cf["conn"], "@"); i > -1 {
		cf["password"] = cf["conn"][0:i]
		cf["conn"] = cf["conn"][i+1:]
	}

	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	if _, ok := cf["maxIdle"]; !ok {
		cf["maxIdle"] = "3"
	}
	rc.key = cf["key"]
	rc.conninfo = cf["conn"]
	rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	rc.password = cf["password"]
	rc.maxIdle, _ = strconv.Atoi(cf["maxIdle"])
	rc.connectType = cf["connectType"]
	rc.sentinelMasterName = cf["sentinelMasterName"]

	rc.connectInit()

	c := rc.p.Get()
	defer c.Close()

	return c.Err()
}

// connect to engine_config.
func (rc *Cache) connectInit() {
	rc.p = &redis.Pool{
		MaxActive:   300,
		MaxIdle:     rc.maxIdle,
		IdleTimeout: time.Duration(10) * time.Second,
		Wait:        false,
	}

	if rc.connectType == "sentinel" {
		//哨兵
		sntnl := &Sentinel{
			Addrs:      strings.Split(rc.conninfo, ";"),
			MasterName: rc.sentinelMasterName,
			Dial: func(addr string) (redis.Conn, error) {
				c, err := redis.Dial("tcp", addr,
					redis.DialConnectTimeout(time.Duration(500)*time.Millisecond),
					redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
					redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		}
		rc.p.Dial = func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				logs.Error("sentinel dial master addr failed: %v", err)
				return nil, err
			}
			conn, err := redis.Dial("tcp", masterAddr,
				redis.DialDatabase(rc.dbNum),
				redis.DialPassword(rc.password),
				redis.DialConnectTimeout(time.Duration(500)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond))
			if err != nil {
				logs.Error("Cache dial failed during pool initialization: %v", err)
				return nil, err
			}
			return conn, nil
		}
		rc.p.TestOnBorrow = func(c redis.Conn, t time.Time) error {
			if !TestRole(c, "master") {
				return errors.New("Role check failed")
			} else {
				return nil
			}
		}
	} else {
		//直连
		rc.p.Dial = func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", rc.conninfo,
				redis.DialDatabase(rc.dbNum),
				redis.DialPassword(rc.password),
				redis.DialConnectTimeout(time.Duration(500)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond))
			if err != nil {
				logs.Error("Cache dial failed during pool initialization: %v", err)
				return nil, err
			}
			return conn, nil
		}
	}
}
