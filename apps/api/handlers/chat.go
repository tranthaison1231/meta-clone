package handlers

import (
	"strconv"
	"strings"

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

func GetChat(c *gin.Context) {
	chatID := c.Param("chatID")

	chat, err := services.GetChat(chatID)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}
	h.Success(c, gin.H{
		"chat": chat,
	})
}

func GetChats(c *gin.Context) {
	pagination := h.ConstructPaginateRequest(c)
	memberIdsStr := c.Request.URL.Query().Get("memberIds")
	isSingleChat, err := strconv.ParseBool(c.Request.URL.Query().Get("isSingleChat"))

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	memberIds := strings.Split(memberIdsStr, ",")

	chats, err := services.GetChats(&models.GetChatsRequest{
		PaginateRequest: *pagination,
		MemberIds:       memberIds,
		IsSingleChat:    isSingleChat,
	})

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, chats)
}

func CreateChat(c *gin.Context) {
	var req models.CreateChatRequest
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	chat, err := services.CreateChat(req.MemberIDs)

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

func GetChatMessages(c *gin.Context) {
	requestParams := h.ConstructPaginateRequest(c)
	targetMessageId := c.Request.URL.Query().Get("targetMessageId")

	chatID := c.Param("chatID")

	isUp, err := strconv.ParseBool(c.Request.URL.Query().Get("isUp"))
	if err != nil {
		h.Fail400(c, err.Error())

	}

	messages, err := services.GetChatMessages(&models.GetChatMessagesRequest{
		PaginateRequest: *requestParams,
		ChatID:          chatID,
		IsUp:            isUp,
		TargetMessageID: targetMessageId,
	})

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, messages)
}
