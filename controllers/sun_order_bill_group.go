package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
	"github.com/astaxie/beego"
)

// oprations for SunOrderBillGroup
type SunOrderBillGroupController struct {
	BaseController
}

func (c *SunOrderBillGroupController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetList", c.GetList)
}

// @Title Post
// @Description create SunOrderBillGroup
// @Param	body		body 	models.SunOrderBillGroup	true		"body for SunOrderBillGroup content"
// @Success 201 {int} models.SunOrderBillGroup
// @Failure 403 body is empty
// @router / [post]
func (c *SunOrderBillGroupController) Post() {
	var v models.SunOrderBillGroup
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSunOrderBillGroup(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get SunOrderBillGroup by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunOrderBillGroup
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunOrderBillGroupController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunOrderBillGroupById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunOrderBillGroup
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunOrderBillGroup
// @Failure 403
// @router / [get]
func (c *SunOrderBillGroupController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSunOrderBillGroup(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunOrderBillGroup
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunOrderBillGroup	true		"body for SunOrderBillGroup content"
// @Success 200 {object} models.SunOrderBillGroup
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunOrderBillGroupController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunOrderBillGroup{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunOrderBillGroupById(&v); err == nil {
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
// @Description delete the SunOrderBillGroup
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunOrderBillGroupController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunOrderBillGroup(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get All List
// @Description get SunOrderBillGroup
// @Success 200 {object} models.SunOrderBillGroup
// @Failure 403
// @router /list [get]
func (c *SunOrderBillGroupController) GetList() {
	var groupIds []int
	GroupId := c.GetString("GroupId")
	PeriodType := c.GetString("PeriodType")
	groupId, _ := strconv.Atoi(GroupId)
	periodType, _ := strconv.Atoi(PeriodType)
	childGroups, _ := models.GetGroupChildList(groupId)
	groupIds = append(groupIds, groupId)
	if len(childGroups) > 0 {
		for _, item := range childGroups {
			group := item.(models.SunGroup)
			groupIds = append(groupIds, group.Id)
		}
	}
	beego.Info("groupIds>>>>",groupIds)
	if list, count, err := models.GetSunOrderBillGroupByGroupId(groupIds, periodType); err == nil {
		c.HandleResult(nil, "查询清单成功", constant.OK, list, int(count))
	} else {
		c.HandleResult(err, "查询清单失败", constant.FAIL, err, 0)
	}

}
