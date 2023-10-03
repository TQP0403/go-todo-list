package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"todo-list/db"
	"todo-list/modules/item"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(".env file not found")
	}

	myDb := db.GetDB()

	app := gin.Default()

	item.AddRoute(app, myDb)

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	addr := fmt.Sprintf(":%v", os.Getenv("PORT"))

	app.Run(addr)
}
