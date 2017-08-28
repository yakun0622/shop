package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for Transport
type TransportController struct {
	BaseController
}

func (c *TransportController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetList)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Transport
// @Param	body		body 	models.Transport	true		"body for Transport content"
// @Success 201 {int} models.Transport
// @Failure 403 body is empty
// @router / [post]
func (c *TransportController) Post() {
	var v models.Transport
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTransport(&v); err == nil {
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

// @Title 新增运费模板
// @Failure 403 body is empty
// @router /add [post]
func (c *TransportController) AddTransport() {
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, false, 0)
		return
	}

	var transport models.Transport
	transport.StoreId = uint32(storeID)
	transport.UpdateTime = c.GetCurrentTime()
	err := json.Unmarshal(c.GetDataBytes(), &transport)

	if err != nil {
		c.HandleResult(nil, "数据解码失败", constant.FAIL, false, 0)
		return
	}
	//插入主表
	transportId, err := models.AddTransport(&transport)
	if err != nil {
		c.HandleResult(nil, "插入运费模板主表失败", constant.FAIL, false, 0)
		return
	}

	//插入字表
	transportExtend := transport.Extends[0]
	transportExtend.TransportId = transportId
	_, err = models.AddTransportExtend(&transportExtend)
	if err != nil {
		c.HandleResult(nil, "插入运费模板主表失败", constant.FAIL, false, 0)
		return
	}
	if err != nil {
		c.HandleResult(err, "新增运费策略失败", constant.FAIL, false, 0)
	} else {
		c.HandleResult(err, "新增运费策略成功", constant.OK, true, 0)
	}
}

// @Title Get
// @Description get Transport by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Transport
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TransportController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunTransportById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Transport
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Transport
// @Failure 403
// @router /list [get]
func (c *TransportController) GetList() {
	storeID := c.GetCurrentStoreID()
	fmt.Println("当前商户ID：", storeID)
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}
	transport_list, count, err := models.GetTransportListByStoreId(storeID)

	beego.Info(transport_list)
	for i := 0; i < len(transport_list); i++ {
		if transport_extend_list, err := models.GetByTransportExtTransportId(int64(transport_list[i].Id)); err == nil {
			transport_list[i].Extends = transport_extend_list
		} else {
			c.HandleResult(err, "获取运费模板失败", constant.FAIL, nil, 0)
			return
		}
	}

	if err == nil {
		c.HandleResult(err, "获取运费模板成功", constant.OK, transport_list, int(count))
	} else {
		c.HandleResult(err, "获取运费模板失败", constant.FAIL, nil, 0)
	}
}

// @Title Update
// @Description update the Transport
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Transport	true		"body for Transport content"
// @Success 200 {object} models.Transport
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TransportController) Put() {
	idStr := c.GetString(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Transport{Id: id}
	if err := json.Unmarshal(c.GetDataBytes(), &v); err == nil {
		v.UpdateTime = c.GetCurrentTime()
		if err := models.UpdateTransportById(&v); err == nil {
			c.HandleResult(err, "更新运费模板成功", constant.OK, true, 0)
		} else {
			c.HandleResult(err, "更新运费模板失败", constant.FAIL, false, 0)
		}
	} else {
		c.HandleResult(err, "更新运费模板失败", constant.FAIL, false, 0)
	}
}

// @Title Delete
// @Description delete the Transport
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TransportController) Delete() {
	id,_ := c.GetInt64(":id")
	//删除字表
	err := models.DeleteTransportExtendByTransId(id)
	if err != nil {
		c.HandleResult(err, "删除运费模板子表失败", constant.FAIL, false, 0)
	}
	err = models.DeleteSunTransport(int(id))
	if err != nil {
		c.HandleResult(err, "删除运费模板主表失败", constant.FAIL, false, 0)
	}else {
		c.HandleResult(err, "删除运费模板成功", constant.OK, true, 0)
	}
}
