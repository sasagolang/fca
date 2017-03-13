package model

import "github.com/jinzhu/gorm"

type BaseConfig struct {
	gorm.Model
	CompanyName      string
	CopyRight        string
	FeeDescription   string
	ServicePhone     string
	QQ               string
	SiteName         string
	RegisterProtocol string `gorm:"size:4000"`
}
