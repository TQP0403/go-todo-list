package business

import (
	"context"
	"todo-list/modules/item/model"
	"todo-list/modules/item/storage"
)

type GetItemBusiness struct {
	store *storage.SqlStore
}

func NewGetItemBusiness(store *storage.SqlStore) *GetItemBusiness {
	return &GetItemBusiness{store: store}
}

func (b GetItemBusiness) Execute(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := b.store.GetItem(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
