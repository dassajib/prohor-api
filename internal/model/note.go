package model

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	// User User `gorm:"foreignKey:UserID"`
	Title     string `gorm:"not null;size:255"`
	Content   string `gorm:"type:text"`
	Tag       string `gorm:"not null;size:30"`
	Date      time.Time
	Pinned    bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// for soft delete
	// `gorm:"index"` | create index col for fast query
	// when soft delete, this field will mark instead of delete
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
