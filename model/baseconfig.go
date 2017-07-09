package model

import "github.com/jinzhu/gorm"

type BaseConfig struct {
	gorm.Model
	CompanyName      string
	CopyRight        string
	FeeDescription   string `gorm:"size:8000"`
	ServicePhone     string
	QQ               string
	SiteName         string
	RegisterProtocol string `gorm:"size:8000"`
}
