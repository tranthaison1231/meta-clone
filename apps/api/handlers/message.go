package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	g "github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/gin-gonic/gin"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
)

func GetMessages(c *gin.Context) {
	chatID, err := strconv.ParseInt(c.Param("chatID"), 10, 64)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	messages, err := services.GetMessages(chatID)

	if err != nil {
		h.Fail400(c, err.Error())
	}

	h.Success(c, gin.H{
		"messages": messages,
	})
}

// func GetBatchMessages(c *gin.Context) {
// 	requestParams := h.ConstructPaginateRequest(c)
// 	messageId, err := strconv.ParseInt(c.Request.URL.Query().Get("messageId"), 10, 64)

// 	if err != nil {
// 		h.Fail400(c, err.Error())
// 	}

// 	isUp, err := strconv.ParseBool(c.Request.URL.Query().Get("isUp"))
// 	if err != nil {
// 		h.Fail400(c, err.Error())
// 	}

// 	// messages, err := services.GetBatchMessages(uint(messageId), isUp, requestParams)

// 	if err != nil {
// 		h.Fail400(c, err.Error())
// 		return
// 	}

// 	h.Success(c, messages)
// }

func SendMessage(c *gin.Context) {
	var req models.MessageRequest
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	chatID, err := strconv.ParseUint(c.Param("chatID"), 10, 64)

	message, err := services.CreateMessage(models.Message{
		Content: req.Content,
		ChatID:  chatID,
		OwnerID: c.MustGet("user").(*models.User).ID,
	})

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	err = services.UpdateLastMessage(chatID, *message)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"message": message,
	})
}

type Response[T any] struct {
	Action   string `json:"action"`
	Response T      `json:"response"`
}

type MessageRequestPayload struct {
	Message string `json:"message"`
}

func SendMessageSocket(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print("SendMessageSocket", event.Body)

	api, err := h.NewAPIGatewayManagementAPI(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	newMessage := Response[MessageRequestPayload]{
		Action: "SEND_MESSAGE",
		Response: MessageRequestPayload{
			Message: "Xin chao",
		},
	}

	data, err := json.Marshal(newMessage)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	input := &g.PostToConnectionInput{
		ConnectionId: aws.String(event.RequestContext.ConnectionID),
		Data:         data,
	}

	if _, err = api.PostToConnection(ctx, input); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}
