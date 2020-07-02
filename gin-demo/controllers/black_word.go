package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-demo/models/comjia"
	"gin-demo/services"
)

type BlackWordController struct {
	BaseController
	blackService *services.BlackService
}

func NewBlackWordController() *BlackWordController {
	return &BlackWordController{
		blackService: services.NewBlackService(),
	}
}

func (c *BlackWordController) GetAllBlackWord(ctx *gin.Context) {
	list := c.blackService.GetAll()
	c.JsonOK(ctx, list)
}

func (c *BlackWordController) GetBlackWord(ctx *gin.Context) {
	id := ctx.Param("id")
	_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JsonError(ctx, "参数id格式不正确")
		return
	}
	blackWord := c.blackService.GetBlackWord(_id)
	c.JsonOK(ctx, blackWord)
}

func (c *BlackWordController) CreateBlackWord(ctx *gin.Context) {
	var blackWord comjia.CjBlackWord
	if err := ctx.ShouldBind(&blackWord); err != nil {
		c.JsonError(ctx, fmt.Sprintf("参数验证失败:%v", err))
		return
	}
	ok := c.blackService.CreateBlackWord(&blackWord)
	if ok {
		c.JsonOK(ctx, blackWord)
	} else {
		c.JsonError(ctx, "创建失败")
	}
}

func (c *BlackWordController) UpdateBlackWord(ctx *gin.Context) {
	id := ctx.Param("id")
	_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JsonError(ctx, "参数id格式不正确")
		return
	}
	var blackWord comjia.CjBlackWord
	if err := ctx.ShouldBind(&blackWord); err != nil {
		c.JsonError(ctx, fmt.Sprintf("参数验证失败:%v", err))
		return
	}
	blackWord.ID = _id
	ok := c.blackService.UpdateBlackWord(&blackWord)
	if ok {
		c.JsonOK(ctx, blackWord)
	} else {
		c.JsonError(ctx, "更新失败")
	}
}

func (c *BlackWordController) DeleteBlackWord(ctx *gin.Context) {
	id := ctx.Param("id")
	_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JsonError(ctx, "参数id格式不正确")
		return
	}
	result := c.blackService.DeleteBlackWord(_id)
	c.JsonOK(ctx, result)
}
