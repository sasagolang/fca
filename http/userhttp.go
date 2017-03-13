package http

import (
	"encoding/base64"
	"fca/model"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetUser(w rest.ResponseWriter, r *rest.Request) {
	uid, _ := strconv.Atoi(r.PathParam("uid"))
	fmt.Printf("Logic:%v\n", Logic)
	user, _ := Logic.GetUserByUid(uid)
	fmt.Printf("user:%v\n", user)
	result := Result{}
	result.ResultCode = 1
	result.ResultMessage = ""
	result.Data = user
	w.WriteJson(result)
}
func SendCode(w rest.ResponseWriter, r *rest.Request) {
	mobile := r.PathParam("mobile")
	d := make([]byte, 4)
	s := NewLen(4)
	ss := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		ss += strconv.FormatInt(int64(d[v]), 32)
	}
	err := Logic.SendMsgByMobile(mobile, ss)
	if err != nil {
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", "")
}
func UserLogin(w rest.ResponseWriter, r *rest.Request) {
	request := model.UserLoginRequest{}
	err := r.DecodeJsonPayload(&request)
	user, err := Logic.UserLogin(request.UserName, request.Password, request.Code, 1)
	//w.WriteJson(user)
	if err != nil {
		fmt.Printf("UserLogin error(%v)\n", err)
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	res := model.UserLoginResponse{}
	res.Token = strconv.FormatInt(time.Now().UnixNano(), 10)
	res.User = *user
	WriterResponse(w, 1, "", res)
}
func RegisterUser(w rest.ResponseWriter, r *rest.Request) {
	user := model.User{}
	err := r.DecodeJsonPayload(&user)
	fmt.Printf("RegisterUser:%v\n", user)

	u, err := Logic.RegisterUser(user.Mobile, user.Email, user.Name, user.Password)
	//	user := Logic.UserLogin("aaa", "bbb", "ccc")
	result := Result{}
	result.ResultCode = 1
	result.ResultMessage = ""
	result.Data = err
	if err != nil {
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", u)
	//w.WriteJson(result)
}
func RegisterByMobile(w rest.ResponseWriter, r *rest.Request) {
	reg := model.RegisterUserRequest{}
	err := r.DecodeJsonPayload(&reg)
	if err != nil {
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	u, err := Logic.RegisterUser(reg.Mobile, "", "anonymous", reg.Pwd)
	if err != nil {
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", u)
}
func UpdateUserInfoFunc(w rest.ResponseWriter, r *rest.Request) {

	info := model.UpdateUserInfoRequest{}
	err := r.DecodeJsonPayload(&info)
	if err != nil {
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	if info.HeadImg != "" {
		filename := "/img/headimg/" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".png" //成图片文件并把文件写入到buffer

		ddd, err3 := base64.StdEncoding.DecodeString(info.HeadImg) //成图片文件并把文件写入到buffer
		if err3 != nil {
			panic(err3)
		}
		err2 := ioutil.WriteFile(getCurrentPath()+filename, ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		if err2 != nil {
			//panic(err2)
			fmt.Printf("UpdateUserInfoFunc:%v\n", err2)
		} else {
			info.HeadImg = filename
		}
	}
	err = Logic.UpdateUserInfo(info.UID, info.NickName, info.HeadImg, info.Password, info.CarSetID)
	if err != nil {
		WriterResponse(w, 2, err.Error(), err)
		return
	}
	WriterResponse(w, 1, "", "")
}
