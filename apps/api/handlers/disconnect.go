package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tranthaison1231/meta-clone/api/services"
)

func Disconnect(_ context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Disconnected, ID= %s", event.RequestContext.ConnectionID)

	connectionID := event.RequestContext.ConnectionID

	err := services.RemoveConnection(connectionID)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}
