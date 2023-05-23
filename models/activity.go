package models

import "time"

type Activity struct {
	Id        uint   `gorm:"column:id"`
	Title     string `gorm:"column:title" json:"title"`
	Email     string `gorm:"column:email" json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
