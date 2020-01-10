package controllers

import (
	"beeweb/models"
	"strconv"

	"github.com/astaxie/beego"
)

//Topic 文章控制
type TopicController struct {
	beego.Controller
}

func (t *TopicController) Get() {
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	t.Data["IsTopic"] = true
	t.Data["topics"] = models.GetAllTopic(false)
	t.TplName = "topic.html"
}
func (t *TopicController) Add() {

	t.Data["IsTopic"] = true
	t.TplName = "topic_add.html"

}

func (t *TopicController) Post() {

	if !checkAccount(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	op := t.Input().Get("op")
	title := t.Input().Get("title")
	content := t.Input().Get("content")
	if op == "update" {
		id, err := strconv.ParseInt(t.Input().Get("tid"), 10, 64)
		if err == nil {
			models.UpdateTopic(id, title, content)
		}
		beego.Error(err)
	} else {
		models.AddTopic(title, content)
	}

	t.Redirect("/topic", 302)
}

//View 客户端GET请求
func (t *TopicController) View() {
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	t.Data["IsTopic"] = true
	t.TplName = "topic_view.html"
	id, err := strconv.ParseInt(t.Ctx.Input.Param("0"), 10, 64)
	if err != nil {
		beego.Error(err)
		t.Redirect("/topic", 302)
		return
	}

	t.Data["topic"] = models.GetTopic(id)
}

func (t *TopicController) Del() {
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	t.Data["IsTopic"] = true
	t.TplName = "topic_view.html"
	id, err := strconv.ParseInt(t.Ctx.Input.Param("0"), 10, 64)
	if err != nil {
		beego.Error(err)

	} else {
		models.DelTopic(id)
	}

	t.Redirect("/topic", 302)
}
