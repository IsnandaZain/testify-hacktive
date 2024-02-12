package service

import (
	"fmt"
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be product not found")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Kaca mata",
	}

	productRepository.Mock.On("FindById", "2").Return(product)
	result, err := productService.GetOneProduct("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'Kaca mata'")
	assert.Equal(t, &product, result, "result has to be a product data with id '2'")
}

// All Product
func TestProductServiceGetAllProduct(t *testing.T) {
	products := []entity.Product{
		{Id: "1", Name: "Kaca mata"},
		{Id: "2", Name: "Laptop"},
	}

	t.Run("GetAll", func(t *testing.T) {
		productRepository.Mock.On("FindAll", "2").Return(products)
		result, err := productService.GetAllProduct("2")

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, products, result, "result has to be products")
		assert.Equal(t, len(products), len(result), "number of product returned should match")

	})

	t.Run("GetAllNotFound", func(t *testing.T) {

		productRepository.Mock.On("FindAll", "1").Return(nil)
		resultNotFound, err := productService.GetAllProduct("1")
		fmt.Println(resultNotFound)

		assert.Nil(t, resultNotFound)
		assert.NotNil(t, err)
		assert.Equal(t, "Product not found", err.Error(), "error response has to be 'product not found'")
	})
}
