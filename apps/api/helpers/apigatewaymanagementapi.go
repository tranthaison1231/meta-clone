package h

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

func NewAPIGatewayManagementAPI(ctx context.Context) (*apigatewaymanagementapi.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		isOffline := os.Getenv("IS_OFFLINE") == "true"

		fmt.Println("==================== | NEW API GATEWAY MANAGEMENT API | IS OFFLINE ====================", isOffline)

		endpoint := os.Getenv("APIG_ENDPOINT")

		if isOffline {
			endpoint = "http://localhost:8082"
		}

		fmt.Println("==================== | NEW API GATEWAY MANAGEMENT API | ENDPOINT ====================", endpoint)

		return aws.Endpoint{
			URL: endpoint,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(customResolver),
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("==================== | NEW API GATEWAY MANAGEMENT API | CONFIG ====================", cfg)

	return apigatewaymanagementapi.NewFromConfig(cfg), nil
}
