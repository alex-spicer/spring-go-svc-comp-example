package service

import (
	"product-go/internal/model"
	"sync"
	"sync/atomic"
)

type ProductService struct {
	products map[int64]model.Product
	mu       sync.RWMutex
	nextID   atomic.Int64
}

func NewProductService() *ProductService {
	return &ProductService{
		products: make(map[int64]model.Product),
	}
}

func (s *ProductService) GetAll() []model.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]model.Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}
	return products
}

func (s *ProductService) Get(id int64) (model.Product, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, exists := s.products[id]
	return p, exists
}

func (s *ProductService) Create(p model.Product) model.Product {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID.Add(1)
	p.ID = id
	s.products[id] = p
	return p
}

func (s *ProductService) Update(id int64, p model.Product) (model.Product, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.products[id]; !exists {
		return model.Product{}, false
	}

	p.ID = id
	s.products[id] = p
	return p, true
}

func (s *ProductService) Delete(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.products[id]; !exists {
		return false
	}

	delete(s.products, id)
	return true
}
