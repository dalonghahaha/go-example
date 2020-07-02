package controllers

import "github.com/gin-gonic/gin"

import "net/http"

import "io/ioutil"

type ResponseController struct {
	BaseController
}

func NewResponseController() *ResponseController {
	return &ResponseController{}
}

func (c *ResponseController) Html(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "demo.html", gin.H{
		"id":   1,
		"name": "乔布斯",
		"options": map[string]string{
			"1": "男",
			"2": "女",
		},
	})
}

func (c *ResponseController) String(ctx *gin.Context) {
	ctx.String(http.StatusOK, "逗你玩")
}

func (c *ResponseController) JSON(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "乔布斯",
		"options": map[string]string{
			"1": "男",
			"2": "女",
		},
	})
}

func (c *ResponseController) XML(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"id":   1,
		"name": "乔布斯",
	})
}

func (c *ResponseController) Data(ctx *gin.Context) {
	data, _ := ioutil.ReadFile("assets/image/demo.jpeg")
	ctx.Data(http.StatusOK, "image/jpeg", data)
}

func (c *ResponseController) Layout(ctx *gin.Context) {
	c.LayoutHtml(ctx, "frontend", "index", gin.H{
		"title": "Frontend title!",
	})
}
