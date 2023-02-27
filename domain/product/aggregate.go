package product

import (
	"errors"

	"github.com/feynmaz/shop"
	"github.com/google/uuid"
)

var (
	ErrMissingValue = errors.New("missing important values")
)

type Product struct {
	item     *shop.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}

	return Product{
		item: &shop.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *shop.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
