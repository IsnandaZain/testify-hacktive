package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Struct untuk data produk
type Product struct {
	ID   string
	Name string
}

// Interface untuk repository produk
type ProductRepository interface {
	FindAll() []Product
}

// Mock repository produk
type MockProductRepository struct {
	mock.Mock
}

// Implementasi metode FindAll untuk mock repository
func (m *MockProductRepository) FindAll() []Product {
	// Mendefinisikan argumen yang akan dikembalikan saat dipanggil
	args := m.Called()
	return args.Get(0).([]Product)
}

// Fungsi yang akan diuji
func ProcessProducts(repository ProductRepository) int {
	products := repository.FindAll()
	return len(products)
}

func TestProcessProducts(t *testing.T) {
	// Membuat slice produk untuk pengujian
	testProducts := []Product{
		{ID: "1", Name: "Kaca mata"},
		{ID: "2", Name: "Laptop"},
		{ID: "3", Name: "Keyboard"},
	}

	// Membuat mock repository
	mockRepo := new(MockProductRepository)
	// Mendefinisikan perilaku mock untuk metode FindAll
	mockRepo.On("FindAll").Return(testProducts)

	// Memanggil fungsi yang akan diuji dengan mock repository sebagai argumen
	result := ProcessProducts(mockRepo)

	// Memeriksa apakah hasilnya sesuai dengan yang diharapkan (panjang slice produk)
	assert.Equal(t, len(testProducts), result)

	// Memeriksa apakah metode FindAll pada mock repository dipanggil
	mockRepo.AssertCalled(t, "FindAll")
}
