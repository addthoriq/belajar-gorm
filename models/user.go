package models

import "time"

type Base struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Base
	Name     string
	Email    *string // Nullable
	Password string
}
