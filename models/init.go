package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type BaseModel struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// 表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "mfw_auth_" + defaultTableName;
	}
	// 打开数据库连接
	db, err := gorm.Open("mysql", "web:123456@/test?charset=utf8&parseTime=True&loc=Local")

	if err == nil {
		//init casbin
		InitCasbin()
		DB = db
		db.LogMode(true)
		db.AutoMigrate(&Operation{}, &App{}, &Resource{}, &Auth{},&UserAuth{})
		db.Model(&Auth{}).AddForeignKey("resource_id", "mfw_auth_resources(id)", "CASCADE", "NO ACTION")
		db.Model(&UserAuth{}).AddForeignKey("auth_id", "mfw_auth_auths(id)", "CASCADE", "NO ACTION")

		return db, err
	}
	return nil, err
}
