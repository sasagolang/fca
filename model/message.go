package model

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	UID        int
	Title      string
	Content    string
	Status     int
	CreateTime int64
}
