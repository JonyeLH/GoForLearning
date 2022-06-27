package router

import (
	"MyGo_middleware/common/config"
	"MyGo_middleware/logs"
	"MyGo_middleware/proxy"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_http/controller"
	"log"
	"net/http"
	"time"
)

var server *http.Server

func router(ginRouter *gin.Engine) {
	Group := ginRouter.Group("/")
	Group.Handle(http.MethodPost, "/log_test", controller.HttpBase.LogTest)
	Group.Handle(http.MethodPost, "/mysql_test", controller.HttpBase.MysqlTest)
}

func HttpServerStart() {
	ginRouter := proxy.NewGinRouter()
	router(ginRouter)

	server = &http.Server{
		Addr:    fmt.Sprintf(":%s", config.System.HttpPort),
		Handler: ginRouter,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("start http service err:%s", err.Error())
		}
	}()

	logs.Info("start http service bind on %s", config.System.HttpPort)
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logs.Error("HttpServerStop err:%s\n", err.Error())
	}
	logs.Info("HttpServer stopped")
}
