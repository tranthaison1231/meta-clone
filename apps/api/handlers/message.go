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
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
)

func SendMessage(c *gin.Context) {
	var req models.MessageRequest
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	chatID := c.Param("chatID")

	message, err := services.CreateMessage(models.Message{
		Content: req.Content,
		ChatID:  chatID,
		OwnerID: c.MustGet("user").(*models.User).ID,
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
