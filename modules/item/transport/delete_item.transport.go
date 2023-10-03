package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/modules/common"
	"todo-list/modules/item/business"
	"todo-list/modules/item/storage"
)

type DeleteItemTransport interface {
	HandleDeleteItem() func(ctx *gin.Context)
}

func (t *Transport) HandleDeleteItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.GetResponseFail(err))

			return
		}

		store := storage.NewSqlStore(t.db)
		b := business.NewDeleteItemBusiness(store)

		if err := b.Execute(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusInternalServerError, common.GetResponseFail(err))

			return
		}

		ctx.JSON(http.StatusOK, common.GetResponseSuccess(nil, nil))
	}
}
