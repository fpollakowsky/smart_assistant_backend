package models

import (
	"gorm.io/gorm"
	"time"
)

type Routine struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Title       string         `json:"title"`
	Status      bool           `json:"status"`
	TriggerTime string         `json:"trigger_time"`
	Payload     []Payload      `json:"payload"`
}

type Payload struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	RoutineID int            `json:"-"`
	DeviceID  int            `json:"-"`
	Device    Device         `json:"device"`
	Value     int            `json:"value"`
}
