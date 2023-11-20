package middlewares

import (
	"github.com/tranthaison1231/messenger-clone/api/services"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	userClaims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(401, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}
	user, err := services.GetUserByMail(userClaims.Email)
	if err != nil {
		c.JSON(401, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("user", user)
	c.Next()
}
