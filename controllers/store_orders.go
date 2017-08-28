package controllers

import (
	"strconv"

	"encoding/json"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

//SearchControllerController search
type StoreOrdersController struct {
	BaseController
}

func (c *StoreOrdersController) URLMapping() {
	c.Mapping("StoreOrders", c.StoreOrders)
	c.Mapping("GetOrderInfo", c.GetOrderInfo)
	c.Mapping("EditShippingFee", c.EditShippingFee)
	c.Mapping("ChangeState", c.ChangeState)
	c.Mapping("GetOrderGoodes", c.GetOrderGoodes)
}

// @router /list [get]
func (c *StoreOrdersController) StoreOrders() {
	//state, _ := c.GetInt("OrderState")
	//groupId, _ := c.GetInt("GroupId")
	offset, _ := c.GetInt("offset")
	limit, _ := c.GetInt("limit")

	var ORDER_STATES = map[string]int{
		"state_new":          constant.ORDER_STATE_NEW,
		"state_approve_pass": constant.ORDER_STATE_APPROVE_PASS,
		"state_send":         constant.ORDER_STATE_SEND,
		"state_success":      constant.ORDER_STATE_SUCCESS,
		"state_cancel":       constant.ORDER_STATE_CANCEL,
	}
	orderState := c.GetString("state_type")
	state := -1
	if len(orderState) > 0 {
		state = ORDER_STATES[orderState]
	}
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}

	orders, count, err := models.GetOrderByStoreId(
		int(storeID),
		state,
		limit,
		offset,
	)
	if err != nil {
		c.HandleResult(err, "获取商户订单数据失败", constant.FAIL, nil, 0)
		c.apiResult.Code = constant.FAIL
	} else {
		c.HandleResult(nil, "获取商户订单数据成功", constant.OK, orders, count)
	}
}

// @Title Get Order Info
// @Description 获取商户下的订单列表
// @Failure 403
// @router /orderInfo [post]
func (c *StoreOrdersController) GetOrderInfo() {
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}

	orderId, err := strconv.Atoi(c.GetString("orderId"))
	if orderId <= 0 {
		c.HandleResult(nil, "订单编号不合法", constant.FAIL, nil, 0)
	}

	orderInfo, err := models.GetOrderById(orderId)
	if err != nil {
		c.HandleResult(nil, "获取商户"+strconv.Itoa(int(storeID))+"订单信息"+strconv.Itoa(int(orderId))+"错误", constant.FAIL, nil, 0)
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = orderInfo
	}
	c.ServeJSON()
}

// @Title edit_shippingFee
// @Description 修改订单运费
// @Failure 403
// @router /edit_shippingFee [post]
func (c *StoreOrdersController) EditShippingFee() {
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}

	var shippingFeeData map[string]interface{}
	err := json.Unmarshal(c.GetDataBytes(), &shippingFeeData)
	var result = false
	if err == nil {
		result = models.EditShippingFee(shippingFeeData)
	}

	if result {
		c.HandleResult(err, "发货提交成功", constant.OK, result, 0)
	} else {
		c.HandleResult(err, "发货提交失败", constant.FAIL, nil, 0)
	}
}

// @Title edit_shippingFee
// @Description 修改订单运费
// @Failure 403
// @router /change_state/:id [post]
func (c *StoreOrdersController) ChangeState() {
	result := false
	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}
	//订单状态
	state_type := c.GetString("state_type")
	order_id, _ := c.GetInt("order_id")

	//获取order信息
	orderInfo, err := models.GetOrderById(order_id)
	if orderInfo.StoreId != storeID {
		c.HandleResult(err, "无权取消非本商户下的订单", constant.FAIL, result, 0)
		return
	}

	if state_type == "order_cancel" {
		//result = orderCancel(order_id, )
	}

	if result {
		c.HandleResult(err, "发货提交成功", constant.OK, result, 0)
	} else {
		c.HandleResult(err, "发货提交失败", constant.FAIL, nil, 0)
	}
}

// @router /goodses [post]
func (c *StoreOrdersController) GetOrderGoodes() {
	orderId, _ := c.GetInt("OrderId")

	goodses, err := models.GetSunOrderGoodsByOrderID(orderId)
	if err != nil {
		c.apiResult.Code = constant.FAIL
		c.ServeJSON()
		return
	}

	if err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = map[string]interface{}{
			"goodses": goodses,
		}
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}
