package services

import (
	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/models"
)

type OrderServiceImpl struct {
	OrderDataLayer interfaces.OrderDataLayer
}

func NewOrderServiceImpl(odl interfaces.OrderDataLayer) interfaces.OrderServiceLayer {
	return &OrderServiceImpl{
		OrderDataLayer: odl,
	}
}

func (osl *OrderServiceImpl) CreateOrder(order *models.Order) error {
	err := osl.OrderDataLayer.CreateOrder(order)
	return err
}

func (osl *OrderServiceImpl) GetOrder(orderID string) (*models.Order, error) {
	var order *models.Order
	order, err := osl.OrderDataLayer.GetOrder(orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (osl *OrderServiceImpl) GetOrdersByUserID(userID string) ([]*models.Order, error) {
	var orders []*models.Order
	orders, err := osl.OrderDataLayer.GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
