package models

import "time"

type User struct {
	Id        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Password  string
	createdAt time.Time
	updatedAt time.Time
}
