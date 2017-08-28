package controllers

import (
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/constant/permission"
	"github.com/yakun0622/shop/models"
)

type Role struct {
	BaseController
}

func (c *Role) URLMapping() {
	c.Mapping("Post", c.Post)
	// c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Lock", c.Lock)
	c.Mapping("UnLock", c.UnLock)
	c.Mapping("Permission", c.Permission)
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
func (c *Role) GetAll() {
	query, fields, sortby, order, offset, limit, err := c.GetAllSeachParams()
	if err != nil {
		c.apiResult.Code = constant.InvalidKeyValue
		c.apiResult.Error = err
	} else {
		l, err := models.GetAllRole(query, fields, sortby, order, offset, limit)
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

// @Title Post
// @router / [post]
func (r *Role) Post() {
	v := models.Role{}
	v.RoleCtime = r.GetCurrentTime()
	v.RoleLock = 1
	r.BindModelWithPost(&v)
}

// @Title Get All
// @Description get permission
// @Failure 403
// @router /permission/ [get]
func (r *Role) Permission() {
	r.apiResult.Code = constant.OK
	r.apiResult.Data = permission.Permission
	r.ServeJSON()
}

// @Title Update
// @Description update the Role
// @Success 200 {object} models.Role
// @Failure 403 :id is not int
// @router / [put]
func (c *Role) Put() {
	v := models.Role{}
	c.BindModelWithPut(&v, "role_name", "role_permission", "approve_year", "approved_year", "approve_month", "approved_month", "approve_order", "approve_level")
}

// @Title Put
// @Description lock the Role
// @Success 200 {object} models.Role
// @Failure 403 :id is not int
// @router /lock/:id [put]
func (c *Role) Lock() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Role{Id: uint(id), RoleLock: 0}
	c.BindModelWithPut(&v, "role_lock")
}

// @Title Put
// @Description unlock the Role
// @Success 200 {object} models.Role
// @Failure 403 :id is not int
// @router /unlock/:id [put]
func (c *Role) UnLock() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Role{Id: uint(id), RoleLock: 1}
	c.BindModelWithPut(&v, "role_lock")
}
