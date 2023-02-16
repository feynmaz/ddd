package customer

import (
	"errors"

	"github.com/feynmaz/ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrUpdateCustomer      = errors.New("failed update customer")
)

type CustomerRepository interface {
	Get(id uuid.UUID) (aggregate.Customer, error)
	Add(customer aggregate.Customer) error
	Update(aggregate.Customer) error
}

