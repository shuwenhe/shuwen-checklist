package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/shuwen-checklist/controller"
)

func SetupRouter() *gin.Engine { // Generate route
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", controller.IndexHandler)

	todo := r.Group("v1")
	{
		todo.POST("/todo", controller.CreateTodo)           // 创建待办事项
		todo.GET("/todo", controller.GetTodoList)           // 查询所有待办事项
		todo.PUT("/todo/:id", controller.UpdateTodoById)    // 更新待办事项
		todo.DELETE("/todo/:id", controller.DeleteTodoById) // 删除待办事项
	}

	return r
}
