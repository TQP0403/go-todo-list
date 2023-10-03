package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list/modules/common"
	"todo-list/modules/item/business"
	"todo-list/modules/item/dto"
	"todo-list/modules/item/storage"
)

type CreateItemTransport interface {
	HandleCreateItem() func(ctx *gin.Context)
}

func (t *Transport) HandleCreateItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data dto.CreateTodoItemDto

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.GetResponseFail(err))

			return
		}

		store := storage.NewSqlStore(t.db)
		b := business.NewCreateItemBusiness(store)

		if err := b.Execute(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.GetResponseFail(err))

			return
		}

		ctx.JSON(http.StatusOK, common.GetResponseSuccess(data.Id, nil))
	}
}
