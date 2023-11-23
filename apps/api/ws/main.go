package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tranthaison1231/meta-clone/api/handlers"
)

type Body struct {
	action  string
	payload interface{}
}

func Handler(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := json.Marshal(event)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	body := Body{}
	err = json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	switch event.RequestContext.RouteKey {
	case "$connect":
		return handlers.Connect(ctx, event)
	case "$disconnect":
		return handlers.Disconnect(ctx, event)
	case "$default":
		fmt.Print(event.Body)
		fmt.Println(body.action)
		switch body.action {
		case "SEND_MESSAGE":
			return handlers.SendMessageSocket(ctx, event)
		default:
			return events.APIGatewayProxyResponse{Body: "no handler", StatusCode: 200}, nil
		}
	default:
		return events.APIGatewayProxyResponse{Body: "no handler", StatusCode: 200}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
