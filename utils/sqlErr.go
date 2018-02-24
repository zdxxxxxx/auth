package utils

import "github.com/astaxie/beego/orm"

func GetErrInfo(Error error) (int, string) {
	var err string
	var code int
	if (Error == orm.ErrNoRows) {
		code = 201
		err = "未发现查询结果"
	}
	return code, err
}
