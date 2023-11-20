package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessagesByContact(c *gin.Context) {
	contacts := []string{"John Doe", "Jane Smith", "Mike Johnson"}

	c.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
	})
}

func SendMessageToContact(c *gin.Context) {
	contacts := []string{"John Doe", "Jane Smith", "Mike Johnson"}
	c.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
	})
}
