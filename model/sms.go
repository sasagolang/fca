package model

import "github.com/jinzhu/gorm"
import "time"

type SMSInfo struct {
	gorm.Model
	UID      int `gorm:"column:uid;"`
	Mobile   string
	Code     string
	SendTime time.Time
}
