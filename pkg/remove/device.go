package remove

import (
	"errors"
	"home-ai-backend/models"
	"home-ai-backend/server/config"
)

func Device(device models.Device) error {
	result := config.DB.Delete(&device)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("REMOVE_ROUTINE:: no rows affected")
	}

	return nil
}
