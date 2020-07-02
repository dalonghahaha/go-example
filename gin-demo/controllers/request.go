package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"julive/tools/coding"
)

type RequestController struct {
	BaseController
}

func NewRequestController() *RequestController {
	return &RequestController{}
}

//获取Get信息
func (c *RequestController) Get(ctx *gin.Context) {
	//通过参数名称获取值,参数对应值不存在则为""
	id := ctx.Query("id")
	name := ctx.Query("name")
	//通过参数名称获取值,带有默认值,默认值为第2个参数
	sex := ctx.DefaultQuery("sex", "1")
	//获取全部信息,返回值结构为map[string][]string
	all := ctx.Request.URL.Query()
	c.JsonOK(ctx, gin.H{
		"id":   id,
		"name": name,
		"sex":  sex,
		"all":  all,
	})
}

//获取URL绑定的信息
func (c *RequestController) URL(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.Param("name")
	c.JsonOK(ctx, gin.H{
		"id":   id,
		"name": name,
	})
}

//获取Post信息
func (c *RequestController) Post(ctx *gin.Context) {
	//通过参数名称获取值,参数对应值不存在则为""
	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	//通过参数名称获取值,带有默认值,默认值为第2个参数
	sex := ctx.DefaultPostForm("sex", "1")
	//获取全部信息,返回值结构为map[string][]string
	all := ctx.Request.PostForm
	c.JsonOK(ctx, gin.H{
		"id":   id,
		"name": name,
		"sex":  sex,
		"all":  all,
	})
}

//获取Form表单提交的文件信息
func (c *RequestController) File(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	c.JsonOK(ctx, gin.H{
		"file_name": file.Filename,
		"file_size": file.Size,
	})
}

//获取Post请求的Body源数据
func (c *RequestController) Body(ctx *gin.Context) {
	body, _ := ctx.GetRawData()
	if body != nil {
		c.JsonOK(ctx, gin.H{
			"body": string(body),
		})
	}
}

//获取Post请求的JsonBody源数据
func (c *RequestController) JsonBody(ctx *gin.Context) {
	body, _ := ctx.GetRawData()
	data := coding.JSONDecode(string(body))
	if body != nil {
		c.JsonOK(ctx, gin.H{
			"body": data,
		})
	}
}

//获取Header信息
func (c *RequestController) Header(ctx *gin.Context) {
	id := ctx.GetHeader("id")
	name := ctx.GetHeader("name")
	c.JsonOK(ctx, gin.H{
		"id":   id,
		"name": name,
	})
}

//获取Cookie信息
func (c *RequestController) Cookie(ctx *gin.Context) {
	id, _ := ctx.Cookie("id")
	name, _ := ctx.Cookie("name")
	if id == "" {
		ctx.SetCookie("id", "123", 3600, "/", "localhost", false, true)
	}
	if name == "" {
		ctx.SetCookie("name", "hehehe", 3600, "/", "localhost", false, true)
	}
	c.JsonOK(ctx, gin.H{
		"id":   id,
		"name": name,
	})
}

//请求重定向
func (c *RequestController) Redirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
}
