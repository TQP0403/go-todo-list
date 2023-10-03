package dto

import (
	"github.com/TQP0403/go-todo-list/modules/item/model"
)

type CreateTodoItemDto struct {
	Id          uint             `json:"-" gorm:"column:id"`
	Tittle      string           `json:"tittle" gorm:"column:tittle"`
	Description string           `json:"description" gorm:"column:description"`
	Status      model.ItemStatus `json:"-"  gorm:"column:status"`
}

func (item *CreateTodoItemDto) TableName() string {
	return "item"
}
