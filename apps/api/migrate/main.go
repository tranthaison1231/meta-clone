package main

import (
	"api/db"
	"api/model"
	"fmt"
)

func init() {
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&model.User{})
	fmt.Println("👍 Migration complete")
}
