package product

import "products/model"

type UseCase interface {
	Create(m model.Product) error
	GetByID(ID uint) (model.Product, error)
	GetAll() (model.Products, error)
	GetBoxPriceByID(ID uint, currency string, quantity uint) (float64, error)
}

type Storage interface {
	Create(m model.Product) error
	GetByID(ID uint) (model.Product, error)
	GetAll() (model.Products, error)
}

type ClientConvertCurrency interface {
	Convert(from, to string, amount float64) (float64, error)
}
