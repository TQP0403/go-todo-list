package transport

import (
	"github.com/TQP0403/go-todo-list/modules/common"
	"github.com/TQP0403/go-todo-list/modules/item/business"
	"github.com/TQP0403/go-todo-list/modules/item/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
