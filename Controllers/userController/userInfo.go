package userController

import (
	"net/http"

	"github.com/user/gogo/models"

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
		return []byte("keysecret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		var user models.UserModel
		var userID = claims["userID"]

		db.First(&user, userID)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "fail",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
			"user":   user,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err})
	}

}
