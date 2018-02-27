package controllers

import (
	"github.com/gin-gonic/gin"
	"auth2/utils"
	"auth2/models"
	"fmt"
)

func CreateUserAuth(c *gin.Context) {
	var req = new(utils.ReqData)
	data := &UserAuthClientJson{}
	_err := c.BindJSON(data)
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	var e bool
	uid := data.Uid
	auth := data.Auth
	for _, i := range auth {
		ua := &models.UserAuth{Uid: uid, AuthId: uint(i)}
		ua.Insert()
		a, _ := models.GetAuthById(uint(i))
		r, _ := models.GetResourceById(a.ResourceId)
		o, _ := models.GetOperationById(a.OperationId)
		fmt.Println(uid, r.Path, o.Value)
		e = models.AddAuth(uid, r.Path, o.Value)
	}
	if e {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, "授权失败！")
	}
	c.JSON(200, req)
}

func CheckAuth(c *gin.Context) {
	var req = new(utils.ReqData)
	data := &CheckAuthClientJson{}
	_err := c.BindJSON(data)
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	uid := data.Uid
	path := data.Path
	op := data.Operation
	result := models.CheckAuth(uid, path, op)
	req.SetResult(0, result)
	c.JSON(200, req)
}
