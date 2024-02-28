package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)	

type CustomerController struct {
}

func NewCustomerController() *CustomerController {
	return &CustomerController{}
}

func (c *CustomerController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", c.Create)
	router.GET("/:customerId/payments-methods", c.GetPaymentMethods)
	router.POST("/:customerId/payments-methods", c.AddPaymentMethod)
	router.GET("/:customerId/payments-methods/:paymentMethodId", c.GetPaymentMethod)
}

func (c *CustomerController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Customer created",
	})
}

func (c *CustomerController) GetPaymentMethods(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get payment methods",
	})
}

func (c *CustomerController) AddPaymentMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Payment method added",
	})
}

func (c *CustomerController) GetPaymentMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get payment method",
	})
}