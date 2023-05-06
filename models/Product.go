package models

import "gorm.io/gorm"

type Item struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64
	Rating      int64
	IsActive    bool
}
type Order struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint
	orderStatus bool
}

type Comment struct {
	gorm.Model
	Rating []Rating `gorm:"foreignKey:CommentRefer"`
}

type Rating struct {
	gorm.Model
	Number       string
	CommentRefer uint
}
