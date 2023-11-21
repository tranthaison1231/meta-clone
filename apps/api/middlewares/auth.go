package middlewares

import (
	"net/http"
	"strings"

	"github.com/tranthaison1231/messenger-clone/api/services"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  "No authorization header provided",
		})
		c.Abort()
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userClaims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}
	user, err := services.GetUserByMail(userClaims.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("user", user)
	c.Next()
}
