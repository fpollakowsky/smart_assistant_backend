package add

import (
	"errors"
	"fmt"
	"shome-backend/models"
	"shome-backend/server/config"
)

func Routine(routine models.Routine) error {
	result := config.DB.Create(&routine)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Println(result.RowsAffected)
		return errors.New("ADD_ROUTINE:: no rows affected")
	}

	return nil
}
