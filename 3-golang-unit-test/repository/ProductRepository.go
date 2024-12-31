package repository

import "3-golang-unit-test/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}