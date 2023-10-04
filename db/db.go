package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func loadFileDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("env file not found")
	}
}

func GetDB() *gorm.DB {
	loadFileDotEnv()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	connStr, ok := os.LookupEnv("DB_CONNECTION")

	//connectionStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)

	connectionStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, password, name, port)

	fmt.Printf("connection string: %v \n", connectionStr)

	if ok {
		connectionStr = connStr
	}

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println("connect failed")
		panic(err)
	}

	fmt.Println("connect succeed")

	return db
}
