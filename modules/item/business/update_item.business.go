package business

import (
	"context"
	"github.com/TQP0403/go-todo-list/modules/item/dto"
	"github.com/TQP0403/go-todo-list/modules/item/storage"
)

type UpdateItemBusiness struct {
	store *storage.SqlStore
}

func NewUpdateItemBusiness(store *storage.SqlStore) *UpdateItemBusiness {
	return &UpdateItemBusiness{store: store}
}

func (b UpdateItemBusiness) Execute(ctx context.Context, id int, data *dto.UpdateTodoItemDto) error {
	if _, err := b.store.GetItem(id); err != nil {
		return err
	}

	if err := b.store.UpdateItem(id, data); err != nil {
		return err
	}

	return nil
}
