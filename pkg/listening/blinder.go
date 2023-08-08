package listening

import (
	"github.com/gin-gonic/gin"
	"home-ai-backend/models"
	"home-ai-backend/pkg/update"
	"home-ai-backend/server/mqtt"
	"net/http"
	"strconv"
)

func ChangeBlinderValue(c *gin.Context) {
	var blinder models.Blinder

	if err := c.ShouldBindJSON(&blinder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var status, err = strconv.ParseFloat(blinder.Value, 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	mqtt.NewRequest(blinder.Channel, blinder.Room, blinder.Value)

	err = update.DeviceValue(blinder.Channel, blinder.Room, status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mysql failed"})
		return
	}
}
