package models

import (
	"time"
)

type Item struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64
	Rating      float64
	IsActive    bool
	Comments    []Comment `gorm:"many2many:item_comments"`
}
type UserOrders struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Orders []Item `gorm:"many2many:user_orders_items"`
	Status bool
}

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	ItemID    uint
	UserID    uint
	Comment   string
	CreatedAt time.Time
}

type Rating struct {
	ID     uint `gorm:"primaryKey"`
	ItemID uint
	UserID uint
	Rating int64
}
