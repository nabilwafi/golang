package service

import (
	"3-golang-unit-test/entity"
	"3-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ProductRepositoryMock = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var ProductServiceMock = ProductService{Repository: ProductRepositoryMock}

func TestProductService_GetNotFound(t *testing.T) {
	ProductRepositoryMock.Mock.On("FindById", "1").Return(nil)

	product, err := ProductServiceMock.Get("1")
	
	assert.Nil(t, product, "product must be nil")
	assert.NotNil(t, err, "err must not be nil")
}

func TestProductService_Success(t *testing.T) {
	product := entity.Product{
			Id: "1",
			Name: "Laptop RGB",
	}

	ProductRepositoryMock.Mock.On("FindById", "2").Return(product)

	res, err := ProductServiceMock.Get("2")
	
	assert.Nil(t, err, "err must be nil")
	assert.NotNil(t, res, "res must not be nil")
	assert.Equal(t, product.Id, res.Id, "return Id must be '1'")
	assert.Equal(t, product.Name, res.Name, "return Name must be 'Laptop RGB'")
}