package handlers

import (
	"fmt"
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

func GetChats(c *gin.Context) {
	requestParams := h.ConstructPaginateRequest(c)
	memberIdsStr := c.Request.URL.Query().Get("memberIds")
	memberIds := strings.Split(memberIdsStr, ",")

	var parsedMemberIds []uint

	if len(memberIds) > 1 {
		for _, value := range memberIds {
			parsedValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				h.Fail400(c, err.Error())
				return
			}
			parsedMemberIds = append(parsedMemberIds, uint(parsedValue))
		}
	}

	chats, err := services.GetChats(requestParams, parsedMemberIds)

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
	targetMessageId, err := strconv.ParseInt(c.Request.URL.Query().Get("targetMessageId"), 10, 64)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	chatID, err := strconv.ParseUint(c.Param("chatID"), 10, 64)

	fmt.Println("chatId", chatID)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	isUp, err := strconv.ParseBool(c.Request.URL.Query().Get("isUp"))
	if err != nil {
		h.Fail400(c, err.Error())

	}

	messages, err := services.GetChatMessages(&models.GetChatMessagesRequest{
		PaginateRequest: *requestParams,
		ChatID:          uint(chatID),
		IsUp:            isUp,
		TargetMessageID: uint(targetMessageId),
	})

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, messages)
}
