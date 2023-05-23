package models

import "time"

type Todo struct {
	Id              uint   `gorm:"column:id"`
	Title           string `gorm:"column:title" json:"title"`
	ActivityGroupId uint   `gorm:"column:activity_group_id" json:"activity_group_id"`
	IsActive        bool   `gorm:"column:is_active" json:"is_active"`
	Priority        string `gorm:"column:priority" json:"priority"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
