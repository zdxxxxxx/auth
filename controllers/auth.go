package controllers

import (
	"auth2/utils"
	"github.com/gin-gonic/gin"
	"auth2/models"
	"strconv"
)

func CreateAuth(c *gin.Context) {
	var req = new(utils.ReqData)
	data := &AuthClientJson{}
	_err := c.BindJSON(data)
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	app_id := data.AppId
	path := data.Path
	ops := data.Operations
	d := &reqAuthsData{}
	r := &models.Resource{AppId: uint(app_id), Path: path}
	err := r.Insert()
	if err == nil {
		for _, item := range ops {
			a := &models.Auth{ResourceId: uint(r.Id), OperationId: item}
			e := a.Insert()
			if e != nil {
				panic(e)
				break
			}
		}
		d.ResourceId = r.Id
		d.Path = r.Path
		d.Operations = ops
		req.SetResult(0, d)
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func DeleteAuth(c *gin.Context) {
	var req = new(utils.ReqData)
	id := c.Param("id")
	_id, _err1 := strconv.ParseUint(id, 10, 64)

	// 参数判断
	if _err1 != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	r := new(models.Resource)
	r.Id = uint(_id)
	err := models.DeleteResources(uint(_id))
	if err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func UpdateAuth(c *gin.Context) {
	var req = new(utils.ReqData)
	id := c.Param("id")
	_id, _err1 := strconv.ParseUint(id, 10, 64)
	data := &AuthClientJson{}
	_err := c.BindJSON(data)
	// 参数判断
	if _err1 != nil || _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	path := data.Path
	ops := data.Operations
	err := models.UpdateResources(uint(_id), path)
	if err == nil {
		models.DeleteAuthByResourceId(uint(_id))
		for _, item := range ops {
			a := &models.Auth{ResourceId: uint(_id), OperationId: item}
			e := a.Insert()
			if e != nil {
				panic(e)
				break
			}
		}
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func GetAuth(c *gin.Context) {
	var req = new(utils.ReqData)
	id := c.Param("id");
	_id, _err := strconv.ParseUint(id, 10, 64)
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	r, _ := models.GetResourceById(uint(_id))
	auths, _ := models.GetAuthsByResourceId(uint(_id))
	var ops []int
	for _, i := range auths {
		ops = append(ops, i.OperationId)
	}
	data := &reqAuthsData{ResourceId: r.Id, Path: r.Path, Operations: ops}
	req.SetResult(0, data)
	c.JSON(200, req)
}
