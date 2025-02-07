package models

import (
	"errors"

	"gorm.io/gorm"
)

// User model
type User struct {
	Base
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Age   int    `gorm:"default:18"`
}

// Hooks (ActiveRecord-style)
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

// Find a user by ID
func FindUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	return &user, err
}

// Find users by condition
func FindUsersByEmail(db *gorm.DB, email string) ([]User, error) {
	var users []User
	err := db.Where("email = ?", email).Find(&users).Error
	return users, err
}

// Create a new user
func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

// Update user
func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}
