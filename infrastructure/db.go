package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Database struct
type Database struct {
	DB *gorm.DB
}

//NewDatabase : intializes and returns mysql db
func NewDatabase() Database {
	USER := "user"
	PASS := "123"
	HOST := "localhost"
	DBNAME := "golang_task"

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)
	fmt.Println(URL)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")

	}
	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
