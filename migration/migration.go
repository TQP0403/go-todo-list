package main

import (
	"fmt"
	"github.com/TQP0403/go-todo-list/db"
	"github.com/TQP0403/go-todo-list/modules/item/model"
)

func main() {
	migrator := db.GetDB().Migrator()

	if !migrator.HasTable(&model.TodoItem{}) {
		if err := migrator.CreateTable(&model.TodoItem{}); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("database migrated")
}
