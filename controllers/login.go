package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//LoginController 登陆控制
type LoginController struct {
	beego.Controller
}

//Get 登陆Get请求
func (l *LoginController) Get() {
	isExit := l.Input().Get("exit") == "true"
	if isExit {
		l.Ctx.SetCookie("uname", "", -1, "/")
		l.Ctx.SetCookie("pwd", "", -1, "/")
		l.Redirect("/", 302)
		return
	}

	l.Data["IsHome"] = "true"
	l.TplName = "login.html"
}

//Post 登陆Post请求
func (l *LoginController) Post() {

	uname := l.Input().Get("username")
	pwd := l.Input().Get("password")
	autologin := l.Input().Get("autologin") == "on"

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autologin {
			maxAge = 1<<32 - 1
		}
		l.Ctx.SetCookie("uname", uname, maxAge, "/")
		l.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	l.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	uname, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	pwd, err := ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	return beego.AppConfig.String("uname") == uname.Value &&
		beego.AppConfig.String("pwd") == pwd.Value
}
