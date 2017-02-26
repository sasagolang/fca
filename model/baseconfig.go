package model

import "github.com/jinzhu/gorm"

type BaseConfig struct {
	gorm.Model
	SiteName         string
	RegisterProtocol string `gorm:"size:4000"`
}
