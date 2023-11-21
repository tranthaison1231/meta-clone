package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tranthaison1231/messenger-clone/api/models"
	"github.com/tranthaison1231/messenger-clone/api/services"
)

func GetMessages(c *gin.Context) {
	chatID, err := strconv.ParseInt(c.Param("chatID"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	messages, err := services.GetMessages(chatID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"messages": messages,
	})
}

func SendMessage(c *gin.Context) {
	var req models.MessageRequest
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
	chatID, err := strconv.ParseUint(c.Param("chatID"), 10, 64)

	message, err := services.CreateMessage(models.Message{
		Content: req.Content,
		ChatID:  chatID,
		OwnerID: c.MustGet("user").(*models.User).ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	err = services.UpdateLastMessage(chatID, *message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
