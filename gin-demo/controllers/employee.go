package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-demo/services"
)

type EmployeeController struct {
	BaseController
	employeeService *services.EmployeeService
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{
		employeeService: services.NewEmployeeService(),
	}
}

func (c *EmployeeController) GetEmployee(ctx *gin.Context) {
	id := ctx.Query("id")
	_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JsonError(ctx, "参数id格式不正确")
		return
	}
	employee := c.employeeService.GetEmployee(_id)
	c.JsonOK(ctx, employee)
}
