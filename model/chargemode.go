package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ChargeFee struct {
	gorm.Model
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Price     int
}

type ChargeRule struct {
	gorm.Model
	Name       string
	ChargeFees []ChargeFee `gorm:"many2many:rule_fees;"`
}
