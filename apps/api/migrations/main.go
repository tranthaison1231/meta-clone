package main

import (
	"fmt"

	"github.com/tranthaison1231/messenger-clone/api/models"

	"github.com/tranthaison1231/messenger-clone/api/db"
)

func init() {
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.User{}, &models.Chat{}, &models.Message{})
	fmt.Println("👍 Migration complete")
}
