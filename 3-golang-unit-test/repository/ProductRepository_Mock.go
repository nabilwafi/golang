package repository

import (
	"3-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}

	category := args.Get(0).(entity.Product)
	return &category
}