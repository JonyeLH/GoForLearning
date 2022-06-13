package go_http

import "go_http/config"

func main() {
	Init()


}

func Init(){
	config.InitConfigs()
	config.InitLog()


}