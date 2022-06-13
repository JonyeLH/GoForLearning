package main

import "go_http/config"

func main() {
	Init()

}

func Init() {
	config.InitConfigs()
	config.InitLog()
	config.InitMysql()
	config.InitRedis()

}
