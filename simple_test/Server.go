package main

import (
	"fmt"
	"net/http"
)

/*
todo:
1、内部结构的优化：分出controller层、business层
2、功能完善：完成更多的路由功能
*/
type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc handlerFunc) {
	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	fmt.Println("启动监听端口：", address)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server { //go的不定参数写法
	handler := NewHandlerBasedMap()
	var root Filter = handler.ServerHTTP
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	return &sdkHttpServer{ //当返回实际类型所实现的接口的时候，需要返回指针
		Name:    name,
		handler: handler,
		root:    root,
	}
}

func SignUp(ctx *Context) {
	resp := &signUpReq{}
	err := ctx.ReadJson(resp)
	if err != nil {
		err = ctx.BadRequestJson(err)
		if err != nil {
			return
		}
	}
	rep := commonResponse{
		BizCode: 2,
		Msg:     "success",
		Data:    123,
	}
	err = ctx.OkJson(rep)
	if err != nil {
		err = ctx.BadRequestJson(err)
		if err != nil {
			return
		}
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
