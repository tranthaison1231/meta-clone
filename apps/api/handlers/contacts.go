package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	text string
	from string
}

type Contact struct {
	Name        string
	Members     []string
	LastMessage Message
}

func GetMyContacts(c *gin.Context) {
	contacts := []Contact{
		{"Friends", []string{"John Doe", "Jane Smith"}, Message{text: "Hello, how are you?", from: "Jane Smith"}},
		{"Family", []string{"Mike Johnson", "Sarah Smith"}, Message{text: "Hello, how are you?", from: "Jane Smith"}},
	}

	c.JSON(http.StatusOK, gin.H{
		"Contacts": contacts,
	})
}
