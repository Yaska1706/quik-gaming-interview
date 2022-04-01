package repository

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yaska1706/quik-gaming-interview/pkg/api"
)

var (
	err error
)

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() (*gorm.DB, error) {

	var db *gorm.DB

	driver := os.Getenv("DB_DRIVER")
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	db, err = gorm.Open(driver, username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}

	db.AutoMigrate(api.Wallet{})
	return db, nil
}
