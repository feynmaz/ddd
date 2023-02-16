package services

import (
	"log"

	"github.com/feynmaz/ddd/domain/customer"
	"github.com/feynmaz/ddd/domain/customer/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository rapplies a customer repository to the OrderService
func WithCustomerRepository(ct customer.CustomerRepository) OrderConfiguration {
	// Return a function that matches the orderConfiguration alias
	return func(os *OrderService) error {
		os.customers = ct
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) error {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}
	log.Println(c)

	return nil
}
