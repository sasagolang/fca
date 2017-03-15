package model

import "github.com/jinzhu/gorm"

type Pole struct {
	gorm.Model
	ElectricPile ElectricPile
	ElectricPileID uint
	UUID string
	Imei string
	Imsi string
	Status int
}
