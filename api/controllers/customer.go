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
}

func (c *CustomerController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Customer created",
	})
}