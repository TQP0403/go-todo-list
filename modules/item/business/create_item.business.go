package business

import (
	"context"
	"github.com/TQP0403/go-todo-list/modules/item/dto"
	"github.com/TQP0403/go-todo-list/modules/item/model"
	"github.com/TQP0403/go-todo-list/modules/item/storage"
	"strings"
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
