package read

import (
	"shome-backend/models"
	"shome-backend/server/config"
)

func ApiKeys() []models.ApiKey {
	var apiKey []models.ApiKey

	result := config.DB.Find(&apiKey)

	if result.Error != nil {
		return nil
	}

	if result.RowsAffected == 0 {
		return nil
	}

	return apiKey
}
