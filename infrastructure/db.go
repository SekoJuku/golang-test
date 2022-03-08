package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	DBPORT := os.Getenv("DB_PORT")
	URL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty", HOST, USER,
		PASS, DBNAME, DBPORT)
	fmt.Println(URL)
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")

	}
	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}
}
