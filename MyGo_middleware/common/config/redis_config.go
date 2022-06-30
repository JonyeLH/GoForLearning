package config

import (
	redis "MyGo_middleware/redis"
	"fmt"
	"strconv"
	"strings"
)

type RedisConfig struct {
	ConnectType        string `yaml:"connect_type"`
	SentinelMasterName string `yaml:"sentinel_master_name"`
	Conn               string `yaml:"conn"`
	DefaultDbNum       string `yaml:"default_db_num"`
	Password           string `yaml:"password"`
	MaxIdle            string `yaml:"max_idle"`
	RedisListKey       string `yaml:"redis_list_key"`
}
var RedisConf RedisConfig

func InitRedis() bool {
	var myRedis RedisConfig
	_ = initConf(BathPath+"redis_conf.yml", &myRedis)
	addrList := strings.Split(myRedis.Conn, ";")
	db, err := strconv.Atoi(myRedis.DefaultDbNum)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if !redis.InitRedis(myRedis.SentinelMasterName,
		myRedis.Password, db, addrList) {
		return false
	}
	fmt.Println("redis init  success ")
	RedisConf = myRedis
	return true
}
