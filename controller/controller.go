package controller

import (
	"net/http"

	"github.com/shuwenhe/shuwen-checklist/dao"
	"github.com/shuwenhe/shuwen-checklist/model"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) { // 创建
	var todo model.Todo          // 定义变量接收绑定的值
	c.BindJSON(&todo)            // 从请求中把数据拿出来
	err := dao.CreateTodo(&todo) // 从model中获取创建事项
	if err != nil {              // 存入数据库
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(), // 错误返回错误类型
		})
	} else {
		c.JSON(http.StatusOK, todo) // 成功:给前端返回响应todo变量
	}
}

func GetTodoList(c *gin.Context) { // 查看所有待办事项
	todoList, err := dao.GetTodoList()
	if err != nil { // 查询todo表所有数据
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodoById(c *gin.Context) { // U更新操作:前端传id，后端拿到id更新状态
	id, ok := c.Params.Get("id") // 请求路径id拿到
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}

	todo, err := dao.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.BindJSON(&todo)                            // 从请求中把数据拿出来
	if err := dao.UpdateTodo(todo); err != nil { // 更新操作，拿到数据，数据保存到数据库
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodoById(c *gin.Context) { // 删除待办事项
	id, ok := c.Params.Get("id") // 路径id参数拿到
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效ID",
		})
		return
	}
	if err := dao.DeleteTodoById(id); err != nil { //
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}
}
