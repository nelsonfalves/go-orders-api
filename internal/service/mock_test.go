package service

import (
	"github.com/nelsonalves117/go-orders-api/internal/canonical"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) GetAllOrders() ([]canonical.Order, error) {
	args := m.Called()
	return args.Get(0).([]canonical.Order), args.Error(1)
}

func (m *mockRepository) GetOrderById(id string) (canonical.Order, error) {
	args := m.Called(id)
	return args.Get(0).(canonical.Order), args.Error(1)
}

func (m *mockRepository) CreateOrder(order canonical.Order) (canonical.Order, error) {
	args := m.Called(order)
	return args.Get(0).(canonical.Order), args.Error(1)
}

func (m *mockRepository) UpdateOrder(id string, order canonical.Order) (canonical.Order, error) {
	args := m.Called(id, order)
	return args.Get(0).(canonical.Order), args.Error(1)
}

func (m *mockRepository) DeleteOrder(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
