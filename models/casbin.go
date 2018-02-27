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

func CheckAuth(params ...interface{}) bool {
	// casbin init
	e.LoadPolicy()
	return e.Enforce(params...)
}

func AddAuth(params ...interface{}) bool {
	return e.AddPolicy(params...)
}

func DeleteAuth(params ...interface{}) {
	e.RemovePolicy(params...)
}
