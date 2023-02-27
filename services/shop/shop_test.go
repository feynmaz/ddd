package shop

import (
	"testing"

	"github.com/feynmaz/shop/domain/product"
	order "github.com/feynmaz/shop/services/order"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	apple, err := product.New(
		"Apple", "Granny Smith", 0.99,
	)
	if err != nil {
		t.Fatal(err)
	}

	orange, err := product.New("Orange", "Clementine, Seven Seas", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	banana, err := product.New("Banana", "Lady Finger", 0.49)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{apple, orange, banana}
}

func Test_Shop(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	shop, err := NewShop(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	custID, err := os.AddCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = shop.Order(custID, order)
	if err != nil {
		t.Fatal(err)
	}
}
