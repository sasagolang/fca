package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UID          int
	Name         string
	Email        string
	Password     string
	Mobile       string
	RegisterTime int64
	//Status       int
	NickName string
	HeadImg  string
	//ModelBase
}
type UserLoginRequest struct {
	UserName  string
	Password  string
	LoginType int
	Code      string
}

type RegisterUserRequest struct {
	Mobile     string
	VerifyCode string
	Pwd        string
}
