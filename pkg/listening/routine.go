package listening

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shome-backend/models"
	"shome-backend/pkg/add"
	"shome-backend/pkg/read"
	"shome-backend/pkg/remove"
	"shome-backend/pkg/update"
)

func GetAllRoutines(c *gin.Context) {
	err, devices := read.Routine()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, devices)
}

func AddRoutine(c *gin.Context) {
	var routine models.Routine

	if err := c.Bind(&routine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := add.Routine(routine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully added routine"})
}

func RemoveRoutine(c *gin.Context) {
	var routine models.Routine

	if err := c.ShouldBind(&routine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := remove.Routine(routine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully removed routine"})
}

func UpdateRoutine(c *gin.Context) {
	var routine models.Routine

	if err := c.ShouldBind(&routine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := update.Routine(routine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated routine"})
}
