package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("%s&parseTime=True", os.Getenv("DSN"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError:                           true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
