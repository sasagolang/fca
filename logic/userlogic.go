package logic

import (
	"encoding/base64"
	"errors"
	"fca/dal"
	"fca/libs"
	"fca/model"
	"fmt"
	"time"
)

var Aeskey = "qwerasdfzxcvtgbyhn567890123456bh"

func (logic LogicBase) UserLogin(username string, pwd string, code string, loginType int) (user *model.User, err error) {
	u := new(model.User)

	if dal.DB.Where("mobile=?", username).First(&u).RecordNotFound() {
		return nil, errors.New("用户不存在")
	}

	return u, nil
}
func (logic LogicBase) RegisterUser(mobile string, email string, cname string, pwd string) (u *model.User, err error) {

	user := new(model.User)

	if !dal.DB.Where("mobile=?", mobile).First(&user).RecordNotFound() {
		return nil, errors.New("用户已注册")
	}

	user.UID = time.Now().Nanosecond()
	user.Name = cname
	user.Email = email
	user.Mobile = mobile
	user.NickName = cname
	if cname == "" {
		user.Name = "anonymous"
		user.NickName = "anonymous"
	}
	data, _ := libs.AesEncrypt([]byte(pwd), []byte(Aeskey))
	pwdstr := base64.StdEncoding.EncodeToString(data)
	user.Password = pwdstr
	user.RegisterTime = time.Now().UnixNano()
	//user.Status = 1
	fmt.Printf("user:%v\n", user)
	//err = Dao.AddUser(user)
	dal.DB.NewRecord(user)
	dal.DB.Create(user)
	user.Password = ""
	return user, nil

}
func (logic LogicBase) GetUserByUid(uid int) (user *model.User, err error) {
	var usertmp model.User
	fmt.Printf("GetUserByUid:%v\n", uid)
	if dal.DB.Where("uid=?", uid).First(&usertmp).RecordNotFound() {
		fmt.Printf("GetUserByUid:%v\n", usertmp)
		return nil, nil
	}
	usertmp.Password = ""
	fmt.Printf("GetUserByUid:%v\n", usertmp)
	return &usertmp, nil
}
