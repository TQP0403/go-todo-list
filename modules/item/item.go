package item

import (
	"github.com/TQP0403/go-todo-list/modules/item/transport"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRoute(engine *gin.Engine, db *gorm.DB) {
	group := engine.Group("/item")
	trans := transport.NewTransport(db)

	// method POST: create an item
	group.POST("/", trans.HandleCreateItem())

	// method GET: get list items
	group.GET("", trans.HandleListItems())

	// method GET: get an item by id
	group.GET("/:id", trans.HandleGetItem())

	// method PUT: update an item by id
	group.PUT("/:id", trans.HandleUpdateItem())

	// method DELETE: delete an item by id
	group.DELETE("/:id", trans.HandleDeleteItem())
}
