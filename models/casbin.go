package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
)

var e *casbin.Enforcer

func InitCasbin() {
	var adapter = gormadapter.NewAdapter("mysql", "web:123456@/test?charset=utf8&parseTime=True&loc=Local", true)
	e = casbin.NewEnforcer("conf/rbac_model.conf", adapter)
}

func CheckAuth(params ...interface{}) {
	// casbin init
	e.LoadPolicy()
	e.Enforce(params...)
}

func AddAuth(params ...interface{}) {
	e.AddPolicy(params...)
}

func DeleteAuth(params ...interface{}) {
	e.RemovePolicy(params...)
}
