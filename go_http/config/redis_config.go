package config

type RedisConfig struct {
	ConnectType        string `yaml:"connect_type"`
	SentinelMasterName string `yaml:"sentinel_master_name"`
	Conn               string `yaml:"conn"`
	DefaultDbNum       string `yaml:"default_db_num"`
	Password           string `yaml:"password"`
	MaxIdle            string `yaml:"max_idle"`
	RedisListKey       string `yaml:"redis_list_key"`
}

func InitRedis() {
	var redis RedisConfig
	_ = initConf(BathPath+"redis_conf.yml", &redis)

}
