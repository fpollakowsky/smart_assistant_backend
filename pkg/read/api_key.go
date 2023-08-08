package read

import (
	"home-ai-backend/models"
	"home-ai-backend/server/config"
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
