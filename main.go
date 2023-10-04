package main

import (
	"fmt"
	"github.com/TQP0403/go-todo-list/db"
	"github.com/TQP0403/go-todo-list/modules/item"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
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
