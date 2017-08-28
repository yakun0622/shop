package controllers

import (
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
	"strconv"
)

// oprations for SunOffpayArea
type SunOffpayAreaController struct {
	BaseController
}

func (c *SunOffpayAreaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Save", c.Save)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Get
// @Description get SunAdmin by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunAdmin
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunOffpayAreaController) GetOne() {
	//idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.Atoi(idStr)
	//v, err := models.GetSunAdminById(id)
	//if err != nil {
	//	c.Data["json"] = err.Error()
	//} else {
	//	c.Data["json"] = v
	//}
	//c.ServeJSON()
}

// @Title Get All
// @Description get SunOffpayArea
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunOffpayArea
// @Failure 403
// @router / [get]
func (c *SunOffpayAreaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	query, fields, sortby, order, offset, limit, err := c.GetAllSeachParams()

	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}
	query["Id"] = strconv.Itoa(int(storeID))

	var result interface{}
	l, err := models.GetAllSunOffpayArea(query, fields, sortby, order, offset, limit)
	if len(l) > 0 {
		result = l[0]
	}
	if err != nil {
		c.HandleResult(err, "获取商户货到付款地区列表错误", constant.FAIL, nil, 0)
	} else {
		c.HandleResult(err, "获取商户货到付款地区列表成功", constant.OK, result, 0)
	}
}

// @Title 保存
// @Description 保存商户货到付款区域
// @router /save [post]
func (c *SunOffpayAreaController) Save() {
	var err error
	areaId := c.GetString("areaId")
	if len(areaId) <= 0 {
		c.HandleResult(nil, "areaId不能为空", constant.FAIL, false, 0)
		return
	}

	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, false, 0)
		return
	}
	//查找商户下的数据
	offpayArea, err := models.GetSunOffpayAreaById(int(storeID))
	if offpayArea == nil && err == nil {
		//新增数据
		offpayArea.Id = int(storeID)
		offpayArea.AreaId = areaId
		_, err = models.AddSunOffpayArea(offpayArea)
	} else {
		//修改数据
		offpayArea.AreaId = areaId
		err = models.UpdateSunOffpayAreaById(offpayArea)
	}
	if err != nil {
		c.HandleResult(nil, "保存货到付款地区失败", constant.FAIL, false, 0)
	} else {
		c.HandleResult(nil, "保存货到付款地区成功", constant.OK, true, 0)
	}
}
