package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Topic 标题
type Topic struct {
	ID              int64
	UID             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserID int64
}

func AddTopic(title, content string) bool {

	topic := &Topic{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}
	o := orm.NewOrm()
	_, err := o.Insert(topic)
	if err != nil {
		beego.Error(err)
		return false
	}
	return true

}

func GetAllTopic(isDesc bool) []*Topic {
	topics := make([]*Topic, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	if isDesc {
		_, err := qs.OrderBy("-created").All(&topics)
		if err != nil {
			beego.Error(err)
			return nil
		}
	} else {
		_, err := qs.All(&topics)
		if err != nil {
			beego.Error(err)
			return nil
		}
	}

	return topics
}
