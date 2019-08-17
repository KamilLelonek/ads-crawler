package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "crawler_development"
)

var psqlInfo = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname,
)

func DB() *gorm.DB {
	return db
}

func SetupDB() *gorm.DB {
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("db err: ", err)
	}

	return db
}
