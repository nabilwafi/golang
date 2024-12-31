package service

import (
	"3-golang-unit-test/entity"
	"3-golang-unit-test/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (s ProductService) Get(id string) (*entity.Product, error) {
	category := s.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Product not found");
	}

	return category, nil
}