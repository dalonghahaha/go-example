package controllers

import (
	"fmt"
	"net/http"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (c *BaseController) JsonOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"result":  data,
	})
}

func (c *BaseController) JsonError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": message,
		"result":  nil,
	})
}

func (c *BaseController) LayoutHtml(ctx *gin.Context, layout, page string, obj map[string]interface{}) {
	gvFront := goview.New(goview.Config{
		Root:         "views/" + layout,
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})
	err := gvFront.Render(ctx.Writer, http.StatusOK, page, obj)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Render index error: %v!", err)
	}
}
