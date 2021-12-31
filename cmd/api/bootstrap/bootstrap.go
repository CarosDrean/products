package bootstrap

import (
	"fmt"

	"products/infraestructure/handler"
)

func Run() error {
	config := newConfiguration("./configuration.json")
	api := newEcho()
	db := newDB(config)

	handler.InitRoutes(api, db, config)

	port := fmt.Sprintf(":%d", config.Port)
	return api.Start(port)
}
