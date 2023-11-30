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
	fmt.Println("================== SOCKET HANDLER | MESSAGE CASE ==================", event.RequestContext.RouteKey)

	switch event.RequestContext.RouteKey {
	case "$connect":
		return handlers.Connect(ctx, event)
	case "$disconnect":
		return handlers.Disconnect(ctx, event)
	case "MESSAGE":
		fmt.Println("================== SOCKET HANDLER | MESSAGE CASE | ROUTE KEY ==================", event.RequestContext.RouteKey)
		fmt.Println("================== SOCKET HANDLER | MESSAGE CASE | BODY ACTION ==================", body.Action)

		return handlers.SendMessageSocket(ctx, event)
	case "$default":
		// fmt.Println("================== SOCKET HANDLER | DEFAULT CASE | BODY ACTION ==================", body.Action)
		// switch body.Action {
		// case "MESSAGE":
		// 	fmt.Println("================== SOCKET HANDLER | MESSAGE CASE ==================", event.RequestContext.RouteKey)
		// 	return handlers.SendMessageSocket(ctx, event)
		// default:
		// 	return events.APIGatewayProxyResponse{Body: "no handler", StatusCode: 200}, nil
		// }
		return events.APIGatewayProxyResponse{Body: "no handler", StatusCode: 200}, nil

	default:
		return events.APIGatewayProxyResponse{Body: "no handler", StatusCode: 200}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
