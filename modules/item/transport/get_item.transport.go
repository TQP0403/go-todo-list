package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/modules/common"
	"todo-list/modules/item/business"
	"todo-list/modules/item/storage"
)

type GetItemTransport interface {
	HandleGetItem() func(ctx *gin.Context)
}

func (t *Transport) HandleGetItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.GetResponseFail(err))

			return
		}

		store := storage.NewSqlStore(t.db)
		b := business.NewGetItemBusiness(store)

		data, err := b.Execute(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, common.GetResponseFail(err))
			return
		}

		ctx.JSON(http.StatusOK, common.GetResponseSuccess(data, nil))
	}
}
