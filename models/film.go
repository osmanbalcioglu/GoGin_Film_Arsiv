package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	Title    string  `gorm:"not null" json:"title" binding:"required"`
	Genre    string  `gorm:"not null" json:"genre" binding:"required"`
	Director string  `gorm:"not null" json:"director" binding:"required"`
	Year     int     `gorm:"not null" json:"year" binding:"required,gt=1887"`
	Rating   float64 `gorm:"not null" json:"rating" binding:"required,min=0,max=10"`
	Poster   string  `json:"poster"`
}
