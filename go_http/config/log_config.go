package config

type LogConfig struct {
	EventLogFileName 	string		`yaml:"event_log_file_name"`
	TraceLogFileName	string		`yaml:"trace_log_file_name"`
}

var Log LogConfig

func InitLog() {
	_ = initConf(BathPath + "log_conf.yml", &Log)
}