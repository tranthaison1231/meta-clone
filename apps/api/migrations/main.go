package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tranthaison1231/messenger-clone/api/models"

	"github.com/tranthaison1231/messenger-clone/api/db"
)

func init() {
	godotenv.Load()
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.User{}, &models.Chat{}, &models.Message{}, &models.FriendRequest{}, &models.Community{})
	fmt.Println("👍 Migration complete")
}
