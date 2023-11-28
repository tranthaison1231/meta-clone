package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tranthaison1231/meta-clone/api/conf"
	"github.com/tranthaison1231/meta-clone/api/handlers"
	"github.com/tranthaison1231/meta-clone/api/services"

	"github.com/tranthaison1231/meta-clone/api/middlewares"

	"github.com/tranthaison1231/meta-clone/api/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	ginLambda *ginadapter.GinLambda
	server    *http.Server
	port      string
)

func initRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/sign-in", handlers.SignIn)
	r.POST("/sign-up", handlers.SignUp)

	auth := r.Group("", middlewares.Auth)

	auth.GET("/me", handlers.GetMe)
	auth.PUT("/me", handlers.UpdateMe)
	auth.GET("/chats", handlers.GetChats)
	auth.POST("/chats", handlers.CreateChat)
	auth.POST("/news-feed", handlers.GetNewsFeed)
	auth.GET("/posts", handlers.GetPosts)
	auth.POST("/posts", handlers.CreatePost)
	auth.POST("/chats/:chatID/join", handlers.AddMemberToChat)
	auth.GET("/communities", handlers.GetCommunities)
	auth.POST("/communities", handlers.CreateCommunity)
	auth.GET("/chats/:chatID/messages", handlers.GetMessages)
	auth.POST("/chats/:chatID/messages", handlers.SendMessage)
	auth.POST("/users/add-friend", handlers.AddFriend)
	auth.POST("/users/accept-friend", handlers.AcceptFriend)
	auth.GET("/users/:userID/friends", handlers.GetUserFriends)
	auth.GET("/users", handlers.GetUsers)
}

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

	initRoutes(r)

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

// @title Meta-Clone
// @version 1.0
// @description API for Meta-Clone
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if port != "" {
		server.ListenAndServe()
	} else {
		lambda.Start(Handler)
	}
}
