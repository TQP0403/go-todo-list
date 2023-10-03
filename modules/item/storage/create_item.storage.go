package storage

import "todo-list/modules/item/dto"

type CreateItemStorage interface {
	CreateItem(data *dto.CreateTodoItemDto) error
}

func (sql *SqlStore) CreateItem(data *dto.CreateTodoItemDto) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
