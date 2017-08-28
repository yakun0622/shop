package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yakun0622/shop/constant"

	"github.com/yakun0622/shop/models"
)

// oprations for SunGoods
type SunGoodsController struct {
	BaseController
}

func (c *SunGoodsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Details", c.Details)

}

type goodsDatas struct {
	GoodsCommon models.SunGoodsCommon
	Goodses     []models.SunGoods
}

// @router /details [get]
func (c *SunGoodsController) Details() {
	goodsId := c.GetString("GoodsId")
	goodsCommonId := c.GetString("GoodsCommonId")
	if goodsId != "" || goodsCommonId != "" {
		goodsCommon, goodses, err := models.GetGoodsCommonAndGoodses(goodsId, goodsCommonId)
		if err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = goodsDatas{
				GoodsCommon: goodsCommon,
				Goodses:     goodses,
			}
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Post
// @Description create SunGoods
// @Param	body		body 	models.SunGoods	true		"body for SunGoods content"
// @Success 201 {int} models.SunGoods
// @Failure 403 body is empty
// @router / [post]
func (c *SunGoodsController) Post() {
	var v models.SunGoods
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSunGoods(&v); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v
		} else {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.InvalidRequestData
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Get
// @Description get SunGoods by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunGoods
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunGoodsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunGoodsById(id)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunGoods
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunGoods
// @Failure 403
// @router / [get]
func (c *SunGoodsController) GetAll() {
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
				c.apiResult.Code = constant.InvalidKeyValue
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSunGoods(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunGoods
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunGoods	true		"body for SunGoods content"
// @Success 200 {object} models.SunGoods
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunGoodsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunGoods{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunGoodsById(&v); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v
		} else {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.InvalidRequestData
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the SunGoods
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunGoodsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunGoods(id); err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.ActionFaild
		c.apiResult.Error = err
	}
	c.ServeJSON()
}
