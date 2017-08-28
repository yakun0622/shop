package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunCart
type SunCartController struct {
	BaseController
}

func (c *SunCartController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create SunCart
// @Param	body		body 	models.SunCart	true		"body for SunCart content"
// @Success 201 {int} models.SunCart
// @Failure 403 body is empty
// @router / [post]
func (c *SunCartController) Post() {
	var v models.SunCart
	if err := json.Unmarshal([]byte(c.GetString("data")), &v); err == nil {
		//校验请求数据是否跟登录用户一致
		if !c.ValidUser(int(v.BuyerId)) {
			return
		}
		count := v.GoodsNum
		if exist := models.IsExistInCart(&v); exist {
			v.GoodsNum = v.GoodsNum + count
			if err = models.UpdateSunCartById(&v); err == nil {
				c.apiResult.Code = constant.OK
				c.apiResult.Data = v
			} else {
				c.apiResult.Code = constant.ActionFaild
				c.apiResult.Error = err
			}
		} else {
			if _, err := models.AddSunCart(&v); err == nil {
				c.apiResult.Code = constant.OK
				c.apiResult.Data = v
			} else {
				c.apiResult.Code = constant.ActionFaild
				c.apiResult.Error = err
			}
		}
	} else {
		c.apiResult.Code = constant.InvalidRequestData
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Get
// @Description get SunCart by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunCart
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunCartController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//TODO:带入用户ID进行查找，无需进行校验
	v, err := models.GetSunCartById(id)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		if !c.ValidUser(int(v.BuyerId)) {
			return
		}
		c.apiResult.Code = constant.OK
		c.apiResult.Data = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description 获取当前登录用户的购物车信息
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunCart
// @Failure 403
// @router / [get]
func (c *SunCartController) GetAll() {
	if m, err := models.GetAllSunCart(c.GetUserID()); err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = m
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunCart
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunCart	true		"body for SunCart content"
// @Success 200 {object} models.SunCart
// @Failure 403 :id is not int
// @router /change_num [post]
func (c *SunCartController) Put() {
	id, _ := c.GetInt("Id")
	goodsNum, _ := c.GetInt16("GoodsNum")
	v := models.SunCart{Id: id, GoodsNum: uint16(goodsNum)}
	c.BindModelWithPut(&v, "GoodsNum")
}

// @Title Delete
// @Description delete the SunCart
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunCartController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if m, err := models.GetSunCartById(id); err == nil {
		if !c.ValidUser(int(m.BuyerId)) {
			return
		}
		if err := models.DeleteSunCart(id); err == nil {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.ActionFaild
		c.apiResult.Error = err
	}
	c.ServeJSON()
}
