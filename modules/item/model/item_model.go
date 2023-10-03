package model

import (
	"errors"
	"github.com/TQP0403/go-todo-list/modules/common"
	"time"
)

var (
	ErrEmptyTittle = common.NewBadRequestError(errors.New("tittle is empty"), "tittle is empty")
)

type TodoItem struct {
	Id          uint       `json:"id" gorm:"column:id;primaryKey"`
	Tittle      string     `json:"tittle" gorm:"column:tittle"`
	Description string     `json:"description" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status;default:0"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (item *TodoItem) TableName() string {
	return "item"
}
