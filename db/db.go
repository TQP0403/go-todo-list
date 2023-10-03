package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDB() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	schema := os.Getenv("DB_SCHEMA")
	connStr, ok := os.LookupEnv("DB_CONNECTION")

	connectionStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, schema)

	if ok {
		connectionStr = connStr
	}

	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})

	if err != nil {
		fmt.Println("connect failed")
		panic(err)
	}

	fmt.Println("connect succeed")

	return db
}
