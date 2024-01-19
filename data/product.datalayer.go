package data

import (
	"context"
	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductDataLayerImpl struct {
	prodsCollection *mongo.Collection
	ctx             context.Context
}

func NewProductDataLayer(productsCollection *mongo.Collection, ctx context.Context) interfaces.ProductDataLayer {
	return &ProductDataLayerImpl{
		prodsCollection: productsCollection,
		ctx:             ctx,
	}
}

func (pd *ProductDataLayerImpl) CreateProduct(product *models.Products) error {
	_, err := pd.prodsCollection.InsertOne(pd.ctx, product)
	return err
}

func (pd *ProductDataLayerImpl) GetAllProducts() ([]*models.Products, error) {
	var allProducts []*models.Products

	cursor, err := pd.prodsCollection.Find(pd.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(pd.ctx) {
		var model models.Products
		err := cursor.Decode(&model)

		if err != nil {
			return nil, err
		}
		allProducts = append(allProducts, &model)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(pd.ctx)

	if len(allProducts) == 0 {
		return nil, errors.New("documents not found")
	}
	return allProducts, nil

}

func (pd *ProductDataLayerImpl) GetProductByName(name *string) (*models.Products, error) {

	var product *models.Products
	filter := bson.D{{Key: "name", Value: name}}
	err := pd.prodsCollection.FindOne(pd.ctx, filter).Decode(&product)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return product, nil
}

func (pd *ProductDataLayerImpl) DeleteProductByName(name *string) error {
	filter := bson.D{{Key: "name", Value: name}}
	_, err := pd.prodsCollection.DeleteOne(pd.ctx, filter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("product not found")
		}
		return err
	}
	return nil
}

func (pd *ProductDataLayerImpl) UpdateProductByName(name string, newProduct *models.Products) error {

	filter := bson.D{{Key: "name", Value: name}}
	update := bson.D{{Key: "$set", Value: newProduct}}
	result, err := pd.prodsCollection.UpdateOne(pd.ctx, filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}
