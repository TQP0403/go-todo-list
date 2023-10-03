package storage

import (
	"github.com/TQP0403/go-todo-list/modules/common"
	"github.com/TQP0403/go-todo-list/modules/item/dto"
	"github.com/TQP0403/go-todo-list/modules/item/model"
)

type ListItemStorage interface {
	ListItem(filter *dto.FilterItemDto, p *common.Pagination) (*[]model.TodoItem, error)
}

func (sql *SqlStore) ListItem(filter *dto.FilterItemDto, p *common.Pagination) (*[]model.TodoItem, error) {
	var data []model.TodoItem

	query := sql.db.Model(model.TodoItem{})

	if filter != nil {
		if filter.Status != nil && filter.Status.String() != "" {
			query = query.Where("status = ?", filter.Status)
		}
	}

	if err := query.Count(&p.Total).Offset((p.Page - 1) * p.Limit).Limit(p.Limit).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
