package transport

import (
	"github.com/TQP0403/go-todo-list/modules/common"
	"github.com/TQP0403/go-todo-list/modules/item/business"
	"github.com/TQP0403/go-todo-list/modules/item/dto"
	"github.com/TQP0403/go-todo-list/modules/item/storage"
	"github.com/gin-gonic/gin"
	"net/http"
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

		ctx.JSON(http.StatusCreated, common.GetResponseSuccess(data.Id, nil))
	}
}
