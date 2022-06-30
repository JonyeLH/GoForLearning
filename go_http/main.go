package main

import (
	config2 "MyGo_middleware/common/config"
	"go_http/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Init()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	UnInit()
}

func Init() {
	config2.InitConfigs()
	config2.InitLog()
	config2.InitMysql()
	config2.InitRedis()
	config2.InitSystem()

	router.HttpServerStart()
}

func UnInit() {
	router.HttpServerStop()
}
