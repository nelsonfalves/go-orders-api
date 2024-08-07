package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/nelsonalves117/go-orders-api/internal/canonical"
	"github.com/nelsonalves117/go-orders-api/internal/repositories"
	"github.com/sirupsen/logrus"
	"time"
)

type Service interface {
	GetAllOrders() ([]canonical.Order, error)
	GetOrderById(id string) (canonical.Order, error)
	CreateOrder(order canonical.Order) (canonical.Order, error)
	UpdateOrder(id string, order canonical.Order) (canonical.Order, error)
	DeleteOrder(id string) error
}

type service struct {
	repo repositories.Repository
}

func New() Service {
	return &service{
		repo: repositories.New(),
	}
}

func (service *service) GetAllOrders() ([]canonical.Order, error) {
	order, err := service.repo.GetAllOrders()
	if err != nil {
		logrus.WithError(err).Error("error occurred when trying to get all orders") // Error at beginning of file
		return []canonical.Order{}, err
	}

	return order, nil
}

func (service *service) GetOrderById(id string) (canonical.Order, error) {
	order, err := service.repo.GetOrderById(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred when trying to get an order")
		return canonical.Order{}, err
	}

	return order, nil
}

func (service *service) CreateOrder(order canonical.Order) (canonical.Order, error) {
	order.Id = uuid.NewString()
	order.CreatedAt = time.Now()

	order, err := service.repo.CreateOrder(order)
	if err != nil {
		logrus.WithError(err).Error("error occurred when trying to create an order")
		return canonical.Order{}, err
	}

	return order, nil
}

func (service *service) UpdateOrder(id string, order canonical.Order) (canonical.Order, error) {
	order, err := service.repo.UpdateOrder(id, order)
	if err != nil {
		logrus.WithError(err).Error("error occurred when trying to update an order")
		return canonical.Order{}, err
	}

	return order, nil
}

func (service *service) DeleteOrder(id string) error {
	order, err := service.repo.GetOrderById(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred when trying to get an order")
		return err
	}

	if order.Id == "" {
		return fmt.Errorf("order not found on db")
	}

	err = service.repo.DeleteOrder(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred when trying to delete an order")
		return err
	}

	return nil
}
