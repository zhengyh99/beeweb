package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DBName    = "root:123@tcp(127.0.0.1:3306)/beeblog"
	_SQLDriver = "mysql"
)

func RegisterDB() {

	orm.RegisterModel(new(Topic), new(Category))
	orm.RegisterDataBase("default", _SQLDriver, _DBName, 10)

}
