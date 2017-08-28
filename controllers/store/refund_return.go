package store

import (
	"encoding/json"
	"strconv"

	"github.com/yakun0622/shop/models"

	"github.com/astaxie/beego"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/controllers"
	"github.com/yakun0622/shop/tools"
)

// oprations for SunRefundReturn
type RefundReturnController struct {
	controllers.BaseController
}

func (c *RefundReturnController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetList", c.GetList)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create SunRefundReturn
// @Param	body		body 	models.SunRefundReturn	true		"body for SunRefundReturn content"
// @Success 201 {int} models.SunRefundReturn
// @Failure 403 body is empty
// @router / [post]
func (c *RefundReturnController) Post() {
	var m models.SunRefundReturn
	var goodInfo []map[string]int
	var addFlag int
	var totalRefundAmount float64

	addFlag = 0
	totalRefundAmount = 0
	OrderId,_ := c.GetInt("OrderId")
	beego.Info("get_orderid:",OrderId)
	GoodsInfo := c.GetString("GoodsInfo")
	json.Unmarshal([]byte(GoodsInfo), &goodInfo)
	OrderInfo, _ := models.GetOrderById(OrderId)
	m.OrderId = uint(OrderId)
	m.OrderSn = OrderInfo.OrderSn
	m.AddTime = tools.GetTime()
	m.SellerTime = OrderInfo.AddTime
	m.OrderGoodsType = uint8(OrderInfo.OrderType)
	m.StoreId = OrderInfo.StoreId
	m.StoreName = OrderInfo.StoreName
	m.BuyerId = OrderInfo.BuyerId
	m.BuyerName = OrderInfo.BuyerName
	m.RefundSn = tools.GetTimeString(2) + strconv.Itoa(OrderId)
	m.GroupId = int(OrderInfo.GroupId)
	m.GroupName = OrderInfo.GroupName
	beego.Info("show ob:",goodInfo)
	for _, v := range goodInfo{
		for key, value := range v {
			id, _ := strconv.Atoi(key)
			num := value
			goodInfo, _ := models.GetSunGoodsById(id)
			orderGoodsInfo, _ := models.GetSingeSunOrderGoodsByOrderID(OrderId, id)
			m.Goods = goodInfo
			m.GoodsName = goodInfo.GoodsName
			m.GoodsNum = uint(num)
			m.GoodsImage = goodInfo.GoodsImage
			m.RefundAmount =  float64(float64(num) * orderGoodsInfo.GoodsPrice)
			totalRefundAmount += m.RefundAmount
			beego.Info("get id:", orderGoodsInfo)
			if _, err := models.AddSunRefundReturn(&m); err == nil {
				addFlag = 1
			}

		}
	}
	beego.Info("totalRefundAmount:",totalRefundAmount)
	if addFlag == 1{
		OrderInfo.Id = uint64(OrderId)
		OrderInfo.RefundAmount = totalRefundAmount
		if totalRefundAmount == OrderInfo.OrderAmount{
			OrderInfo.RefundState = 2
		} else {
			OrderInfo.RefundState = 1
		}
		models.UpdateSunOrderById(OrderInfo)
		c.HandleResult(nil, "添加成功", constant.OK, m, 1)
	}else{
		c.HandleResult(nil, "添加失败", constant.FAIL, nil, 0)
	}
}

// @Title Get
// @Description get SunRefundReturn by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunRefundReturn
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RefundReturnController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunRefundReturnById(id)
	if err != nil {
		c.HandleResult(err, "未找到商户订单", constant.FAIL, nil, 0)
	} else {
		c.HandleResult(err, "成功获取商户订单", constant.OK, v, 1)
	}
	//c.ServeJSON()
}

// @Title Get All
// @Description get SunRefundReturn
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunRefundReturn
// @Failure 403
// @router /list [get]
func (c *RefundReturnController) GetList() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64
	var offset int64

	query, fields, sortby, order, offset, limit, err := c.GetAllSeachParams()
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}
	query["StoreId"] = strconv.Itoa(int(storeID))
	list, count, err := models.GetAllSunRefundReturn(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.HandleResult(err, "获取商户订单列表错误", constant.FAIL, nil, 0)
	} else {
		c.HandleResult(err, "成功获取商户订单列表", constant.OK, list, int(count))
	}
}

// @Title Update
// @Description update the SunRefundReturn
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunRefundReturn	true		"body for SunRefundReturn content"
// @Success 200 {object} models.SunRefundReturn
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RefundReturnController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, _ := models.GetSunRefundReturnById(id)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunRefundReturnById(v); err == nil {
			if v.SellerState == 2 {
				vi, _ := models.GetSunRefundReturnById(id)
				models.UpdateGoodsStorageByID(vi.Goods.Id, vi.GoodsNum)
			}
			c.HandleResult(nil, "修改订单成功", constant.OK, v, 1)
		} else {
			c.HandleResult(err, "修改订单失败", constant.FAIL, nil, 0)
		}
	}
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	//	if err := models.UpdateSunRefundReturnById(&v); err == nil {
	//		c.Data["json"] = "OK"
	//	} else {
	//		c.Data["json"] = err.Error()
	//	}
	//} else {
	//	c.Data["json"] = err.Error()
	//}
	//c.ServeJSON()
}

// @Title Delete
// @Description delete the SunRefundReturn
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RefundReturnController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunRefundReturn(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
