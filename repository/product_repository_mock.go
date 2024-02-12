package repository

import (
	"fmt"
	"go-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	arguments := repository.Mock.Called(id)
	fmt.Println(arguments)
	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(entity.Product)
	return &product
}

func (repository *ProductRepositoryMock) FindAll(id string) []entity.Product {
	arguments := repository.Mock.Called(id)
	fmt.Println(arguments)
	if arguments.Get(0) == nil {
		return nil
	}
	products := arguments.Get(0).([]entity.Product)

	return products
}
