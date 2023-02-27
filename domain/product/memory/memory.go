package memory

import (
	"fmt"
	"sync"

	"github.com/feynmaz/ddd/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepositoty struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepositoty {
	return &MemoryProductRepositoty{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mpr *MemoryProductRepositoty) GetAll() ([]product.Product, error) {
	products := make([]product.Product, 0, len(mpr.products))

	for _, product := range mpr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mpr *MemoryProductRepositoty) GetByID(id uuid.UUID) (product.Product, error) {
	if p, ok := mpr.products[id]; ok {
		return p, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepositoty) Add(p product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if mpr.products == nil {
		mpr.products = make(map[uuid.UUID]product.Product)
	}

	if _, ok := mpr.products[p.GetID()]; ok {
		return fmt.Errorf("product already exists: %w", product.ErrFailedToAddProduct)
	}

	mpr.products[p.GetID()] = p
	return nil
}

func (mpr *MemoryProductRepositoty) Update(p product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[p.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[p.GetID()] = p
	return nil
}

func (mpr *MemoryProductRepositoty) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
