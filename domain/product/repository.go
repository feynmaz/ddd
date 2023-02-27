package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound    = errors.New("product not found")
	ErrFailedToAddProduct = errors.New("failed to add product")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(p Product) error
	Update(p Product) error
	Delete(id uuid.UUID) error
}
