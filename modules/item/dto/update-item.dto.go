package dto

import (
	"github.com/TQP0403/go-todo-list/modules/item/model"
	"time"
)

type UpdateTodoItemDto struct {
	Id          uint              `json:"-" gorm:"column:id"`
	Tittle      *string           `json:"tittle" gorm:"column:tittle"`
	Description *string           `json:"description" gorm:"column:description"`
	Status      *model.ItemStatus `json:"status"  gorm:"column:status"`
	UpdatedAt   time.Time         `json:"-" gorm:"column:updated_at"`
}

func (item *UpdateTodoItemDto) TableName() string {
	return "item"
}
