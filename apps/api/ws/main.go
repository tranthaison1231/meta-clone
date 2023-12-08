package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tranthaison1231/meta-clone/api/db"
	"github.com/tranthaison1231/meta-clone/api/handlers"
)

type Body struct {
	Action string `json:"action"`
}

func init() {
	db.ConnectDB()
}

func Handler(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	bodyStr := []byte(event.Body)

	fmt.Println("==================== HANDLER | EVENT REQUEST CONTEXT | ROUTE KEY ====================", event.RequestContext.RouteKey)
	fmt.Println("==================== HANDLER | EVENT REQUEST CONTEXT | CONNECTION ID ====================", event.RequestContext.ConnectionID)

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
	case "ONLINE":
		fmt.Println("==================== HANDLER | ONLINE CASE ====================")
		return handlers.OnUserOnline(ctx, event)
	case "SEND_MESSAGE":
		fmt.Println("==================== HANDLER | MESSAGE CASE ====================")
		return handlers.SendMessageSocket(ctx, event)
	case "$default":
		switch body.Action {
		case "MESSAGE":
			return handlers.SendMessageSocket(ctx, event)
		default:
			return events.APIGatewayProxyResponse{Body: "no handler", StatusCode: 200}, nil
		}
		// return events.APIGatewayProxyResponse{Body: "no handler", StatusCode: 200}, nil

	default:
		return events.APIGatewayProxyResponse{Body: "no handler", StatusCode: 200}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
