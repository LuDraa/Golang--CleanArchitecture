package services

import (
	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/models"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServiceLayer struct {
	ProductDataLayer interfaces.ProductDataLayer
}

func NewProductServiceLayer(ProductDataLayer interfaces.ProductDataLayer) interfaces.ProductServiceLayer {
	return &ProductServiceLayer{
		ProductDataLayer: ProductDataLayer,
	}
}

func (ps *ProductServiceLayer) CreateProduct(product *models.Products) error {
	err := ps.ProductDataLayer.CreateProduct(product)
	return err
}

func (ps *ProductServiceLayer) GetAllProducts() ([]*models.Products, error) {
	var allProducts []*models.Products

	allProducts, err := ps.ProductDataLayer.GetAllProducts()

	if err != nil {
		return nil, err
	}

	return allProducts, nil

}

func (ps *ProductServiceLayer) GetProductByName(name *string) (*models.Products, error) {

	var product *models.Products
	product, err := ps.ProductDataLayer.GetProductByName(name)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return product, nil
}

func (ps *ProductServiceLayer) DeleteProductByName(name *string) error {
	err := ps.ProductDataLayer.DeleteProductByName(name)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("product not found")
		}
		return err
	}
	return nil
}

func (ps *ProductServiceLayer) UpdateProductByName(name string, newProduct *models.Products) error {

	err := ps.ProductDataLayer.UpdateProductByName(name, newProduct)

	if err != nil {
		return err
	}

	return nil
}
