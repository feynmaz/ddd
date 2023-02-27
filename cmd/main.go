package main

import (
	"log"

	"github.com/google/uuid"

	"github.com/feynmaz/shop/domain/product"
	"github.com/feynmaz/shop/services/order"
	"github.com/feynmaz/shop/services/shop"
)

func main() {
	products := productInventory()
	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	business, err := shop.NewShop(
		shop.WithOrderService(os),
	)
	if err != nil {
		panic(err)
	}

	custID, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
	}

	total, err := business.Order(custID, order)
	if err != nil {
		panic(err)
	}

	log.Print(total)
}

func productInventory() []product.Product {
	apple, err := product.New(
		"Apple", "Granny Smith", 0.99,
	)
	if err != nil {
		panic(err)
	}

	orange, err := product.New("Orange", "Clementine, Seven Seas", 1.99)
	if err != nil {
		panic(err)
	}

	banana, err := product.New("Banana", "Lady Finger", 0.49)
	if err != nil {
		panic(err)
	}

	return []product.Product{apple, orange, banana}
}
