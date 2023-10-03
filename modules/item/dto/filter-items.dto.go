package dto

import "todo-list/modules/item/model"

type FilterItemDto struct {
	Status *model.ItemStatus `json:"status" form:"status"`
}
