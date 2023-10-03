package storage

import (
	"errors"
	"gorm.io/gorm"
	"todo-list/modules/common"
	"todo-list/modules/item/model"
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
