package storage

import (
	"github.com/TQP0403/go-todo-list/modules/item/dto"
	"time"
)

type UpdateStorage interface {
	UpdateItem(id int, data *dto.UpdateTodoItemDto) error
}

func (sql *SqlStore) UpdateItem(id int, data *dto.UpdateTodoItemDto) error {
	data.UpdatedAt = time.Now()
	if err := sql.db.Where(id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
