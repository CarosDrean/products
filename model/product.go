package model

type Product struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type Products []Product

func (ps Products) GetPriceTotal() float64 {
	priceTotal := 0.0
	for _, product := range ps {
		priceTotal += product.Price
	}

	return priceTotal
}
