package memory

import (
	"fmt"
	"sync"

	"github.com/feynmaz/ddd/aggregate"
	domain "github.com/feynmaz/ddd/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepositoty struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepositoty {
	return &MemoryProductRepositoty{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepositoty) GetAll() ([]aggregate.Product, error) {
	products := make([]aggregate.Product, 0, len(mpr.products))

	for _, product := range mpr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mpr *MemoryProductRepositoty) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return aggregate.Product{}, domain.ErrProductNotFound
}

func (mpr *MemoryProductRepositoty) Add(product aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if mpr.products == nil {
		mpr.products = make(map[uuid.UUID]aggregate.Product)
	}

	if _, ok := mpr.products[product.GetID()]; ok {
		return fmt.Errorf("product already exists: %w", domain.ErrFailedToAddProduct)
	}

	mpr.products[product.GetID()] = product
	return nil
}

func (mpr *MemoryProductRepositoty) Update(product aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[product.GetID()]; !ok {
		return domain.ErrProductNotFound
	}

	mpr.products[product.GetID()] = product
	return nil
}

func (mpr *MemoryProductRepositoty) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return domain.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
