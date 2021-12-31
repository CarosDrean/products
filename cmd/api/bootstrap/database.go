package bootstrap

import "products/model"

func newDB(config model.Configuration) string {
	return config.Database.Name
}
