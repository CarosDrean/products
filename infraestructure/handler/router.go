package handler

import (
	"products/infraestructure/handler/product"
	"products/model"

	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Echo, db string, config model.Configuration) {
	product.NewRoutes(app, db, config)
}
