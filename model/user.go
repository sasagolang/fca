package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UID          int `gorm:"column:uid;"`
	Name         string
	Email        string
	Password     string
	Mobile       string
	RegisterTime int64
	//Status       int
	NickName string
	HeadImg  string
	CarSet   CarSet
	CarSetID uint
	//ModelBase
}
type UserLoginRequest struct {
	UserName  string
	Password  string
	LoginType int
	Code      string
}
type UserLoginResponse struct {
	Token string
	User  User
}
type RegisterUserRequest struct {
	Mobile     string
	VerifyCode string
	Pwd        string
}
type UpdateUserInfoRequest struct {
	UID      int 
	NickName string
	HeadImg  string
	Password string
	CarSetID int
}
