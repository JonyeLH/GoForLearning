package main

import (
	"net/http"
)

/*
todo:
1、添加日志系统
2、接入gin框架
3、接入gorm数据操作框架
 */

func main() {
	server := NewHttpServer("Server_test")
	server.Route(http.MethodGet, "/user/signUp", SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err) //如果进入错误，panic直接停止main函数，panic是一种不可恢复的错误判断，一旦使用了，后面的代码就不会执行包括goroutine
		//在一旦错误失败的出现，整个代码就没意义执行的场景，使用panic,一般使用error即可
	}
}
