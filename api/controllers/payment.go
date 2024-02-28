package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/payment"
)

type PaymentController struct {
	paymentService payment.Service
}

type PaymentControllerParams struct {
	PaymentService payment.Service
}

func NewPaymentController(params PaymentControllerParams) *PaymentController {
	return &PaymentController{
		paymentService: params.PaymentService,
	}
}

func (c *PaymentController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", c.Create)
	router.GET("/:paymentId", c.GetPayment)
}

func (c *PaymentController) Create(ctx *gin.Context) {
	merchant := ctx.GetString("merchant")
	var request payment.CreatePaymentRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.Merchant = merchant
	response, err := c.paymentService.Create(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *PaymentController) GetPayment(ctx *gin.Context) {
	merchant := ctx.GetString("merchant")
	paymentId := ctx.Param("paymentId")

	response, err := c.paymentService.Get(payment.GetPaymentRequest{
		Merchant:  merchant,
		PaymentID: paymentId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
