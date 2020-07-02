package server

import (
	"fmt"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"julive/components/db"
	"julive/components/logger"
	"julive/middlewares"
)

var server *gin.Engine

var registers = [...]func() error{
	logger.Register,
	db.Register,
}

func Init() (err error) {
	//依次初始化组件,只要其中一个失败直接退出
	for _, register := range registers {
		err = register()
		if err != nil {
			return
		}
	}
	//初始化服务器
	if viper.GetString("server.mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	server = gin.New()
	server.Use(middlewares.Logger)
	server.Use(middlewares.Recover)
	viewConfig := goview.DefaultConfig
	viewConfig.Root = "views"
	server.HTMLRender = ginview.New(viewConfig)
	server.Static("/assets", "assets")
	return nil
}

func Run() {
	setupRouter()
	addr := fmt.Sprintf(":%s", viper.GetString("server.port"))
	err := server.Run(addr)
	if err != nil {
		panic("web服务启动失败!")
	}
}
