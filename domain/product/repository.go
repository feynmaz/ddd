package product

import (
	"errors"

	"github.com/feynmaz/ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrFailedToAddProduct = errors.New("failed to add product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
