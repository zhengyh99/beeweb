package controllers

import (
	"beeweb/models"
	"strconv"

	"github.com/astaxie/beego"
)

//CategoryController 主页控制
type CategoryController struct {
	beego.Controller
}

//Get 客户端GET请求
func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		title := c.Input().Get("title")
		if len(title) == 0 {
			return
		}
		err := models.AddCategory(title)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			return
		}
		cid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
			return
		}
		models.DelCategory(cid)
		c.Redirect("/category", 302)
		return

	}
	c.Data["categories"] = models.GetAllCategory()
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsCategory"] = "true"
	c.TplName = "category.html"
}
