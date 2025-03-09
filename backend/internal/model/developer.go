package model

import "time"

type Developer struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Seniority int `gorm:"not null"`
	WeeklyWorkHours int `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
