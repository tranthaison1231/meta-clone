package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
)

type ConnectPayload struct {
	UserID string `json:"userId"`
}

func OnUserOnline(_ context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	connectionID := event.RequestContext.ConnectionID

	fmt.Println("==================== | ON USER ONLINE | CONNECTION ID ====================", connectionID)

	var err error
	var body models.SocketRequest[ConnectPayload]

	err = json.Unmarshal([]byte(event.Body), &body)

	fmt.Println("==================== | ON USER ONLINE | ERROR JSON UNMARSHAL ====================", err)

	fmt.Println("==================== | ON USER ONLINE | BODY ====================", body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}

	_, err = services.CreateNewConnection(body.Payload.UserID, connectionID)

	fmt.Println("==================== | ON USER ONLINE | CREATE NEW CONNECTION ====================", err)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, err
}
