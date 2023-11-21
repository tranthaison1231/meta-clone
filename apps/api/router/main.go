package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tranthaison1231/messenger-clone/api/conf"
	"github.com/tranthaison1231/messenger-clone/api/handlers"
	"github.com/tranthaison1231/messenger-clone/api/services"

	"github.com/tranthaison1231/messenger-clone/api/middlewares"

	"github.com/tranthaison1231/messenger-clone/api/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	ginLambda *ginadapter.GinLambda
	server    *http.Server
	port      string
)

func init() {
	godotenv.Load()

	db.ConnectDB()
	conf.Conf = conf.DefaultConfig()
	services.SecretKey = []byte(conf.Conf.JwtSecret)

	port := os.Getenv("PORT")
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	r.POST("/sign-in", handlers.SignIn)
	r.POST("/sign-up", handlers.SignUp)

	auth := r.Group("", middlewares.Auth)

	auth.GET("/me", handlers.GetMe)
	auth.GET("/chats", handlers.GetChats)
	auth.POST("/chats", handlers.CreateChat)
	auth.POST("/chats/:chatID/join", handlers.AddMemberToChat)
	auth.GET("/communities", handlers.GetCommunities)
	auth.POST("/communities", handlers.CreateCommunity)
	auth.GET("/chats/:chatID/messages", handlers.GetMessages)
	auth.POST("/chats/:chatID/messages", handlers.SendMessage)

	if port != "" {
		server = &http.Server{
			Addr:    ":" + port,
			Handler: r,
		}
		log.Fatal(server.ListenAndServe())
	} else {
		ginLambda = ginadapter.New(r)
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if port != "" {
		server.ListenAndServe()
	} else {
		lambda.Start(Handler)
	}
}
