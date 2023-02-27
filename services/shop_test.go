package services

import (
	"testing"

	"github.com/feynmaz/ddd/domain/customer"
	"github.com/google/uuid"
)

func Test_Shop(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	shop, err := NewShop(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	customer, err := customer.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.customers.Add(customer); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = shop.Order(customer.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}
}
