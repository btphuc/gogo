package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/user/gogo/Controllers/todoController"
	// models
)

func main() {

	router := gin.Default()

	todo := router.Group("/api/v1/todos")
	{
		// todo.GET("/test", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{"ok": 1})
		// })
		todo.POST("/", todoController.CreateTodo)
		todo.GET("/", todoController.FetchAllTodo)
		todo.GET("/:id", todoController.FetchSingleTodo)
		todo.PUT("/:id", todoController.UpdateTodo)
		todo.DELETE("/:id", todoController.DeleteTodo)
	}

	user := router.Group("/api/user/")
	{
		user.POST("/register", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": 1})
		})
		user.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": 1})
		})
	}

	router.Run(":3000")

}
