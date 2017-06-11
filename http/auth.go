package http

import (
	"html/template"
	"net/http"

	"fmt"

	"github.com/qor/admin"
	"github.com/qor/qor"
)

type Auth struct {
}
type AdminUser struct {
	Name string
}

func InitAuth() {
	http.HandleFunc("/auth/login", AuthView)
}
func AuthView(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template")
	//template.ParseGlob
	s, err := t.ParseFiles(GetCurrentPath() + "/view/login.html.tpl")
	//	err := t.Execute(w, nil)
	fmt.Printf("%v,%v\n", s, err)

}
func (Auth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	if userid, err := c.Request.Cookie("userid"); err == nil || true {
		//var user  auth
		user := AdminUser{}
		user.Name = userid.Value
		return user
	}

	return nil
}

func (u AdminUser) DisplayName() string {
	return u.Name
}
