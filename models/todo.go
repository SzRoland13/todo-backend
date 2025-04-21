package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Priority    Priority  `gorm:"default:1" json:"priority"`
	Progress    Progress  `gorm:"default:1" json:"progress"`
	DueDate     time.Time `json:"due_date"`

	UserID uint `json:"user_id"`
}

type Priority int

const (
	LOW Priority = iota + 1
	MEDIUM
	HIGH
)

func (p Priority) String() string {
	return [...]string{"LOW", "MEDIUM", "HIGH"}[p-1]
}

func (p Priority) EnumIndex() int {
	return int(p)
}

type Progress int

const (
	TO_DO Progress = iota + 1
	IN_PROGRESS
	ON_HOLD
	COMPLETED
)

func (p Progress) String() string {
	return [...]string{"TO_DO", "IN_PROGRESS", "ON_HOLD", "COMPLETED"}[p-1]
}

func (p Progress) EnumIndex() int {
	return int(p)
}
