package product

import (
	"products/domain/product"
	"products/infraestructure/client/currencylayer"
	productStorage "products/infraestructure/memorydb/product"
	"products/model"

	"github.com/labstack/echo/v4"
)

const (
	privateRoutePrefix = "beers"
)

func NewRoutes(app *echo.Echo, db string, config model.Configuration) {
	useCase := product.New(productStorage.New(db), currencylayer.New(config.ApiKeyCurrencyLayer))

	handler := newHandler(useCase)

	privateRoutes(app, handler)
}

func privateRoutes(app *echo.Echo, handler Handler) {
	api := app.Group(privateRoutePrefix)

	api.POST("", handler.create)
	api.GET("/:id", handler.getByID)
	api.GET("", handler.getAll)
	api.GET("/:id/boxprice", handler.getBoxPriceByID)
}
