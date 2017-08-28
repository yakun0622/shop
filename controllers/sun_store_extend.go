package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunStoreExtend
type SunStoreExtendController struct {
	BaseController
}

func (c *SunStoreExtendController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create SunStoreExtend
// @Param	body		body 	models.SunStoreExtend	true		"body for SunStoreExtend content"
// @Success 201 {int} models.SunStoreExtend
// @Failure 403 body is empty
// @router / [post]
func (c *SunStoreExtendController) Post() {
	var v models.SunStoreExtend
	if err := json.Unmarshal(c.GetDataBytes(), &v); err == nil {
		v.Id = int(c.GetCurrentStoreID())
		if _, err := models.AddSunStoreExtend(&v); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v.Express
		} else {
			c.apiResult.Code = constant.OK
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.FAIL
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Get
// @Description get SunStoreExtend by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunStoreExtend
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunStoreExtendController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunStoreExtendById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunStoreExtend
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunStoreExtend
// @Failure 403
// @router / [get]
func (c *SunStoreExtendController) GetAll() {
	query, fields, sortby, order, offset, limit, err := c.GetAllSeachParams()
	if err != nil {
		c.apiResult.Code = constant.InvalidKeyValue
		c.apiResult.Error = err
	} else {
		query["Id"] = strconv.Itoa(int(c.GetCurrentStoreID()))
		l, err := models.GetAllSunStoreExtend(query, fields, sortby, order, offset, limit)
		if err != nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Error = err
		} else {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = l
		}
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunStoreExtend
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunStoreExtend	true		"body for SunStoreExtend content"
// @Success 200 {object} models.SunStoreExtend
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunStoreExtendController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunStoreExtend{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunStoreExtendById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the SunStoreExtend
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunStoreExtendController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunStoreExtend(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
