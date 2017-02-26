package model

import "github.com/jinzhu/gorm"

type MyWallet struct {
	gorm.Model
	UID           int
	Balance       int64
	FreezeBalance int64
}

type MyWalletLog struct {
	gorm.Model
	UID  int
	Type string
	Info string
}
