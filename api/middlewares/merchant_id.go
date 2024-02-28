package middlewares

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
)

func MerchantId(service merchant.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		fields := strings.Fields(authorization)

		var secretKey string
		if len(fields) == 2 {
			secretKey = fields[1]
		} else {
			secretKey = fields[0]
		}

		log.Println("secretKey", secretKey)

		if secretKey == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
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
