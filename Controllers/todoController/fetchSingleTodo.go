rpackage todoController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/gogo/models"
)

//FetchSingleTodo GET:id
func FetchSingleTodo(c *gin.Context) {
	var todo models.TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}
	_todo := models.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}
