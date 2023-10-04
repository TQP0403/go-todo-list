package transport

import (
	"github.com/TQP0403/go-todo-list/modules/common"
	"github.com/TQP0403/go-todo-list/modules/item/business"
	"github.com/TQP0403/go-todo-list/modules/item/dto"
	"github.com/TQP0403/go-todo-list/modules/item/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListItemsTransport interface {
	HandleListItems() func(ctx *gin.Context)
}

func (t *Transport) HandleListItems() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Pagination
		var filter dto.FilterItemDto

		if err := ctx.ShouldBindQuery(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.GetResponseFail(err))

			return
		}

		if err := ctx.ShouldBindQuery(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.GetResponseFail(err))

			return
		}

		paging.Validate()

		store := storage.NewSqlStore(t.db)
		b := business.NewListItemsBusiness(store)

		data, err := b.Execute(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusNotFound, common.GetResponseFail(err))
			return
		}

		meta := map[string]any{
			"pagination": paging,
		}

		if filter.Status != nil {
			meta = map[string]any{
				"pagination": paging,
				"filter":     filter,
			}
		}

		ctx.JSON(http.StatusOK, common.GetResponseSuccess(data, meta))
	}
}
