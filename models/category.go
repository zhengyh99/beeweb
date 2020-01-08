package models

import (
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

//Category 文章分类
type Category struct {
	ID              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserID int64
}

func AddCategory(title string) error {
	o := orm.NewOrm()
	ctg := &Category{Title: title}
	qs := o.QueryTable("category")
	err := qs.Filter("title", title).One(ctg)
	if err == nil {
		beego.Error(err)
		return err
	}
	_, err = o.Insert(ctg)
	if err != nil {
		beego.Error(err)
		return err
	}
	return nil
}

func GetAllCategory() []*Category {
	o := orm.NewOrm()
	ctgs := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&ctgs)
	if err != nil {
		beego.Error(err)
		return nil
	}
	beego.Info(ctgs)
	return ctgs
}

func DelCategory(id int64) error {
	o := orm.NewOrm()
	ctg := &Category{ID: id}
	_, err := o.Delete(ctg)
	return err
}
