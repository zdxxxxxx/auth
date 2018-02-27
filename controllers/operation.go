package controllers

import (
	"github.com/gin-gonic/gin"
	"auth2/models"
	"auth2/utils"
	"strconv"
)

func GetOperations(c *gin.Context) {
	var req = new(utils.ReqData)
	o, err := models.GetOperationAll()
	if err == nil {
		req.SetResult(0, o)
	} else {
		req.SetResult(100, err)
	}
	c.JSON(200, req)
}

func GetOperation(c *gin.Context) {
	var req = new(utils.ReqData)
	id, _err := strconv.Atoi(c.Param("id"))
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	o, err := models.GetOperationById(id)
	if err == nil {
		req.SetResult(0, o)
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func CreateOperation(c *gin.Context) {
	var req = new(utils.ReqData)
	data := &OperationClientJson{}
	_err := c.BindJSON(data)
	// 参数判断
	if _err != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	name := data.Name
	value := data.Value
	o := &models.Operation{Name: name, Value: value}
	err := o.Insert()
	if err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func UpdateOperation(c *gin.Context) {
	var req = new(utils.ReqData)
	data := &OperationClientJson{}
	_err := c.BindJSON(data)
	id, _err1 := strconv.Atoi(c.Param("id"))
	// 参数判断
	if _err != nil || _err1 != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	name := data.Name
	value := data.Value
	o := &models.Operation{Id: id, Name: name, Value: value}
	err := o.Update()
	if err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}

func DeleteOperation(c *gin.Context) {
	var req = new(utils.ReqData)
	id, _err1 := strconv.Atoi(c.Param("id"))
	// 参数判断
	if _err1 != nil {
		req.SetResult(101, []int{})
		c.JSON(200, req)
		return
	}
	o := &models.Operation{Id: id}
	err := o.Delete()
	if err == nil {
		req.SetResult(0, []int{})
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}
