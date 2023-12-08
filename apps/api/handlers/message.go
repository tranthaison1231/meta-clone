package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	g "github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/tranthaison1231/meta-clone/api/db"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
)

func SeenMessage(c *gin.Context) {
	var req models.SeenMessageInputDto
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	message, err := services.SeenMessage(req.MessageID)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"message": message,
	})
}

func SendMessage(c *gin.Context) {
	var req models.SendMessageInputDto
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	chatID := c.Param("chatID")

	if req.MessageType == "" {
		h.Fail400(c, errors.New("MessageType is required!").Error())

		return
	}

	message, err := services.CreateMessage(models.Message{
		Content:        req.Content,
		ChatID:         chatID,
		OwnerID:        c.MustGet("user").(*models.User).ID,
		ReplyMessageID: req.ReplyMessageID,
		MessageType:    req.ReplyMessageID,
	})

	services.UpdateLastMessage(chatID, message)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	err = services.UpdateLastMessage(c.Param("chatID"), message)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"message": message,
	})
}

type MessagePayload struct {
	Message models.Message `json:"message"`
}

func SendMessageSocket(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	var err error

	api, err := h.NewAPIGatewayManagementAPI(ctx)

	fmt.Println("==================== | SEND MESSAGE SOCKET | NEW API GATEWAY ERROR ====================", err)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}

	var body models.SocketRequest[MessagePayload]

	err = json.Unmarshal([]byte(event.Body), &body)

	fmt.Println("==================== | SEND MESSAGE SOCKET | ERROR JSON UNMARSHAL ====================", err)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}

	fmt.Println("==================== | SEND MESSAGE SOCKET | BODY ====================", body)

	var chat models.Chat

	err = db.DB.Preload("Members").Where("id = ?", body.Payload.Message.ChatID).First(&chat).Error

	fmt.Println("==================== | SEND MESSAGE SOCKET | QUERY CHAT ERROR ====================", err)

	fmt.Println("==================== | SEND MESSAGE SOCKET | CHAT ====================", chat)

	var chatMemberIDs []string

	for _, member := range chat.Members {
		chatMemberIDs = append(chatMemberIDs, member.ID)
	}

	fmt.Println("==================== | SEND MESSAGE SOCKET | CHAT MEMBER IDS ====================", chatMemberIDs)

	var chatMemberConnectionIDs []string

	err = db.DB.Model(&models.Connection{}).Select("connection_id").Where("user_id IN ?", chatMemberIDs).Find(&chatMemberConnectionIDs).Error

	fmt.Println("==================== | SEND MESSAGE SOCKET | QUERY CHAT MEMBER CONNECTION IDS ERROR ====================", err)
	fmt.Println("==================== | SEND MESSAGE SOCKET | CHAT MEMBER CONNECTION IDS ====================", chatMemberConnectionIDs)

	response, err := json.Marshal(models.SocketResponse[MessagePayload]{
		Action:  "RECEIVE_MESSAGE",
		Payload: body.Payload,
	})

	fmt.Println("==================== | SEND MESSAGE SOCKET | QUERY CHAT MEMBER CONNECTION IDS ERROR ====================", err)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}

	for _, connectionID := range chatMemberIDs {
		if connectionID != event.RequestContext.ConnectionID {
			input := &g.PostToConnectionInput{
				ConnectionId: aws.String(event.RequestContext.ConnectionID),
				Data:         response,
			}

			fmt.Println("==================== | SEND MESSAGE SOCKET | POST TO CONNECTION ID ====================", connectionID)

			if _, err = api.PostToConnection(ctx, input); err != nil {
				fmt.Println("==================== | SEND MESSAGE SOCKET | POST TO CONNECTION ID ERROR ====================", err)

				return events.APIGatewayProxyResponse{}, err
			}
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}
