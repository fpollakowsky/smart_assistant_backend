package update

import (
	"home-ai-backend/models"
	"home-ai-backend/server/config"
)

func Routine(routine models.Routine) error {
	result := config.DB.Model(&models.Routine{}).Where("id = ?", routine.ID).Omit("id").Updates(routine)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
