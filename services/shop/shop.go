package shop

import (
	"log"

	"github.com/feynmaz/shop/services/order"
	"github.com/google/uuid"
)

type ShopConfiguration func(s *Shop) error

type Shop struct {
	OrderService   *order.OrderService
	BillingService interface{}
}

func NewShop(cfgs ...ShopConfiguration) (*Shop, error) {
	s := &Shop{}

	for _, cfg := range cfgs {
		if err := cfg(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func WithOrderService(os *order.OrderService) ShopConfiguration {
	return func(s *Shop) error {
		s.OrderService = os
		return nil
	}
}

func (s *Shop) Order(customer uuid.UUID, products []uuid.UUID) (float64, error) {
	total, err := s.OrderService.CreateOrder(customer, products)
	if err != nil {
		return 0, err
	}

	log.Printf("Bill the customer: %.2f", total)
	return total, nil
}
