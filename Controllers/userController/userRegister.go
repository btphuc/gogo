package userController

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/gogo/models"
)

//UserRegister POST register
func UserRegister(c *gin.Context) {

	name := c.PostForm("name")
	username := c.PostForm("username")
	password := c.PostForm("password")

	//MD5 hash

	h := md5.New()
	io.WriteString(h, password)
	hassPassword := fmt.Sprintf("%x", h.Sum(nil))

	user := models.UserModel{Name: name, Username: username, Password: hassPassword}

	db.Save(&user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "successfull"})
}
