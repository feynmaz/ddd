package order

import (
	"testing"

	"github.com/feynmaz/shop/domain/customer"
	"github.com/feynmaz/shop/domain/product"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	apple, err := product.NewProduct(
		"Apple", "Granny Smith", 0.99,
	)
	if err != nil {
		t.Fatal(err)
	}

	orange, err := product.NewProduct("Orange", "Clementine, Seven Seas", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	banana, err := product.NewProduct("Banana", "Lady Finger", 0.49)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{apple, orange, banana}
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

	cust, err := customer.NewCustomer("Percy")
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
