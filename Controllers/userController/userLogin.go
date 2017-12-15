package userController

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/user/gogo/models"
)

//UserLogin POST
func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// var users []models.UserModel
	var users []models.UserModel

	db.Where("username = ?", username).Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "username does not exist !!"})
		return
	}

	h := md5.New()
	io.WriteString(h, password)
	hassPassword := fmt.Sprintf("%x", h.Sum(nil))

	if strings.Compare(hassPassword, users[0].Password) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "password does not match !!"})
		return
	}

	// fmt.Println(users[0])
	tokenKey := []byte("gogopowerrangers!!")

	type tokenClaim struct {
		username string `json:"username"`
		password string `json:"password"`
		jwt.StandardClaims
	}

	claims := tokenClaim{
		username,
		password,
		jwt.StandardClaims{
			Issuer:    "tui ne",
			ExpiresAt: time.Now().Unix() + 10,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "not good >_< "})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "token": tokenString, "username": username, "password": password})
}
