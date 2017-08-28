package controllers

import (
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// operations for SunExpressController
type SunExpressController struct {
	BaseController
}

func (c *SunExpressController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Create
// @Description create SunExpress
// @Param	body		body 	models.SunExpress	true		"body for SunExpress content"
// @Success 201 {object} models.SunExpress
// @Failure 403 body is empty
// @router / [post]
func (c *SunExpressController) Post() {

}

// @Title GetOne
// @Description get SunExpressController by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunExpressController
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunExpressController) GetOne() {

}

// @Title GetAll
// @Description get SunExpress
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunExpress
// @Failure 403
// @router / [get]
func (c *SunExpressController) GetAll() {
	result, err := models.GetAllSunExpress()
	if err != nil {
		c.HandleResult(err, "获取快递数据错误", constant.FAIL, nil, 0)
	} else {
		c.HandleResult(err, "成功获取快递数据信息", constant.OK, result, 0)
	}
}

// @Title Update
// @Description update the SunExpress
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunExpress	true		"body for SunExpress content"
// @Success 200 {object} models.SunExpress
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunExpressController) Put() {

}

// @Title Delete
// @Description delete the SunExpress
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunExpressController) Delete() {

}
