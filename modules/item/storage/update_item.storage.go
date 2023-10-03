package storage

import "todo-list/modules/item/dto"

type UpdateStorage interface {
	UpdateItem(id int, data *dto.UpdateTodoItemDto) error
}

func (sql *SqlStore) UpdateItem(id int, data *dto.UpdateTodoItemDto) error {
	if err := sql.db.Where(id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
