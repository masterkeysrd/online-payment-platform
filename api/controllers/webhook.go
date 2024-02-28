package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type WebhookController struct {
}

type WebhookControllerParams struct {
}

func NewWebhookController(params WebhookControllerParams) *WebhookController {
	return &WebhookController{}
}

func (c *WebhookController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", c.Webhook)
}

func (c *WebhookController) Webhook(ctx *gin.Context) {
	fmt.Println("Webhook received", ctx.Request.Body)
	ctx.JSON(200, gin.H{
		"message": "Webhook received",
	})
}
