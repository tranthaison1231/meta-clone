package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
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
		h.Fail400(c, err.Error())
	}

	h.Success(c, gin.H{
		"chats": chats,
	})
}

func CreateChat(c *gin.Context) {
	var req models.CreateChatRequest
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	user := c.MustGet("user").(*models.User)

	chat, err := services.CreateChat(models.Chat{
		Name:    req.Name,
		OwnerID: user.ID,
	})

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"chat": chat,
	})
}

func AddMemberToChat(c *gin.Context) {
	var req models.AddMemberToChatRequest
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	chatID, err := strconv.ParseUint(c.Param("chatID"), 10, 64)

	chat, err := services.AddMemberToChat(chatID, req.UserID)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"chat": chat,
	})
}
