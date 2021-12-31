package product

import (
	"errors"
	"fmt"
	"net/http"

	"products/model"
)

const defaultQuantity = 6

type Product struct {
	storage Storage

	clientCurrency ClientConvertCurrency
}

func New(storage Storage, clientCurrency ClientConvertCurrency) *Product {
	return &Product{
		storage:        storage,
		clientCurrency: clientCurrency,
	}
}

func (p Product) Create(m model.Product) error {
	err := p.storage.Create(m)
	if errors.Is(err, model.ErrorIDAlreadyExist) {
		customErr := model.NewError()
		customErr.SetError(err)
		customErr.SetData(m)
		customErr.SetAPIMessage("El id enviado ya existe en la base de datos")
		customErr.SetStatusHTTP(http.StatusConflict)

		return customErr
	}
	if err != nil {
		return fmt.Errorf("products.storage.Create(): %v", err)
	}

	return nil
}

func (p Product) GetByID(ID uint) (model.Product, error) {
	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrorIDNotFound) {
		customErr := model.NewError()
		customErr.SetError(err)
		customErr.SetAPIMessage("El id enviado no existe en la base de datos")
		customErr.SetStatusHTTP(http.StatusNotFound)

		return model.Product{}, customErr
	}
	if err != nil {
		return model.Product{}, fmt.Errorf("products.storage.GetByID(): %v", err)
	}

	return data, nil
}

func (p Product) GetAll() (model.Products, error) {
	data, err := p.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("products.storage.GetAll(): %v", err)
	}

	return data, nil
}

func (p Product) GetBoxPriceByID(ID uint, currency string, quantity uint) (float64, error) {
	data, err := p.GetByID(ID)
	if err != nil {
		return 0, err
	}

	if quantity == model.VoidInt {
		quantity = defaultQuantity
	}

	if currency != model.VoidString {
		price, err := p.clientCurrency.Convert(data.Currency, currency, data.Price)
		if err != nil {
			return 0, fmt.Errorf("produc.clientCurrency.Convert(): %v", err)
		}

		data.Price = price
	}

	return float64(quantity) * data.Price, nil
}
