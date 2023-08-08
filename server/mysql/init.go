package mysql

import (
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"home-ai-backend/models"
	c "home-ai-backend/server/config"
	"home-ai-backend/server/logging"
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
		logging.PrintListFail("Initialize Database")
	}

	// Migrate the schema
	err = c.DB.AutoMigrate(&models.ApiKey{}, &models.Routine{}, &models.Device{}, &models.Payload{})
	if err != nil {
		logging.PrintListFail("Initialize Database")
	}

	logging.PrintListDone("Initialize Database")
}
