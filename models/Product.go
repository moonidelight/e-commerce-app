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
type CommentDetail struct {
	ID        uint `gorm:"primaryKey"`
	Comment   string
	CreatedAt time.Time
}
type Comment struct {
	ID       uint            `gorm:"primaryKey"`
	ItemID   uint            `gorm:"not null"`
	UserID   uint            `gorm:"not null"`
	Comments []CommentDetail `gorm:"many2many:ID"`
}

type RatingDetail struct {
	ID     uint `gorm:"primaryKey"`
	Rating int64
}
type Rating struct {
	RatingID     uint         `gorm:"not null"`
	UserID       uint         `gorm:"not null"`
	ItemID       uint         `gorm:"not null"`
	RatingDetail RatingDetail `gorm:"foreignKey:RatingID"`
}
type Order struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Status    bool
	UserID    uint `gorm:"not null"`
	Total     float64
}
type OrderItem struct {
	OrderID uint `gorm:"not null"`
	ItemID  uint `gorm:"not null"`
	Item    Item `gorm:"foreignKey:ItemID"`
}

type PaymentDetail struct {
	ID          uint `gorm:"primaryKey"`
	Amount      float64
	PaymentDate time.Time
}
type Payment struct {
	PaymentID uint `gorm:"not null"`
	OrderID   uint `gorm:"not null"`
	UserID    uint `gorm:"not null"`
}
