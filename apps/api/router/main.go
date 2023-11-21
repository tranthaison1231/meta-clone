package main

import (
	"context"
	"log"
	"net/http"
	"os"

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
	db.ConnectDB()
	conf.Conf = conf.DefaultConfig()
	services.SecretKey = []byte(conf.Conf.JwtSecret)

	port := os.Getenv("PORT")
	r := gin.Default()

	r.Use(cors.Default())

	var authController = handlers.InitAuthController(db.DB)
	r.POST("/sign-in", authController.SignIn)
	r.POST("/sign-up", authController.SignUp)

	auth := r.Group("", middlewares.Auth)

	auth.GET("/me", authController.GetMe)
	auth.GET("/me/groups", handlers.GetMyGroups)
	auth.GET("/me/contacts", handlers.GetMyContacts)
	auth.GET("/contacts/:contactID/messages", handlers.GetMessagesByContact)
	auth.POST("/contacts/:contactID/messages", handlers.SendMessageToContact)

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
