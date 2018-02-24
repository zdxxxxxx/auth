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
	}
	o, err := models.GetOperationById(id)
	if err == nil {
		req.SetResult(0, o)
	} else {
		req.SetResult(100, err.Error())
	}
	c.JSON(200, req)
}
