package controllers

import (
	"auth2/utils"
	"github.com/gin-gonic/gin"
	"auth2/models"
	"strconv"
	"fmt"
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
	path := utils.BuildPath(data.Path)
	ops := data.Operations
	d := &reqResourceData{}
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

	a, _ := models.DeleteAuthByResourceId(uint(_id))
	models.DeleteUserAuthByAuth(a)
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
	path := utils.BuildPath(data.Path)
	ops := data.Operations
	err := models.UpdateResources(uint(_id), path)
	if err == nil {
		auths, _ := models.GetAuthsByResourceId(uint(_id));
		var old_ops []int
		for _, item := range auths {
			old_ops = append(old_ops, item.OperationId);
		}
		remove_ops := utils.DiffIntSlice(old_ops, ops);
		add_ops := utils.DiffIntSlice(ops, old_ops);
		for _, remove_id := range remove_ops {
			a, _ := models.DeleteAuthByRAndO(uint(_id), uint(remove_id));
			models.DeleteUserAuthByAuth(a)
		}
		for _, add_id := range add_ops {
			a := &models.Auth{ResourceId: uint(_id), OperationId: add_id}
			a.Insert()
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
	data := &reqResourceData{ResourceId: r.Id, Path: r.Path, Operations: ops}
	req.SetResult(0, data)
	c.JSON(200, req)
}

func GetAuths(c *gin.Context) {
	var req = new(utils.ReqData)
	id := c.Param("id");
	_id, _err := strconv.ParseUint(id, 10, 64)
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	rs, _ := models.GetResourceByPath(uint(_id))
	fmt.Println(rs)
	var auths []*reqAuthData
	for _, item := range rs {
		as, _ := models.GetAuthsByResourceId(uint(item.Id))
		for _, a := range as {
			o, _ := models.GetOperationById(a.OperationId)
			auths = append(auths, &reqAuthData{AuthId: a.Id, Path: item.Path, Operation: o.Name})
		}
	}
	req.SetResult(0, auths)
	c.JSON(200, req)
}
