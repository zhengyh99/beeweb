package controllers

import (
	"github.com/astaxie/beego"
)

//LoginController 登陆控制
type LoginController struct {
	beego.Controller
}

//Get 登陆Get请求
func (l *LoginController) Get() {

	l.Data["IsHome"] = "true"
	l.TplName = "login.html"
}
