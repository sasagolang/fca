package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	User           User
	UserID         uint
	UID            int
	Status         uint
	ElectricPile   ElectricPile
	ElectricPileID uint
	Content        string
	Title          string
	CreateTime     int64
}
type CreateCommentRequest struct {
	ElectricPileID int
	Content        string
	Title          string
}
type GetCommentResponse struct {
	UID        int
	Name       string
	NickName   string
	HeadImg    string
	Content    string
	CreateTime int64
}
