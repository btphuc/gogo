package todoController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/gogo/models"
)

//DeleteTodo DELETE:id
func DeleteTodo(c *gin.Context) {
	var todo models.TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
