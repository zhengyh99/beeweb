package controllers

import (
	"github.com/astaxie/beego"
)

//HomeController 主页控制
type HomeController struct {
	beego.Controller
}

//Get 客户端GET请求
func (h *HomeController) Get() {
	h.Data["IsLogin"] = checkAccount(h.Ctx)
	h.Data["IsHome"] = true
	h.TplName = "home.html"
}
