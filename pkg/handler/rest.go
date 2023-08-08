package handler

import (
	"github.com/charmbracelet/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"home-ai-backend/pkg/listening"
	"home-ai-backend/pkg/middleware"
	"home-ai-backend/server/config"
	"io"
	"os"
	"time"
)

func HandleRequests() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.RFC822,
	})

	if !*config.IS_DEBUG {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	} else {
		gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
			logger.Infof("%-15v %-10v %v", absolutePath, httpMethod, handlerName)
		}
	}

	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.MaxMultipartMemory = 8 << 20

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PATCH", "POST", "OPTIONS", "GET", "PUT"},
		AllowHeaders:     []string{"x-api-key,content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	authorized := r.Group("/")

	authorized.Use(middleware.ApiKey())
	{
		authorized.PATCH("/v2/blinder", listening.ChangeBlinderValue) // changes blinder close / open value
		authorized.PATCH("/v2/light", listening.ChangeLightValue)     // changes on / off value of lights

		authorized.GET("/v2/devices", listening.GetAllDevices)   // gets all devices from db
		authorized.PUT("/v2/devices", listening.AddDevice)       // adds device to db
		authorized.DELETE("/v2/devices", listening.RemoveDevice) // removes device from db

		authorized.GET("/v2/routine", listening.GetAllRoutines)   // gets all routines from db
		authorized.PUT("/v2/routine", listening.AddRoutine)       // adds a routine to db
		authorized.DELETE("/v2/routine", listening.RemoveRoutine) // removes a routine to db
		authorized.PATCH("/v2/routine", listening.UpdateRoutine)  // removes a routine to db
	}

	err := r.Run(":80")
	if err != nil {
		// todo add exception
		return
	}
}
