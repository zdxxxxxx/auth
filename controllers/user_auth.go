package controllers

import (
	"github.com/gin-gonic/gin"
	"auth2/utils"
	"auth2/models"
	"strconv"
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
		a, _ := models.GetAuthById(uint(i))
		r, _ := models.GetResourceById(a.ResourceId)
		app, _ := models.GetAppById(r.AppId)
		o, _ := models.GetOperationById(a.OperationId)
		ua := &models.UserAuth{Uid: uid, AuthId: uint(i), AppId: app.Id, Path: r.Path}
		ua.Insert()
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

func DeleteUserAuth(c *gin.Context) {
	var req = new(utils.ReqData)
	id, _err1 := strconv.Atoi(c.Param("id"))
	// 参数判断
	if _err1 != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	ua, _ := models.GetUserAuthById(uint(id))
	a, _ := models.GetAuthById(ua.AuthId)
	r, _ := models.GetResourceById(a.ResourceId)
	o, _ := models.GetOperationById(a.OperationId)
	models.DeleteAuth(ua.Uid, r.Path, o.Value)
	err := ua.Delete()
	if err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func DeleteUserAuths(c *gin.Context) {
	var req = new(utils.ReqData)
	resData := &IdsData{}
	_err := c.BindJSON(resData)
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	var err error
	ids := resData.Ids
	for _, id := range ids {
		ua, _ := models.GetUserAuthById(uint(id))
		a, _ := models.GetAuthById(ua.AuthId)
		r, _ := models.GetResourceById(a.ResourceId)
		o, _ := models.GetOperationById(a.OperationId)
		models.DeleteAuth(ua.Uid, r.Path, o.Value)
		err = ua.Delete()
	}
	if err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func GetUserAuth(c *gin.Context) {
	var req = new(utils.ReqData)
	uid := c.Query("uid")
	app := c.Query("app")
	path := c.Query("path")
	appId, _ := strconv.Atoi(app)
	uas, err := models.GetUserAuth(uid, appId, path)
	var datas = make([]*utils.UserAuthReqData, 0)
	for _, ua := range uas {
		var data = new(utils.UserAuthReqData)
		auth, _ := models.GetAuthById(ua.AuthId)
		op, _ := models.GetOperationById(auth.OperationId)
		re, _ := models.GetResourceById(auth.ResourceId)
		app, _ := models.GetAppById(re.AppId)
		data.Id = int(ua.Id)
		data.Path = re.Path
		data.Uid = ua.Uid
		data.App = app.Name
		data.Operation = op.Name
		datas = append(datas, data)
	}
	if err == nil {
		req.SetResult(0, datas)
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}
