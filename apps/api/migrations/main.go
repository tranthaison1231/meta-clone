package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tranthaison1231/meta-clone/api/models"

	"github.com/tranthaison1231/meta-clone/api/db"
)

func init() {
	godotenv.Load()
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(
		&models.User{},
		&models.Community{},
		&models.Chat{},
		&models.Message{},
		&models.FriendRequest{},
		&models.Post{},
	)
	fmt.Println("👍 Migration complete")
}
