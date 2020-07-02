package main

import (
	"fmt"

	"github.com/spf13/viper"

	"gin-demo/server"
)

func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	//初始化全部的配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//服务初始化
	err = server.Init()
	if err != nil {
		panic(fmt.Errorf("server init error!: %s \n", err))
	}
	server.Run()
}
