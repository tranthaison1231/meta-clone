package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tranthaison1231/messenger-clone/api/models"
	"github.com/tranthaison1231/messenger-clone/api/services"
)

type Message struct {
	text string
	from string
}

type Chat struct {
	Name        string
	Members     []string
	LastMessage Message
}

func GetChats(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	chats, err := services.GetChats(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"chats": chats,
	})
}

func CreateChat(c *gin.Context) {
	var req models.CreateChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := c.MustGet("user").(*models.User)

	chat, err := services.CreateChat(models.Chat{
		Name:    req.Name,
		OwnerID: user.ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chat": chat,
	})
}

func AddMemberToChat(c *gin.Context) {
	var req models.AddMemberToChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := c.MustGet("user").(*models.User)

	chat, err := services.AddMemberToChat(user.ID, req.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chat": chat,
	})
}
