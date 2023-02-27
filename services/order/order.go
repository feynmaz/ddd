package order

import (
	"context"
	"log"

	"github.com/feynmaz/shop/domain/customer"
	custmem "github.com/feynmaz/shop/domain/customer/memory"
	"github.com/feynmaz/shop/domain/customer/mongo"
	"github.com/feynmaz/shop/domain/product"
	prodmem "github.com/feynmaz/shop/domain/product/memory"
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
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Return a function that matches the orderConfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := custmem.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connStr string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connStr)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
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

	products := make([]product.Product, 0, len(productsIDs))
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

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	if err := o.customers.Add(c); err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
