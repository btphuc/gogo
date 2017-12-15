package userController

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/user/gogo/models"
)

//UserRegister POST register
func UserRegister(c *gin.Context) {

	name := c.PostForm("name")
	username := c.PostForm("username")
	password := c.PostForm("password")
	comfirmPassword := c.PostForm("comfirm")

	var allusers []models.UserModel

	db.Find(&allusers)

	if strings.Compare(password, comfirmPassword) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "comfirm password do not match !"})
	} else {
		//MD5 hash
		for _, item := range allusers {
			if strings.Compare(username, item.Username) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "username already exist !"})
				fmt.Println("username already exist !")
				return
			}
		}

		h := md5.New()
		io.WriteString(h, password)
		hassPassword := fmt.Sprintf("%x", h.Sum(nil))

		user := models.UserModel{Name: name, Username: username, Password: hassPassword}

		db.Save(&user)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "successfull"})
	}
}

//
