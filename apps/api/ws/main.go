package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayWebsocketProxyRequest) {
	fmt.Println(request)
	fmt.Println(request.RequestContext.ConnectionID)
}

func main() {
	lambda.Start(Handler)
}
