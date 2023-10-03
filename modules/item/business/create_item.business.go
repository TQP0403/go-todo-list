package business

import (
	"context"
	"strings"
	"todo-list/modules/item/dto"
	"todo-list/modules/item/model"
	"todo-list/modules/item/storage"
)

type CreateItemBusiness struct {
	store *storage.SqlStore
}

func NewCreateItemBusiness(store *storage.SqlStore) *CreateItemBusiness {
	return &CreateItemBusiness{store: store}
}

func (b CreateItemBusiness) Execute(ctx context.Context, data *dto.CreateTodoItemDto) error {
	tittle := strings.TrimSpace(data.Tittle)

	if tittle == "" {
		return model.ErrEmptyTittle
	}

	if err := b.store.CreateItem(data); err != nil {
		return err
	}

	return nil
}
