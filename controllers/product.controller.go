package controllers

import (
	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService interfaces.ProductServiceLayer
}

func NewProductController(ProductService interfaces.ProductServiceLayer) ProductController {
	return ProductController{
		ProductService: ProductService,
	}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Products
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := pc.ProductService.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *ProductController) GetProductByName(ctx *gin.Context) {
	var productName string = ctx.Param("name")
	product, err := pc.ProductService.GetProductByName(&productName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := pc.ProductService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) UpdateProductByName(ctx *gin.Context) {
	var productName string = ctx.Param("name")
	product, err := pc.ProductService.GetProductByName(&productName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProductByName(ctx *gin.Context) {
	var productName string = ctx.Param("name")
	err := pc.ProductService.DeleteProductByName(&productName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *ProductController) RegisterUserRoutes(rg *gin.RouterGroup) {
	productroute := rg.Group("/product")
	productroute.POST("/create", pc.CreateProduct)
	productroute.GET("/get/:name", pc.GetProductByName)
	productroute.GET("/getall", pc.GetAllProducts)
	productroute.PATCH("/update", pc.UpdateProductByName)
	productroute.DELETE("/delete/:name", pc.DeleteProductByName)
}
