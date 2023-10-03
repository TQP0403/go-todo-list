package storage

import (
	"errors"
	"github.com/TQP0403/go-todo-list/modules/common"
	"github.com/TQP0403/go-todo-list/modules/item/model"
	"gorm.io/gorm"
)

type GetItemStorage interface {
	GetItem(id int) (data *model.TodoItem, err error)
}

func (sql *SqlStore) GetItem(id int) (data *model.TodoItem, err error) {
	if err := sql.db.Where(id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewNotFoundError(err, "Item not found")
		}
		return nil, err
	}

	return data, nil
}
