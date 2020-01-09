package controllers

import (
	"beeweb/models"

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
	title := t.Input().Get("title")
	content := t.Input().Get("content")
	models.AddTopic(title, content)

	t.Redirect("/topic", 302)
}
