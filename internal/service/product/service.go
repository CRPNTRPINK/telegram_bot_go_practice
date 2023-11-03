package product

import (
	"errors"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	idx = idx - 1
	if !(len(allProducts) > idx && idx >= 0) {
		return nil, errors.New("error. The index is less than zero or greater than the list")
	}

	return &allProducts[idx], nil
}
