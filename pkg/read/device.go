package read

import (
	"errors"
	"shome-backend/models"
	"shome-backend/server/config"
)

func Device() (error, []models.Device) {
	var device []models.Device

	result := config.DB.Find(&device)
	if result.Error != nil {
		return result.Error, nil
	}

	if result.RowsAffected == 0 {
		return errors.New("zero devices found"), nil
	}

	return nil, device
}
