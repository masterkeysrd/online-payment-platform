package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/customer"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/paymentmethod"
)

type CustomerController struct {
	customerService      customer.Service
	paymentMethodService paymentmethod.Service
}

type CustomerControllerParams struct {
	CustomerService      customer.Service
	PaymentMethodService paymentmethod.Service
}

func NewCustomerController(params CustomerControllerParams) *CustomerController {
	return &CustomerController{
		customerService:      params.CustomerService,
		paymentMethodService: params.PaymentMethodService,
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

	response, err := c.customerService.CreateCustomer(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *CustomerController) GetPaymentMethods(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"error": "Not implemented",
	})
}

func (c *CustomerController) AddPaymentMethod(ctx *gin.Context) {
	merchant := ctx.GetString("merchant")
	customerId := ctx.Param("customerId")

	var request paymentmethod.CreatePaymentMethodRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	request.Merchant = merchant
	request.Customer = customerId

	response, err := c.paymentMethodService.Create(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *CustomerController) GetPaymentMethod(ctx *gin.Context) {
	merchant := ctx.GetString("merchant")
	customerId := ctx.Param("customerId")
	paymentMethodId := ctx.Param("paymentMethodId")

	request := paymentmethod.GetPaymentMethodRequest{
		Merchant:      merchant,
		Customer:      customerId,
		PaymentMethod: paymentMethodId,
	}

	response, err := c.paymentMethodService.Get(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
