package model

import "time"

type Task struct {
	ID        int    `gorm:"primaryKey"`
	Value     int `gorm:"not null"`
	EstimatedDuration time.Duration `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
