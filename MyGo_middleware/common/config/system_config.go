package config

type SystemConfig struct {
	HttpPort    string `yaml:"http_port"`
	GrpcPort    int    `yaml:"grpc_port"`
	SecretKey   string `yaml:"secret_key"`
	ExpiresTime int    `yaml:"expires_time"`
}

var System SystemConfig

func InitSystem() {
	_ = initConf(BathPath+"system_conf.yml", &System)
}
