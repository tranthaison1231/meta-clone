package h

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/tranthaison1231/meta-clone/api/models"
)

type S3PresignClient struct {
	Presigner *s3.PresignClient
}

func (s3PresignClient *S3PresignClient) GetPresignedUrl(ctx *gin.Context, input models.GetPresignedUrlRequestDto) (string, error) {
	bucketName := os.Getenv("S3_BUCKET_NAME")

	request, err := s3PresignClient.Presigner.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(input.FileName),
	})

	if err != nil {
		return "", err
	}

	return request.URL, err
}

func InitS3PresignClient(ctx *gin.Context) (*S3PresignClient, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("S3_BUCKET_REGION")))

	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(
		sdkConfig,
	)
	presigner := s3.NewPresignClient(s3Client)

	s3PresignClient := S3PresignClient{
		Presigner: presigner,
	}

	return &s3PresignClient, nil
}
