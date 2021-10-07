package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	param "shome-backend/flags"
	"shome-backend/middleware"
	"shome-backend/models"
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
		}
	} else {
		r.POST("/v2/blinder", blinderEndpoint)
	}

	r.Run(":80")
}

func blinderEndpoint(c *gin.Context) {
	log.Println("Endpoint Hit: Blinder")
	var _blinder models.Blinder

	if err := c.ShouldBind(&_blinder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
