package listening

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shome-backend/models"
	"shome-backend/pkg/add"
	"shome-backend/pkg/read"
	"shome-backend/pkg/remove"
)

func GetAllDevices(c *gin.Context) {
	err, devices := read.Device()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mysql failed"})
		return
	}

	c.JSON(http.StatusOK, devices)
}

func AddDevice(c *gin.Context) {
	var device models.Device

	if err := c.Bind(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := add.Device(device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully added device"})
}

func RemoveDevice(c *gin.Context) {
	var device models.Device

	if err := c.ShouldBind(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := remove.Device(device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully removed routine"})
}
