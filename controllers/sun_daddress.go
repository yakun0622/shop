package controllers

import (
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunDaddress
type SunDaddressController struct {
	BaseController
}

func (c *SunDaddressController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("SetDefault", c.SetDefault)
}

// @Title Post
// @Description create SunDaddress
// @Param	body		body 	models.SunDaddress	true		"body for SunDaddress content"
// @Success 201 {int} models.SunDaddress
// @Failure 403 body is empty
// @router / [post]
func (c *SunDaddressController) Post() {
	storeId := int(c.GetCurrentStoreID())
	store, _ := models.GetSunStoreById(storeId)
	v := models.SunDaddress{StoreId: uint32(storeId), IsDefault: 2, Company: store.StoreCompanyName}
	c.BindModelWithPost(&v)
}

// @Title Get
// @Description get SunDaddress by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunDaddress
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunDaddressController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunDaddressById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunDaddress
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunDaddress
// @Failure 403
// @router / [get]
func (c *SunDaddressController) GetAll() {
	query, fields, sortby, order, offset, limit, err := c.GetAllSeachParams()
	if err != nil {
		c.apiResult.Code = constant.InvalidKeyValue
		c.apiResult.Error = err
	} else {
		query["StoreId"] = strconv.Itoa(int(c.GetCurrentStoreID()))
		l, err := models.GetAllSunDaddress(query, fields, sortby, order, offset, limit)
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
// @Description update the SunDaddress
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunDaddress	true		"body for SunDaddress content"
// @Success 200 {object} models.SunDaddress
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunDaddressController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunDaddress{Id: id, StoreId: uint32(c.GetCurrentStoreID())}
	c.BindModelWithPut(&v, "SellerName", "Address", "Telphone", "AreaId", "CityId", "AreaInfo")
}

// @Title Delete
// @Description 删除当前登录商户的特定发货地址
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunDaddressController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if result, err := models.DeleteSunDaddressByStore(id, c.GetCurrentStoreID()); err == nil {
		if result {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.apiResult.Code = constant.FAIL
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title SetDefault
// @Description set default daddress
// @Param	id		path 	string	true		"The id you want to set default"
// @Success 200 {object} models.SunDaddress
// @Failure 403 :id is not int
// @router /default/:id [put]
func (c *SunDaddressController) SetDefault() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if models.SetDefaultDAddress(id, c.GetCurrentStoreID()) {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}
