package model

import "github.com/jinzhu/gorm"

type ChargeOrder struct {
	gorm.Model
	NO         int64
	UUID       string
	PoleNO     int
	UID        int `gorm:"column:uid;"`
	Amount     float32
	Duration   int
	State      string
	Status     int
	StatusName string
	Energy     int32
	EPName     string
	EPAddress  string
	StartTime  int64
	EndTime    int64
}
