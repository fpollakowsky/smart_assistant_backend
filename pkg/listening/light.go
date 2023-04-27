package listening

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shome-backend/models"
	"shome-backend/pkg/read"
	"shome-backend/pkg/update"
	"shome-backend/server/mqtt"
)

func ChangeLightValue(c *gin.Context) {
	var light models.Light

	if err := c.ShouldBindJSON(&light); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, value := read.LightValue(light.Channel, light.Room)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if *value == "1" {
		mqtt.NewRequest(light.Channel, light.Room, "0")
		err := update.DeviceValue(light.Channel, light.Room, 0)
		if err != nil {
			c.JSON(http.StatusPreconditionFailed, gin.H{"error": err})
			return
		}
	} else {
		mqtt.NewRequest(light.Channel, light.Room, "1")
		err := update.DeviceValue(light.Channel, light.Room, 1)
		if err != nil {
			c.JSON(http.StatusPreconditionFailed, gin.H{"error": err})
			return
		}
	}
}
