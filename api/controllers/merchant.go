package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
)

type MerchantController struct {
	service merchant.Service
}

type MerchantControllerParams struct {
	Service merchant.Service
}

func NewMerchantController(params MerchantControllerParams) *MerchantController {
	return &MerchantController{
		service: params.Service,
	}
}

func (c *MerchantController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", c.Create)
}

func (c *MerchantController) Create(ctx *gin.Context) {
	var request merchant.CreateMerchantRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := c.service.Create(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
