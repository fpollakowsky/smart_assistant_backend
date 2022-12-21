package models

type ApiKey struct {
	Key     string `json:"key,omitempty" gorm:"primaryKey"`
	Company string `json:"company,omitempty" gorm:"index"`
}
