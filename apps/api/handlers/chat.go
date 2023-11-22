package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	h "github.com/tranthaison1231/messenger-clone/api/helpers"
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
		h.Fail400(c, err.Error())
	}

	h.Success(c, gin.H{
		"chats": chats,
	})
}

func CreateChat(c *gin.Context) {
	var req models.CreateChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Fail400(c, err.Error())
		return
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		h.Fail400(c, err.Error())
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
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Fail400(c, err.Error())
		return
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		h.Fail400(c, err.Error())
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
