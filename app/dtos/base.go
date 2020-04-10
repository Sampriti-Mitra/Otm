package dtos

import "time"

type Base struct {
	Success bool
	Error   error
}

type Model struct {
	Id        int       `gorm:"integer;not null" json:"id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
