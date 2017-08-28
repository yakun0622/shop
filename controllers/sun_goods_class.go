package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"

	"github.com/yakun0622/shop/redis"
)

// oprations for SunGoodsClass
type SunGoodsClassController struct {
	BaseController
}

func (c *SunGoodsClassController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create SunGoodsClass
// @Param	body		body 	models.SunGoodsClass	true		"body for SunGoodsClass content"
// @Success 201 {int} models.SunGoodsClass
// @Failure 403 body is empty
// @router / [post]
func (c *SunGoodsClassController) Post() {
	var v models.SunGoodsClass
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSunGoodsClass(&v); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v
		} else {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.InvalidRequestData
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Get
// @Description get SunGoodsClass by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunGoodsClass
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunGoodsClassController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	goodsClass := redis.Instance().Get("goodsClass." + idStr)
	if goodsClass != nil {
		var dataModel models.SunGoodsClass
		json.Unmarshal(goodsClass.([]byte), &dataModel)
		c.apiResult.Code = constant.OK
		c.apiResult.Data = dataModel
	} else {
		v, err := models.GetSunGoodsClassById(id)
		if err != nil {
			c.apiResult.Code = constant.FindNoData
			c.apiResult.Error = err
		} else {
			resultJSON, _ := json.Marshal(v)
			redis.Instance().Put("goodsClass."+idStr, string(resultJSON), 1000*time.Second)
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v
		}
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunGoodsClass
// @Success 200 {object} models.SunGoodsClass
// @Failure 403
// @router / [get]
func (c *SunGoodsClassController) GetAll() {
	goodsClass := redis.Instance().Get("allGoodsClass")

	if goodsClass != nil {
		var dataModel []models.SunGoodsClass
		json.Unmarshal(goodsClass.([]byte), &dataModel)
		c.apiResult.Code = constant.OK
		c.apiResult.Data = dataModel
	} else {
		l, err := models.GetAllSunGoodsClassSimple()
		if err != nil {
			c.apiResult.Code = constant.FindNoData
			c.apiResult.Error = err
		} else {
			resultJSON, _ := json.Marshal(l)
			redis.Instance().Put("allGoodsClass", string(resultJSON), 1000*time.Hour)
			c.apiResult.Code = constant.OK
			c.apiResult.Data = l
		}
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunGoodsClass
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunGoodsClass	true		"body for SunGoodsClass content"
// @Success 200 {object} models.SunGoodsClass
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunGoodsClassController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunGoodsClass{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunGoodsClassById(&v); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v
		} else {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.InvalidRequestData
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the SunGoodsClass
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunGoodsClassController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunGoodsClass(id); err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.ActionFaild
		c.apiResult.Error = err
	}
	c.ServeJSON()
}
