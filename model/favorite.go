package model

import "github.com/jinzhu/gorm"

type Favorite struct {
	gorm.Model
	ElectricPile   ElectricPile
	ElectricPileID uint
	UID            int
}
