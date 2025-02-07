package models

import (
	"time"

	"gorm.io/gorm"
)

// Base struct for all models (ActiveRecord-like)
type Base struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Save inserts or updates the record
func (b *Base) Save(db *gorm.DB) error {
	return db.Save(b).Error
}

// Delete soft deletes the record
func (b *Base) Delete(db *gorm.DB) error {
	return db.Delete(b).Error
}
