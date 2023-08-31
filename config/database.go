package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func InitialDB() Database {
	DBUser := os.Getenv("MYSQL_DB_USER")
	DBPass := os.Getenv("MYSQL_DB_PASS")
	DBHost := os.Getenv("MYSQL_DB_HOST")
	DBPort := os.Getenv("MYSQL_DB_PORT")
	DBName := os.Getenv("MYSQL_DB_NAME")

	MYSQL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPass, DBHost, DBPort, DBName)
	dsn := MYSQL

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return Database{
		DB: db,
	}
}
