package transport

import (
	"github.com/TQP0403/go-todo-list/modules/common"
	"github.com/TQP0403/go-todo-list/modules/item/business"
	"github.com/TQP0403/go-todo-list/modules/item/dto"
	"github.com/TQP0403/go-todo-list/modules/item/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateItemTransport interface {
	HandleUpdateItem() func(ctx *gin.Context)
}

func (t *Transport) HandleUpdateItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data dto.UpdateTodoItemDto

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.GetResponseFail(err))

			return
		}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.GetResponseFail(err))

			return
		}

		store := storage.NewSqlStore(t.db)
		b := business.NewUpdateItemBusiness(store)

		if err := b.Execute(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.GetResponseFail(err))

			return
		}

		ctx.JSON(http.StatusOK, common.GetResponseSuccess(nil, nil))
	}
}
