package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Disconnect(_ context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("new disconnection. id: %s", event.RequestContext.ConnectionID)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}
