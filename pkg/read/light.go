package read

import (
	"errors"
	"shome-backend/models"
	"shome-backend/server/config"
)

func LightValue(channel, room string) (error, *string) {
	var device models.Device

	result := config.DB.Find(&device, "channel = ? AND room = ?", channel, room)
	if result.Error != nil {
		return result.Error, nil
	}

	if result.RowsAffected == 0 {
		return errors.New("zero devices found"), nil
	}

	return nil, device.Value
}
