package mysql

import (
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"shome-backend/models"
	c "shome-backend/server/config"
)

func InitializeDatabase() {
	// todo remove root access
	var err error

	c.DB, err = gorm.Open(gmysql.Open(c.DSN), &gorm.Config{
		AllowGlobalUpdate:                        true,
		DisableForeignKeyConstraintWhenMigrating: false,
		//QueryFields:                              true,
	})
	if err != nil {
		log.Println("[DATABASE] Failed to open connection")
	} else {
		log.Println("[DATABASE] Connection established")
	}

	// Migrate the schema
	err = c.DB.AutoMigrate(&models.ApiKey{}, &models.Routine{}, &models.Device{}, &models.Payload{})
}
