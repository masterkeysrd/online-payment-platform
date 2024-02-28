package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
}

func NewMerchantController() *MerchantController {
	return &MerchantController{}
}

func (c *MerchantController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", c.Create)
}

func (c *MerchantController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Merchant created",
	})
}
