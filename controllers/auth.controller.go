package controllers

import (
	"ecommerce/gmr/interfaces"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService interfaces.AuthServiceLayer
}

func NewAuthController(as interfaces.AuthServiceLayer) AuthController {
	return AuthController{
		AuthService: as,
	}
}

func (ac *AuthController) RegisterAuthRoutes(group *gin.RouterGroup) {
	group.POST("/login", ac.Login)
}

func (ac *AuthController) Login(c *gin.Context) {

	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	token, err := ac.AuthService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
