package dto

import "github.com/TQP0403/go-todo-list/modules/item/model"

type FilterItemDto struct {
	Status *model.ItemStatus `json:"status" form:"status"`
}
