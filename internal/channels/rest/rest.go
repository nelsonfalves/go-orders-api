package rest

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nelsonalves117/go-orders-api/internal/config"
	"github.com/nelsonalves117/go-orders-api/internal/service"
	"net/http"
)

type Rest interface {
	Start() error
}

type rest struct {
	service service.Service
}

func New() Rest {
	return &rest{
		service: service.New(),
	}
}

func (rest *rest) Start() error {
	router := echo.New()

	router.Use(middleware.Logger())

	router.GET("/orders", rest.GetAllOrders)
	router.GET("/orders/:id", rest.GetOrderById)
	router.POST("/orders/create", rest.CreateOrder)
	router.PUT("/orders/update/:id", rest.UpdateOrder)
	router.DELETE("/orders/delete/:id", rest.DeleteOrder)

	return router.Start(":" + config.Get().Port)
}

func (rest *rest) GetAllOrders(c echo.Context) error {
	orderSlice, err := rest.service.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, orderSlice)
}

func (rest *rest) GetOrderById(c echo.Context) error {
	id := c.Param("id")

	order, err := rest.service.GetOrderById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, order)
}

func (rest *rest) CreateOrder(c echo.Context) error {
	var order orderRequest

	err := c.Bind(&order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	createdOrder, err := rest.service.CreateOrder(toCanonical(order))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusCreated, toResponse(createdOrder))
}

func (rest *rest) UpdateOrder(c echo.Context) error {
	var order orderRequest

	err := c.Bind(&order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	id := c.Param("id")
	updatedOrder, err := rest.service.UpdateOrder(id, toCanonical(order))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, toResponse(updatedOrder))
}

func (rest *rest) DeleteOrder(c echo.Context) error {
	id := c.Param("id")

	err := rest.service.DeleteOrder(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, nil)
}
