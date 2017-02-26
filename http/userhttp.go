package http

import (
	"fca/model"
	"fmt"
	"strconv"

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
	WriterResponse(w, 1, "", user)
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
