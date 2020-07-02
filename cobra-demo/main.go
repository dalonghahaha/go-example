package main

import (
	"fmt"

	"github.com/spf13/viper"

	"cobra-demo/cmd"
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
	//命令初始化
	err = cmd.Init()
	if err != nil {
		panic(fmt.Errorf("server init error!: %s \n", err))
	}
	cmd.Execute()
}
