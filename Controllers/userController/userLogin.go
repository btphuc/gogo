package userController

import "github.com/gin-gonic/gin"

//UserLogin POST
func UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{"ok": 1})
}
