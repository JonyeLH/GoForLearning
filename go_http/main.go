package main

import (
	"go_http/config"
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
	config.InitConfigs()
	config.InitLog()
	config.InitMysql()
	config.InitRedis()
	config.InitSystem()

	router.HttpServerStart()

}

func UnInit() {

}
