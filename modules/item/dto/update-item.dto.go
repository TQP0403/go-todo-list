package dto

import (
	"github.com/TQP0403/go-todo-list/modules/item/model"
)

type UpdateTodoItemDto struct {
	Id          uint              `json:"-" gorm:"column:id"`
	Tittle      *string           `json:"tittle" gorm:"column:tittle"`
	Description *string           `json:"description" gorm:"column:description"`
	Status      *model.ItemStatus `json:"status"  gorm:"column:status"`
}

func (item *UpdateTodoItemDto) TableName() string {
	return "item"
}
