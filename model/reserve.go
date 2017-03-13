package model

import "github.com/jinzhu/gorm"
import "time"

type Reserve struct {
	gorm.Model
	User           User
	UserID         uint
	StartTime      time.Time
	EndTime        time.Time
	Status         uint
	ElectricPile   ElectricPile
	ElectricPileID uint
}
