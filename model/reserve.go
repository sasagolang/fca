package model

import "github.com/jinzhu/gorm"

type Reserve struct {
	gorm.Model
	User           User
	UserID         uint
	UID            int
	StartTime      int64
	EndTime        int64
	Status         uint
	ElectricPile   ElectricPile
	ElectricPileID uint
}

type CreateReserveRequest struct {
	ElectricPileID int
	StartTime      int64
	EndTime        int64
}

type MyReserveResponse struct {
	UserID    uint
	UID       int
	StartTime int64
	EndTime   int64
	Status    uint
	EPName    string
	EPAddress string
	EPNo      string
}
