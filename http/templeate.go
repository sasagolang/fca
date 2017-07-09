package http

import (
	"net/http"
	"time"

	"fmt"
	"html/template"
	"strings"
)

type LoginMsg struct {
	Errmsg string
}

func ExecuteHttp(w http.ResponseWriter, r *http.Request) {

	msg := LoginMsg{}
	t, err := template.ParseFiles("./tmpl/login.tpl")
	fmt.Printf("ExecuteHttp1:%v", err)
	err = t.Execute(w, msg)
	fmt.Printf("ExecuteHttp2:%v", err)
}
func LoginPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var (
			username string = r.PostFormValue("username")
			password string = r.PostFormValue("password")
		)
		if strings.ToLower(username) == "admin" && strings.ToLower(password) == "password" {
			COOKIE_MAX_MAX_AGE := time.Hour * 24 / time.Second // 单位：秒。
			maxAge := int(COOKIE_MAX_MAX_AGE)
			uid := "管理员"

			uid_cookie := &http.Cookie{
				Name:     "fcauid",
				Value:    uid,
				Path:     "/",
				HttpOnly: false,
				MaxAge:   maxAge}

			http.SetCookie(w, uid_cookie)
			http.Redirect(w, r, "/admin", http.StatusFound)
			fmt.Printf("LoginPost:%v", nil)
			return
		} else {
			msg := LoginMsg{Errmsg: "用户名密码错误"}
			t, _ := template.ParseFiles("./tmpl/login.tpl")
			//fmt.Printf("ExecuteHttp1:%v", err)
			_ = t.Execute(w, msg)
			//fmt.Printf("ExecuteHttp2:%v", err)
		}
	}
	//fmt.Printf("ExecuteHttp1:%v", r.Referer())
}
