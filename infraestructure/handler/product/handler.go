package product

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"products/domain/product"
	"products/model"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	useCase product.UseCase
}

func newHandler(useCase product.UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}

func (h Handler) create(c echo.Context) error {
	requestData := model.Product{}
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("c.Bind()"))
	}

	err := h.useCase.Create(requestData)
	if err != nil {
		errData := model.NewError()
		if errors.As(err, &errData) {
			return c.JSON(errData.StatusHTTP(), model.Response{
				Message: errData.APIMessage(),
				Data:    errData.Data(),
			})
		}

		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("created successful"))
}

func (h Handler) getByID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("convert error: %v", err).Error())
	}

	data, err := h.useCase.GetByID(uint(ID))
	if err != nil {
		errData := model.NewError()
		if errors.As(err, &errData) {
			return c.JSON(errData.StatusHTTP(), model.Response{Message: errData.APIMessage()})
		}

		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h Handler) getAll(c echo.Context) error {
	data, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h Handler) getBoxPriceByID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("handler.atoi: %v", err).Error())
	}

	currency := c.FormValue("currency")

	quantity := model.VoidInt
	quantityString := c.FormValue("quantity")

	if quantityString != model.VoidString {
		quantity, err = strconv.Atoi(quantityString)
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("handler.atoi: %v", err).Error())
		}
	}

	data, err := h.useCase.GetBoxPriceByID(uint(ID), currency, uint(quantity))
	if err != nil {
		errData := model.NewError()
		if errors.As(err, &errData) {
			return c.JSON(errData.StatusHTTP(), model.Response{Message: errData.APIMessage()})
		}

		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error: %v", err).Error())
	}

	return c.JSON(http.StatusOK, data)
}
