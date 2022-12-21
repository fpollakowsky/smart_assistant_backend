package add

import (
	"errors"
	"fmt"
	"shome-backend/models"
	"shome-backend/server/config"
)

func Device(device models.Device) error {
	result := config.DB.Create(&device)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Println(result.RowsAffected)
		return errors.New("ADD_DEVICE:: no rows affected")
	}

	return nil
}
