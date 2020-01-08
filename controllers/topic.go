package controllers

import (
	"github.com/astaxie/beego"
)

//Topic 文章控制
type TopicController struct {
	beego.Controller
}

func (t *TopicController) Get() {
	t.Data["IsTopic"] = true
	t.TplName = "topic.html"
}
func (t *TopicController) Add() {
	t.Data["IsTopic"] = true
	t.TplName = "topic_add.html"
}
