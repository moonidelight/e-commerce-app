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
	Comment   string
	CreatedAt time.Time
	UserID    uint `gorm:"not null"`
	ItemID    uint `gorm:"not null"`
	Item      Item `gorm:"foreignKey:ItemID"`
}

type Rating struct {
	ID     uint `gorm:"primaryKey"`
	ItemID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Rating int64
	Item   Item `gorm:"foreignKey:ItemID"`
}

type Order struct {
	ID     uint        `gorm:"primaryKey"`
	UserID uint        `gorm:"not null"`
	Items  []OrderItem `gorm:"foreignKey:OrderID"`
	Total  float64
	Status bool
}
type OrderItem struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint `gorm:"not null"`
	ItemID  uint `gorm:"not null"`
	Item    Item `gorm:"foreignKey:ItemID"`
}

type Payment struct {
	ID          uint `gorm:"primaryKey"`
	Amount      float64
	PaymentDate time.Time
	OrderID     uint  `gorm:"not null"`
	Order       Order `gorm:"foreignKey:OrderID"`
}
