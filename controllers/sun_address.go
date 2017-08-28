package controllers

import (
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunAddress
type SunAddressController struct {
	BaseController
}

func (c *SunAddressController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create SunAddress
// @Param	body		body 	models.SunAddress	true		"body for SunAddress content"
// @Success 201 {int} models.SunAddress
// @Failure 403 body is empty
// @router / [post]
func (c *SunAddressController) Post() {
	v := models.SunAddress{MemberId: uint32(c.GetUserID())}
	c.BindModelWithPost(&v)
}

// @Title Get
// @Description get SunAddress by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunAddress
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunAddressController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunAddressById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunAddress
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunAddress
// @Failure 403
// @router / [get]
func (c *SunAddressController) GetAll() {
	query, fields, sortby, order, offset, limit, err := c.GetAllSeachParams()
	if err != nil {
		c.apiResult.Code = constant.InvalidKeyValue
		c.apiResult.Error = err
	} else {
		query["MemberId"] = strconv.Itoa(int(c.GetUserID()))
		l, err := models.GetAllSunAddress(query, fields, sortby, order, offset, limit)
		if err != nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = nil
		} else {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = l
		}
	}

	c.ServeJSON()
}

// @Title Update
// @Description update the SunAddress
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunAddress	true		"body for SunAddress content"
// @Success 200 {object} models.SunAddress
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunAddressController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunAddress{Id: id}
	c.BindModelWithPut(&v, "TrueName", "AreaInfo", "AreaId", "CityId", "Address", "TelPhone", "MobPhone", "IsDefault")
}

// @Title Delete
// @Description delete the SunAddress
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunAddressController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunAddress{Id: id}
	c.Remove(&v)
}
