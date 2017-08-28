package controllers

import (
	"encoding/json"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for TransportExtend
type TransportExtendController struct {
	BaseController
}

func (c *TransportExtendController) URLMapping() {
	c.Mapping("Save", c.Save)
	c.Mapping("Add", c.Add)
	c.Mapping("Delete", c.Delete)
}

// @Title Delete
// @Description delete the TransportExtend
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete [post]
func (c *TransportExtendController) Delete() {
	id, _ := c.GetInt("id")
	beego.Info("删除ID：" + strconv.Itoa(id))
	if id > 0 {
		if err := models.DeleteSunTransportExtend(id); err == nil {
			c.HandleResult(err, "删除运费模板成功", constant.OK, true, 0)
		} else {
			c.HandleResult(err, "删除运费模板失败", constant.FAIL, false, 0)
		}
	} else {
		c.HandleResult(nil, "删除临时运费模板，id为空...", constant.OK, true, 0)
	}
}

// @Title 保存
// @Description create TransportExtend
// @Param	body		body 	models.TransportExtend	true		"body for TransportExtend content"
// @Success 201 {int} models.TransportExtend
// @Failure 403 body is empty
// @router /save [post]
func (c *TransportExtendController) Save() {
	var v models.TransportExtend
	//beego.Info("json", c.GetString("data"))
	if err := json.Unmarshal(c.GetDataBytes(), &v); err == nil {
		if v.Id > 0 {
			if err := models.UpdateSunTransportExtendById(&v); err == nil {
				c.HandleResult(err, "修改运费模板成功", constant.OK, v.Id, 0)
			} else {
				c.HandleResult(err, "修改运费模板失败", constant.FAIL, false, 0)
			}
		} else {
			if id, err := models.AddTransportExtend(&v); err == nil {
				c.HandleResult(err, "新增运费模板成功", constant.OK, id, 0)
			} else {
				c.HandleResult(err, "新增运费模板失败", constant.FAIL, false, 0)
			}
		}

	} else {
		c.HandleResult(err, "数据解析失败", constant.FAIL, false, 0)
	}
}

// @Title 添加运费模板扩展信息，不做业务处理，方便前端获取index
// @router /add [post]
func (c *TransportExtendController) Add() {
	beego.Info("json", c.GetString("data"))
	c.HandleResult(nil, "修改运费模板成功", constant.OK, true, 0)
}
