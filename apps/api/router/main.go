package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"api/db"
	"api/handlers"

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

	port := os.Getenv("PORT")
	fmt.Println("PORT:", port)
	r := gin.Default()

	r.Use(cors.Default())

	var authController = handlers.InitAuthController(db.DB)
	r.POST("/sign-in", authController.SignIn)
	r.POST("/sign-up", authController.SignUp)
	r.GET("/me", authController.GetMe)
	r.GET("/me/groups", handlers.GetMyGroups)
	r.GET("/me/contacts", handlers.GetMyContacts)
	r.GET("/contacts/:contactID/messages", handlers.GetMessagesByContact)
	r.POST("/contacts/:contactID/messages", handlers.SendMessageToContact)

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
