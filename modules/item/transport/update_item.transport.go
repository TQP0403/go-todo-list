package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/modules/common"
	"todo-list/modules/item/business"
	"todo-list/modules/item/dto"
	"todo-list/modules/item/storage"
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
