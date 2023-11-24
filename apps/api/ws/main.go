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
	Action string `json:"action"`
}

func Handler(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	bodyStr := []byte(event.Body)

	var body Body
	err := json.Unmarshal(bodyStr, &body)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch event.RequestContext.RouteKey {
	case "$connect":
		return handlers.Connect(ctx, event)
	case "$disconnect":
		return handlers.Disconnect(ctx, event)
	case "$default":
		switch body.Action {
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
