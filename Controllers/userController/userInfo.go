package userController

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//UserInfo Get user info
func UserInfo(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	// fmt.Println(token)

	if len(tokenString) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "please log in"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("gogopowerrangers!!"), nil
	})

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "token": token.Claims, "err": err})
}
