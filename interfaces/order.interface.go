package interfaces

import (
	"ecommerce/gmr/models"
)

type OrderDataLayer interface {
	CreateOrder(order *models.Order) error
	GetOrder(orderID string) (*models.Order, error)
	GetOrdersByUserID(userID string) ([]*models.Order, error)
}

type OrderServiceLayer interface {
	CreateOrder(order *models.Order) error
	GetOrder(orderID string) (*models.Order, error)
	GetOrdersByUserID(userID string) ([]*models.Order, error)
}
