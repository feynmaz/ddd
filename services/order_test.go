package services

import (
	"testing"

	"github.com/feynmaz/ddd/aggregate"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	apple, err := aggregate.NewProduct(
		"Apple", "Granny Smith", 0.99,
	)
	if err != nil {
		t.Fatal(err)
	}

	orange, err := aggregate.NewProduct("Orange", "Clementine, Seven Seas", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	banana, err := aggregate.NewProduct("Banana", "Lady Finger", 0.49)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{apple, orange, banana}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}
