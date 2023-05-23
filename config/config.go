package config

import (
	"fmt"
	"os"
	"skyshi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Config() *gorm.DB {
	var err error

	DBHost := os.Getenv("MYSQL_HOST")
	DBPort := os.Getenv("MYSQL_PORT")
	DBName := os.Getenv("MYSQL_DBNAME")
	DBUser := os.Getenv("MYSQL_USER")
	DBpass := os.Getenv("MYSQL_PASSWORD")

	dsn := DBUser + ":" + DBpass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8&parseTime=true&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("ini dsn ==>>>  ", dsn)

	if err != nil {
		panic("failed to connect database !")
	}

	DB.AutoMigrate(&models.Todo{}, &models.Activity{})

	return DB
}
