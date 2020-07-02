package controllers

import (
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

func Panic(ctx *gin.Context) {
	panic("what the fuck!")
}
