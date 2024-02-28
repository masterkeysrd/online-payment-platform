package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
)

func MerchantId(service merchant.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := c.GetHeader("Authorization")

		if secretKey == "" {
			return
		}

		merchant, err := service.GetIDByApiKey(secretKey)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("merchant", merchant)
	}
}
