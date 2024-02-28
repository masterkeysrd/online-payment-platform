package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/customer"
)

type CustomerController struct {
	service customer.Service
}

type CustomerControllerParams struct {
	CustomerService customer.Service
}

func NewCustomerController(params CustomerControllerParams) *CustomerController {
	return &CustomerController{
		service: params.CustomerService,
	}
}

func (c *CustomerController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", c.Create)
	router.GET("/:customerId/payments-methods", c.GetPaymentMethods)
	router.POST("/:customerId/payments-methods", c.AddPaymentMethod)
	router.GET("/:customerId/payments-methods/:paymentMethodId", c.GetPaymentMethod)
}

func (c *CustomerController) Create(ctx *gin.Context) {
	var merchant = ctx.GetString("merchant")

	var request customer.CreateCustomerRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	request.Merchant = merchant

	response, err := c.service.CreateCustomer(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
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
