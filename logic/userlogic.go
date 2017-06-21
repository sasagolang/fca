package logic

import (
	"encoding/base64"
	"errors"
	"fca/dal"
	"fca/libs"
	"fca/model"
	"fmt"
	"strings"
	"time"
)

var Aeskey = "qwerasdfzxcvtgbyhn567890123456bh"

func (logic LogicBase) UserLogin(username string, pwd string, code string, loginType int) (user *model.User, err error) {
	u := new(model.User)

	err = logic.CheckCodeByMobile(username, code)
	if err != nil {
		return nil, err
	}
	if dal.DB.Preload("CarSet").Preload("CarSet.CarBrand").Where("mobile=?", username).First(&u).RecordNotFound() {
		return nil, errors.New("用户不存在")
	}
	data, _ := libs.AesEncrypt([]byte(pwd), []byte(Aeskey))
	pwdstr := base64.StdEncoding.EncodeToString(data)
	if !strings.EqualFold(u.Password, pwdstr) {
		fmt.Printf("%s,%s", u.Password)
		return nil, errors.New("用户名或者密码错误")
	}

	return u, nil
}
func (logic LogicBase) RegisterUser(mobile string, email string, cname string, pwd string) (u *model.User, err error) {

	user := new(model.User)

	if !dal.DB.Where("mobile=?", mobile).First(&user).RecordNotFound() {
		//用户已注册，直接修改密码
		data, _ := libs.AesEncrypt([]byte(pwd), []byte(Aeskey))
		pwdstr := base64.StdEncoding.EncodeToString(data)

		dal.DB.Model(&user).Updates(map[string]interface{}{"Password": pwdstr})
		return user, nil
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
	if dal.DB.Preload("CarSet").Preload("CarSet.CarBrand").Where("uid=?", uid).First(&usertmp).RecordNotFound() {
		fmt.Printf("GetUserByUid:%v\n", usertmp)
		return nil, nil
	}
	usertmp.Password = ""
	fmt.Printf("GetUserByUid:%v\n", usertmp)
	return &usertmp, nil
}
func (logic LogicBase) UpdateUserInfo(uid int, nickName string, headImg string, password string, carsetid int) error {
	u := make(map[string]interface{}, 0)
	var user model.User
	if dal.DB.Where("uid=?", uid).First(&user).RecordNotFound() {
		return errors.New("用户不存在")
	} else {
		if nickName != "" {
			u["nickName"] = nickName
		}
		if password != "" {
			data, _ := libs.AesEncrypt([]byte(password), []byte(Aeskey))
			pwdstr := base64.StdEncoding.EncodeToString(data)
			u["password"] = pwdstr
		}
		if headImg != "" {
			u["headImg"] = headImg
		}
		if carsetid > 0 {
			var carset model.CarSet
			dal.DB.First(&carset, carsetid)
			if &carset != nil && carset.ID > 0 {
				u["car_set_id"] = carsetid
			}
		}
		dal.DB.Model(&user).Updates(u)
	}
	return nil
}
func (logic LogicBase) ForgetPwd(mobile, code, pwd string) (err error) {
	err = logic.CheckCodeByMobile(username, code)
	if err != nil {
		return err
	}
}
