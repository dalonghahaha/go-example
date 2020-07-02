package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-demo/services"
)

type ProjectController struct {
	BaseController
	projectService *services.ProjectService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		projectService: services.NewProjectService(),
	}
}

func (c *ProjectController) GeProject(ctx *gin.Context) {
	id := ctx.Query("id")
	_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JsonError(ctx, "参数id格式不正确")
		return
	}
	project := c.projectService.GetProject(_id)
	c.JsonOK(ctx, project)
}
