package model

import "github.com/jinzhu/gorm"

type Pole struct {
	gorm.Model
	ElectricPile   ElectricPile
	ElectricPileID uint
	UUID           string
	Imei           string
	Imsi           string
	NO             int
	Status         int
	Amount         float32
	Duration       int
	State          string
}
