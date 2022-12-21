package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"shome-backend/models"
	"shome-backend/pkg/read"
)

func ApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		var found = false

		var apiKeys []models.ApiKey
		var apiKey string

		apiKeys = read.ApiKeys()

		header := c.Request.Header
		apiKey = header.Get("x-api-key")

		for i := range apiKeys {
			if apiKeys[i].Key == apiKey {
				found = true
			}
		}

		if found == true {
			c.Next()
		} else {
			log.Println("Unauthenticated access by: " + c.ClientIP())
			c.Status(401)
			c.Abort()
		}
	}
}
