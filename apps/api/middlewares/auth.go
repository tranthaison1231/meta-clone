package middlewares

import (
	"strings"

	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/services"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		h.Fail403(c, "No authorization header provided")
		c.Abort()
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]

	userClaims, err := services.ParseToken(token)
	if err != nil {
		h.Fail403(c, err.Error())
		c.Abort()
		return
	}
	user, err := services.GetUserByMail(userClaims.Email)
	if err != nil {
		h.Fail403(c, err.Error())
		c.Abort()
		return
	}
	c.Set("user", user)
	c.Next()
}
