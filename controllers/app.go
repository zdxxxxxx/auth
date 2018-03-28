package controllers

import (
	"github.com/gin-gonic/gin"
	"auth2/utils"
	"auth2/models"
	"strconv"
	"fmt"
)

func GetApps(c *gin.Context) {
	var req = new(utils.ReqData)
	o, err := models.GetAppAll()
	if err == nil {
		req.SetResult(0, o)
	} else {
		req.SetResult(100, err)
	}
	c.JSON(200, req)
}

func GetAppSelect(c *gin.Context) {
	var req = new(utils.ReqData)
	var datas = make([]*utils.AppReqData, 0)
	o, err := models.GetAppAll()
	for _, item := range o {
		datas = append(datas, &utils.AppReqData{Id: int(item.Id), Name: item.Name})
	}
	if err == nil {
		req.SetResult(0, datas)
	} else {
		req.SetResult(100, err)
	}
	c.JSON(200, req)
}
func GetApp(c *gin.Context) {
	var req = new(utils.ReqData)
	id, _err := strconv.Atoi(c.Param("id"))
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	o, err := models.GetAppById(uint(id))
	r, err := models.GetResourceById(uint(o.ResourceId))
	m := &reqAppData{Name: o.Name, Path: r.Path, Desc: o.Desc, Content: o.Content}
	if err == nil {
		req.SetResult(0, m)
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}
func CreateApp(c *gin.Context) {
	var req = new(utils.ReqData)
	data := &AppClientJson{}
	_err := c.BindJSON(data)
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	name := data.Name
	desc := data.Desc
	path := utils.BuildPath(data.Path)
	o := &models.App{Name: name, Desc: desc}
	err := o.Insert()
	if err == nil {
		r := &models.Resource{AppId: o.Id, Path: path}
		_err := r.Insert()
		fmt.Println(r.Id)
		o.ResourceId = r.Id
		_err1 := o.UpdateResourceId()
		if _err == nil && _err1 == nil {
			req.SetResult(0, o)
		} else {
			req.SetResult(100, _err.Error())
		}
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func UpdateApp(c *gin.Context) {
	var req = new(utils.ReqData)
	id := c.Param("id");
	data := &AppClientJson{}
	_err := c.BindJSON(data)
	_id, _err1 := strconv.ParseUint(id, 10, 64)
	// 参数判断
	if _err != nil || _err1 != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	name := data.Name
	desc := data.Desc
	o := &models.App{Name: name, Desc: desc}
	o.Id = uint(_id)
	err := o.Update()
	if err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func DeleteApp(c *gin.Context) {
	var req = new(utils.ReqData)
	id := c.Param("id");
	_id, _err1 := strconv.ParseUint(id, 10, 64)
	// 参数判断
	if _err1 != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	o := new(models.App)
	o.Id = uint(_id)
	err := o.Delete()
	_err := models.DeleteResourceByAppId(o.Id)
	if err == nil && _err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func UpdateAppContent(c *gin.Context) {
	var req = new(utils.ReqData)
	id := c.Param("id");
	data := &AppClientJson{}
	_err := c.BindJSON(data)
	_id, _err1 := strconv.ParseUint(id, 10, 64)
	// 参数判断
	if _err != nil || _err1 != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	content := data.Content
	o := &models.App{Content: content}
	o.Id = uint(_id)
	err := o.UpdateContent()
	if err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}
