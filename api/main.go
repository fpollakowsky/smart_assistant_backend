package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	param "shome-backend/flags"
	"shome-backend/middleware"
	"shome-backend/models"
	"shome-backend/mqtt"
	"shome-backend/mysql"
)

func HandleRequests() {
	//gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if !param.Debug {
		nethcon := middleware.APIKeyMiddleware{APIToken: make(map[string]string)}
		nethcon.Populate()

		authorized := r.Group("/")
		authorized.Use(nethcon.TokenAuthMiddleware())
		{
			authorized.POST("/v2/blinder", blinderEndpoint)
			authorized.POST("/v2/light", lightEndpoint)
			authorized.POST("/v2/devices", deviceEnpoint)
		}
	} else {
		r.POST("/v2/blinder", blinderEndpoint)
		r.POST("/v2/light", lightEndpoint)
		r.POST("/v2/devices", deviceEnpoint)
	}

	r.Run(":80")
}

func deviceEnpoint(c *gin.Context) {
	log.Println("Endpoint Hit: Device")

	var _devices, err = mysql.GetDevices()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"devices": _devices})
}

func blinderEndpoint(c *gin.Context) {
	log.Println("Endpoint Hit: Blinder")
	var _blinder models.Blinder

	if err := c.ShouldBind(&_blinder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func lightEndpoint(c *gin.Context) {
	log.Println("Endpoint Hit: Light")
	var _light models.Light

	if err := c.ShouldBind(&_light); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client = mqtt.Connect()

	status, err := mysql.GetLightStatus(_light.Channel, _light.Room)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if status == 1 {
		mqtt.NewRequest(client, _light.Channel, _light.Room, "0")
		err := mysql.UpdateLightStatus(_light.Channel, _light.Room, 0)
		if err != nil {
			c.JSON(http.StatusPreconditionFailed, gin.H{"error": err})
			return
		}
	} else {
		mqtt.NewRequest(client, _light.Channel, _light.Room, "1")
		err := mysql.UpdateLightStatus(_light.Channel, _light.Room, 1)
		if err != nil {
			c.JSON(http.StatusPreconditionFailed, gin.H{"error": err})
			return
		}
	}
}
