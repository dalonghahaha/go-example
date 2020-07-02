package server

import (
	"gin-demo/controllers"
)

func setupRouter() {
	server.GET("/", controllers.Index)
	server.GET("/ping", controllers.Ping)
	server.GET("/panic", controllers.Panic)
	//请求参数解析demo
	request := server.Group("/request")
	{
		requestController := controllers.NewRequestController()
		request.GET("/get", requestController.Get)
		request.GET("/url/:id/:name", requestController.URL)
		request.POST("/post", requestController.Post)
		request.POST("/file", requestController.File)
		request.POST("/body", requestController.Body)
		request.POST("/json_body", requestController.JsonBody)
		request.GET("/header", requestController.Header)
		request.GET("/cookie", requestController.Cookie)
		request.GET("/redirect", requestController.Redirect)
	}
	//返回数据类型demo
	response := server.Group("/response")
	{
		responseController := controllers.NewResponseController()
		response.GET("/html", responseController.Html)
		response.GET("/string", responseController.String)
		response.GET("/json", responseController.JSON)
		response.GET("/xml", responseController.XML)
		response.GET("/data", responseController.Data)
		response.GET("/layout", responseController.Layout)
	}
	employee := server.Group("/employee")
	{
		employeeController := controllers.NewEmployeeController()
		employee.GET("/", employeeController.GetEmployee)
	}
	project := server.Group("/project")
	{
		projectController := controllers.NewProjectController()
		project.GET("/", projectController.GeProject)
	}
	//resultful
	black_word := server.Group("/black_word")
	{
		blackWordController := controllers.NewBlackWordController()
		black_word.GET("/", blackWordController.GetAllBlackWord)
		black_word.GET("/:id", blackWordController.GetBlackWord)
		black_word.POST("/", blackWordController.CreateBlackWord)
		black_word.PUT("/:id", blackWordController.UpdateBlackWord)
		black_word.DELETE("/:id", blackWordController.DeleteBlackWord)
	}
}
