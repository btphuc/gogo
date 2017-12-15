package todoController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/gogo/models"
)

//UpdateTodo PUT:id
func UpdateTodo(c *gin.Context) {
	var todo models.TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}
