package product

import (
	"products/model"
)

var memory map[uint]model.Product

type Product struct {
	db string
}

func New(db string) *Product {
	memory = make(map[uint]model.Product)

	return &Product{db: db}
}

func (p Product) Create(m model.Product) error {
	_, exist := memory[m.ID]
	if exist {
		return model.ErrorIDAlreadyExist
	}

	memory[m.ID] = m
	return nil
}

func (p Product) GetByID(ID uint) (model.Product, error) {
	_, exist := memory[ID]
	if !exist {
		return model.Product{}, model.ErrorIDNotFound
	}

	data := memory[ID]
	return data, nil
}

func (p Product) GetAll() (model.Products, error) {
	data := make(model.Products, 0)

	for _, value := range memory {
		data = append(data, value)
	}

	return data, nil
}
