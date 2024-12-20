package service

import (
	"errors"
	"testing"
	"time"

	"github.com/nelsonalves117/go-orders-api/internal/canonical"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllOrders_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	orderTest := []canonical.Order{
		{
			Id: "xpto",
			Products: []string{
				"product1", "product2", "product3",
			},
			Total:     100,
			Status:    "ready",
			CreatedAt: time.Now(),
		},
	}

	mockRepo.On("GetAllOrders").Return(orderTest, nil)

	service := &service{
		repo: mockRepo,
	}

	orders, err := service.GetAllOrders()

	assert.Nil(t, err)
	assert.Equal(t, "xpto", orders[0].Id)
	assert.Equal(t, float32(100), orders[0].Total)
	assert.Equal(t, "ready", orders[0].Status)
	assert.True(t, orders[0].CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))

	mockRepo.AssertExpectations(t)
}

func TestGetAllOrders_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetAllOrders").Return([]canonical.Order{}, errors.New("error occurred when trying to get all orders"))

	service := &service{
		repo: mockRepo,
	}

	orders, err := service.GetAllOrders()

	assert.NotNil(t, err)
	assert.Empty(t, orders)
	assert.Equal(t, "error occurred when trying to get all orders", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetOrderById_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	orderTest := canonical.Order{
		Id: "xpto",
		Products: []string{
			"product1", "product2", "product3",
		},
		Total:     100,
		Status:    "ready",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetOrderById", "xpto").Return(orderTest, nil)

	service := &service{
		repo: mockRepo,
	}

	order, err := service.GetOrderById("xpto")

	assert.Nil(t, err)
	assert.Equal(t, "xpto", order.Id)
	assert.Equal(t, "product1", order.Products[0])
	assert.Equal(t, "product2", order.Products[1])
	assert.Equal(t, "product3", order.Products[2])
	assert.Equal(t, float32(100), order.Total)
	assert.Equal(t, "ready", order.Status)
	assert.True(t, order.CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))
	mockRepo.AssertExpectations(t)
}

func TestGetOrderById_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetOrderById", "xpto").Return(canonical.Order{}, errors.New("error occurred when trying to get an order"))

	service := &service{
		repo: mockRepo,
	}

	order, err := service.GetOrderById("xpto")

	assert.NotNil(t, err)
	assert.Empty(t, order)
	assert.Equal(t, "error occurred when trying to get an order", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestCreateOrder_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	orderTest := canonical.Order{
		Products: []string{
			"product1", "product2", "product3",
		},
		Total:  100,
		Status: "ready",
	}

	mockRepo.On("CreateOrder", mock.MatchedBy(func(order canonical.Order) bool {
		return order.Products[0] == "product1" &&
			order.Products[1] == "product2" &&
			order.Products[2] == "product3" &&
			order.Total == float32(100) &&
			order.Status == "ready"
	})).Return(orderTest, nil)

	service := &service{
		repo: mockRepo,
	}

	order, err := service.CreateOrder(orderTest)

	assert.Nil(t, err)
	assert.Equal(t, "product1", order.Products[0])
	assert.Equal(t, "product2", order.Products[1])
	assert.Equal(t, "product3", order.Products[2])
	assert.Equal(t, float32(100), order.Total)
	assert.Equal(t, "ready", order.Status)

	mockRepo.AssertExpectations(t)
}

func TestCreateOrder_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	orderTest := canonical.Order{
		Products: []string{
			"product1", "product2", "product3",
		},
		Total:  100,
		Status: "ready",
	}

	mockRepo.On("CreateOrder", mock.MatchedBy(func(order canonical.Order) bool {
		return order.Products[0] == "product1" &&
			order.Products[1] == "product2" &&
			order.Products[2] == "product3" &&
			order.Total == float32(100) &&
			order.Status == "ready"
	})).Return(canonical.Order{}, errors.New("error occurred when trying to create an order"))

	service := &service{
		repo: mockRepo,
	}

	order, err := service.CreateOrder(orderTest)

	assert.NotNil(t, err)
	assert.Empty(t, order)
	assert.Equal(t, "error occurred when trying to create an order", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUpdateOrder_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	orderTest := canonical.Order{
		Products: []string{
			"product1", "product2", "product3",
		},
		Total:  100,
		Status: "ready",
	}

	mockRepo.On("UpdateOrder", "xpto", mock.MatchedBy(func(order canonical.Order) bool {
		return order.Products[0] == "product1" &&
			order.Products[1] == "product2" &&
			order.Products[2] == "product3" &&
			order.Total == float32(100) &&
			order.Status == "ready"
	})).Return(orderTest, nil)

	service := &service{
		repo: mockRepo,
	}

	order, err := service.UpdateOrder("xpto", orderTest)

	assert.Nil(t, err)
	assert.Equal(t, orderTest, order)

	mockRepo.AssertExpectations(t)
}

func TestUpdateOrder_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	orderTest := canonical.Order{
		Products: []string{
			"product1", "product2", "product3",
		},
		Total:  100,
		Status: "ready",
	}

	mockRepo.On("UpdateOrder", "xpto", orderTest).Return(canonical.Order{}, errors.New("error occurred when trying to update an order"))

	service := &service{
		repo: mockRepo,
	}

	order, err := service.UpdateOrder("xpto", orderTest)

	assert.NotNil(t, err)
	assert.Empty(t, order)
	assert.Equal(t, "error occurred when trying to update an order", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestDeleteOrder_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	orderTest := canonical.Order{
		Id: "xpto",
		Products: []string{
			"product1", "product2", "product3",
		},
		Total:     100,
		Status:    "ready",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetOrderById", "xpto").Return(orderTest, nil)

	mockRepo.On("DeleteOrder", "xpto").Return(nil)

	service := &service{
		repo: mockRepo,
	}

	err := service.DeleteOrder("xpto")

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteOrder_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	orderTest := canonical.Order{
		Id: "xpto",
		Products: []string{
			"product1", "product2", "product3",
		},
		Total:     100,
		Status:    "ready",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetOrderById", "xpto").Return(orderTest, nil)

	mockRepo.On("DeleteOrder", "xpto").Return(errors.New("error occurred when trying to delete an order"))

	service := &service{
		repo: mockRepo,
	}

	err := service.DeleteOrder("xpto")

	assert.NotNil(t, err)
	assert.Equal(t, "error occurred when trying to delete an order", err.Error())

	mockRepo.AssertExpectations(t)
}
