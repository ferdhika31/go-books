package db

import (
	"fmt"

	"github.com/ferdhika31/go-books/config"
	"github.com/ferdhika31/go-books/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	configuration, _ := config.LoadConfig()
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB.User, configuration.DB.Pass, configuration.DB.Name)
	db, err = gorm.Open("mysql", connect_string)
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&models.Category{}, &models.Book{})
}

func DbManager() *gorm.DB {
	return db
}
