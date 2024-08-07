package rest

import (
	"github.com/nelsonalves117/go-orders-api/internal/canonical"
)

func toCanonical(order orderRequest) canonical.Order {
	return canonical.Order{
		Products: order.Products,
		Total:    order.Total,
		Status:   order.Status,
	}
}

func toResponse(order canonical.Order) orderResponse {
	return orderResponse{
		Id:        order.Id,
		Products:  order.Products,
		Total:     order.Total,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
	}
}
