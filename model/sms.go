package model

import "github.com/jinzhu/gorm"
import "time"

type SMSInfo struct {
	gorm.Model
	UID      int
	Mobile   string
	Code     string
	SendTime time.Time
}
