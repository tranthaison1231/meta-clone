package h

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data, "code": http.StatusOK})
}

func Fail400(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": msg, "code": http.StatusBadRequest})
}

func Fail401(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": msg, "code": http.StatusUnauthorized})
}

func Fail403(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": msg, "code": http.StatusForbidden})
}

func Fail404(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": msg, "code": http.StatusNotFound})
}

func Fail409(c *gin.Context, msg string) {
	c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": msg, "code": http.StatusConflict})
}

func Fail502(c *gin.Context, msg string) {
	c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": msg, "code": http.StatusBadGateway})
}
