package storage

import "github.com/TQP0403/go-todo-list/modules/item/model"

type DeleteItemStorage interface {
	DeleteItem(id int) error
}

func (sql *SqlStore) DeleteItem(id int) error {
	if err := sql.db.Where(id).Delete(model.TodoItem{}).Error; err != nil {
		return err
	}

	return nil
}
