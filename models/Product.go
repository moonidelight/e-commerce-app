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
}

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	ItemID    uint `gorm:"not null"`
	Comment   string
	CreatedAt time.Time
}

type Rating struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	ItemID uint `gorm:"not null"`
	Rating int64
}

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	CreatedAt time.Time
	Status    bool
	Items     []OrderItem `gorm:"foreignKey:OrderID"`
}
type OrderItem struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint `gorm:"not null"`
	ItemID  uint `gorm:"not null"`
}
type Payment struct {
	ID          uint `gorm:"primaryKey"`
	OrderID     uint `gorm:"not null"`
	Amount      float64
	PaymentDate time.Time
}
