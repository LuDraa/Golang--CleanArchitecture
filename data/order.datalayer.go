package data

import (
	"context"
	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderDataLayerImpl struct {
	orderCollection *mongo.Collection
	ctx             context.Context
}

func NewOrderDataLayerImpl(orderCollection *mongo.Collection, ctx context.Context) interfaces.OrderDataLayer {
	return &OrderDataLayerImpl{
		orderCollection: orderCollection,
		ctx:             ctx,
	}
}

func (odl *OrderDataLayerImpl) CreateOrder(order *models.Order) error {
	_, err := odl.orderCollection.InsertOne(odl.ctx, order)
	return err
}

func (odl *OrderDataLayerImpl) GetOrder(orderID string) (*models.Order, error) {
	var order models.Order
	err := odl.orderCollection.FindOne(odl.ctx, bson.M{"_id": orderID}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (odl *OrderDataLayerImpl) GetOrdersByUserID(userID string) ([]*models.Order, error) {
	var orders []*models.Order
	cursor, err := odl.orderCollection.Find(odl.ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(odl.ctx)

	for cursor.Next(odl.ctx) {
		var order *models.Order
		err := cursor.Decode(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
