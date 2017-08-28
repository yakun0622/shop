package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"fmt"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunStore
type SunStoreController struct {
	BaseController
}

func (c *SunStoreController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetSimpleStatic", c.GetSimpleStatic)
	c.Mapping("GetGoodsCommonList", c.GetGoodsCommonList)
	c.Mapping("SetFreeFreight", c.SetFreeFreight)
}

// @Title Post
// @Description create SunStore
// @Param	body		body 	models.SunStore	true		"body for SunStore content"
// @Success 201 {int} models.SunStore
// @Failure 403 body is empty
// @router / [post]
func (c *SunStoreController) Post() {
	var v models.SunStore
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSunStore(&v); err == nil {
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
// @Description get SunStore by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunStore
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunStoreController) GetOne() {
	//根据当前登录用户获取信息
	v, err := models.GetSunStoreById(int(c.GetCurrentStoreID()))
	if err != nil {
		c.apiResult.Error = err
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunStore
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunStore
// @Failure 403
// @router / [get]
func (c *SunStoreController) GetAll() {
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

	l, err := models.GetAllSunStore(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunStore
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunStore	true		"body for SunStore content"
// @Success 200 {object} models.SunStore
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunStoreController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunStore{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunStoreById(&v); err == nil {
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
// @Description delete the SunStore
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunStoreController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunStore(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get simple static
// @Description GetSimpleStatic
// @Success 200 {object} models.SunStore
// @router /static/simple [get]
func (c *SunStoreController) GetSimpleStatic() {
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}
	result := make(map[string]interface{})

	//1.店铺及商品提示
	//出售中商品
	if count, err := models.GetGoodsCommonCountByStatus(storeID, 1); err == nil {
		result["online"] = count
	}
	//等待审核
	if count, err := models.GetGoodsCommonCountByStatus(storeID, 2); err == nil {
		result["waitverify"] = count
	}
	//审核失败
	if count, err := models.GetGoodsCommonCountByStatus(storeID, 3); err == nil {
		result["verifyfail"] = count
	}
	//仓库待上架商品 - 仓库中已审核
	if count, err := models.GetGoodsCommonCountByStatus(storeID, 4); err == nil {
		result["offline"] = count
	}
	//待回复咨询
	if count, err := models.GetGoodsCommonCountByStatus(storeID, 5); err == nil {
		result["consult"] = count
	}
	//违规下架
	if count, err := models.GetGoodsCommonCountByStatus(storeID, 6); err == nil {
		result["lockup"] = count
	}

	//2.交易提示
	//待付款
	if count, err := models.GetOrderCountByStatus(storeID, 1); err == nil {
		result["payment"] = count
	}
	//待发货
	if count, err := models.GetOrderCountByStatus(storeID, 2); err == nil {
		result["delivery"] = count
	}
	//售前退款
	if count, err := models.GetOrderCountByStatus(storeID, 3); err == nil {
		result["refund_lock"] = count
	}
	//售后退款
	if count, err := models.GetOrderCountByStatus(storeID, 4); err == nil {
		result["refund"] = count
	}
	//售前退货
	if count, err := models.GetOrderCountByStatus(storeID, 5); err == nil {
		result["return_lock"] = count
	}
	//售后退货
	if count, err := models.GetOrderCountByStatus(storeID, 6); err == nil {
		result["return"] = count
	}
	//待确认的结算账单
	if count, err := models.GetOrderCountByStatus(storeID, 7); err == nil {
		result["bill_confirm"] = count
	}

	c.apiResult.Code = constant.OK
	c.apiResult.Data = result
	c.ServeJSON()
}

// @Title Get goods_common in a store
// @Description GetGoodsCommonList
// @Success 200 {object} models.SunStore
// @router /goods_common/online [get]
func (c *SunStoreController) GetGoodsCommonList() {
	limit, _ := c.GetInt64("limit")
	offset, _ := c.GetInt64("offset")
	goodsName := c.GetString("goodsName")
	//var classID = 0
	//needSpec := false
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}

	list, totalCount, err := models.GetGoodsCommonByStatus(1, storeID, goodsName, limit, offset)

	if err == nil {
		c.HandleResult(err, "获取出售商品成功", constant.OK, list, int(totalCount))
	} else {
		c.HandleResult(err, "获取出售商品失败", constant.FAIL, nil, 0)
	}

}

// @Title Get goods_common in a store
// @Description GetGoodsCommonList
// @Success 200 {object} models.SunStore
// @router /goods_common/offline [get]
func (c *SunStoreController) GetGoodsCommonOfflineList() {
	limit, _ := c.GetInt64("limit")
	offset, _ := c.GetInt64("offset")
	goodsName := c.GetString("goodsName")
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}

	goodsType := c.GetString("type")
	fmt.Println("goodsType=====", goodsType)
	var result []models.SunGoodsCommon
	var totalCount int64
	var err error
	var statusType int

	switch goodsType {
	case "WaitVerify":
		statusType = 4
		break
	case "lock":
		statusType = 3
		break
	default:
		statusType = 2
		break
	}
	result, totalCount, err = models.GetGoodsCommonByStatus(statusType, storeID, goodsName, limit, offset)
	if err != nil {
		c.HandleResult(err, "获取商品列表失败", constant.FAIL, nil, 0)
	} else {
		c.HandleResult(err, "获取商品列表成功", constant.OK, result, int(totalCount))
	}
	c.ServeJSON()

}

// @Title set free_freight
// @router /free_freight [post]
func (c *SunStoreController) SetFreeFreight() {
	var v models.SunStore
	fmt.Println(c.GetDataString())
	if err := json.Unmarshal([]byte(c.GetString("data")), &v); err == nil {
		if v, err = models.SetStoreFreeFreightByStroreID(c.GetCurrentStoreID(), v.StoreFreePrice); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v
		} else {
			c.apiResult.Error = err
			c.apiResult.Code = constant.OK
		}
	} else {
		c.apiResult.Error = err
		c.apiResult.Code = constant.OK
	}
	c.ServeJSON()
}
