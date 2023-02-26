package services

import (
	"log"

	"github.com/feynmaz/ddd/aggregate"
	"github.com/feynmaz/ddd/domain/customer"
	"github.com/feynmaz/ddd/domain/customer/memory"
	"github.com/feynmaz/ddd/domain/product"
	prodmem "github.com/feynmaz/ddd/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
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

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	products := make([]aggregate.Product, 0, len(productsIDs))
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("customer %s has ordered %d products with total %.2f\n", c.GetID(), len(products), total)
	return total, nil
}
