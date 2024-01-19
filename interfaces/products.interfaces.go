package interfaces

import "ecommerce/gmr/models"

type ProductDataLayer interface {
	CreateProduct(*models.Products) error
	GetAllProducts() ([]*models.Products, error)
	GetProductByName(*string) (*models.Products, error)
	DeleteProductByName(*string) error
	UpdateProductByName(string, *models.Products) error
}

type ProductServiceLayer interface {
	CreateProduct(*models.Products) error
	GetAllProducts() ([]*models.Products, error)
	GetProductByName(*string) (*models.Products, error)
	DeleteProductByName(*string) error
	UpdateProductByName(string, *models.Products) error
}
