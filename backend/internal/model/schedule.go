package model

import "time"

type Schedule struct {
	ID int `gorm:"primaryKey"`
	Task Task `gorm:"not null, foreignKey:TaskID"`
	TaskID int `gorm:"not null"`
	Developer Developer `gorm:"not null, foreignKey:DeveloperID"`
	DeveloperID int `gorm:"not null"`
	SprintWeek int `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}