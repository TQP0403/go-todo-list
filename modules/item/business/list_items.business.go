package business

import (
	"context"
	"github.com/TQP0403/go-todo-list/modules/common"
	"github.com/TQP0403/go-todo-list/modules/item/dto"
	"github.com/TQP0403/go-todo-list/modules/item/model"
	"github.com/TQP0403/go-todo-list/modules/item/storage"
)

type ListItemsBusiness struct {
	store *storage.SqlStore
}

func NewListItemsBusiness(store *storage.SqlStore) *ListItemsBusiness {
	return &ListItemsBusiness{store: store}
}

func (b ListItemsBusiness) Execute(ctx context.Context, f *dto.FilterItemDto, p *common.Pagination) (*[]model.TodoItem, error) {
	data, err := b.store.ListItem(f, p)
	if err != nil {
		return nil, err
	}

	return data, nil
}
