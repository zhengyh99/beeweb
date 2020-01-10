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

func GetTopic(id int64) *Topic {
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err := qs.Filter("id", id).One(topic)
	if err != nil {
		return nil
	}
	topic.Views++
	_, err = o.Update(topic)
	if err != nil {
		beego.Error(err)
		return nil
	}
	return topic
}

func UpdateTopic(id int64, title, content string) bool {
	o := orm.NewOrm()
	topic := new(Topic)
	topic.ID = id
	topic.Title = title
	topic.Content = content
	topic.Updated = time.Now()
	_, err := o.Update(topic)
	if err != nil {
		beego.Error(err)
		return false
	}
	return true
}

func DelTopic(id int64) bool {
	o := orm.NewOrm()
	topic := new(Topic)
	topic.ID = id
	_, err := o.Delete(topic)
	if err != nil {
		beego.Error(err)
		return false
	}
	return true
}
