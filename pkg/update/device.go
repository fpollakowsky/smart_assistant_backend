package update

import (
	"shome-backend/models"
	"shome-backend/server/config"
)

func DeviceValue(channel, room string, value float64) error {
	err := config.DB.Model(&models.Device{}).Where("channel = ? and room = ?", channel, room).Update("value", value)
	if err != nil {
		return err.Error
	}

	return nil
}
