package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrUpdateCustomer      = errors.New("failed update customer")
)

type Repository interface {
	Get(id uuid.UUID) (Customer, error)
	Add(customer Customer) error
	Update(Customer) error
}
