package controllers

import (
	"net/http"

	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/models"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService interfaces.OrderServiceLayer
}

func NewOrderController(orderService interfaces.OrderServiceLayer) OrderController {
	return OrderController{
		OrderService: orderService,
	}
}

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := oc.OrderService.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (oc *OrderController) GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")

	order, err := oc.OrderService.GetOrder(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetOrdersByUser(ctx *gin.Context) {
	userID := ctx.Param("userID")

	orders, err := oc.OrderService.GetOrdersByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (oc *OrderController) RegisterOrderRoutes(rg *gin.RouterGroup) {
	orderRoute := rg.Group("/order")
	orderRoute.POST("/create", oc.CreateOrder)
	orderRoute.GET("/get/:orderID", oc.GetOrder)
	orderRoute.GET("/getbyuser/:userID", oc.GetOrdersByUser)
}
