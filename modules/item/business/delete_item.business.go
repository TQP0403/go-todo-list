package business

import (
	"context"
	"github.com/TQP0403/go-todo-list/modules/item/storage"
)

type DeleteItemBusiness struct {
	store *storage.SqlStore
}

func NewDeleteItemBusiness(store *storage.SqlStore) *DeleteItemBusiness {
	return &DeleteItemBusiness{store: store}
}

func (b DeleteItemBusiness) Execute(ctx context.Context, id int) error {
	if _, err := b.store.GetItem(id); err != nil {
		return err
	}

	if err := b.store.DeleteItem(id); err != nil {
		return err
	}

	return nil
}
