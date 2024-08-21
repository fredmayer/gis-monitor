package models

import (
	"gorm.io/gorm"
)

type Lot struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	Status         string
	LotName        string
	LotDescription string `gorm:"type:longtext"`
	PriceMin       float64
}
