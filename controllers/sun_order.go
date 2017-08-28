package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
	"github.com/yakun0622/shop/tools"
	"github.com/astaxie/beego"
)

// oprations for SunOrder
type SunOrderController struct {
	BaseController
}

func (c *SunOrderController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetGoodes", c.GetGoodesAndApprovers)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("User", c.User)
	c.Mapping("GroupOrders", c.GroupOrders)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetTransportfee", c.GetTransportfee)
}

// @router /user [post]
func (c *SunOrderController) User() {
	orderState := c.GetString("OrderState")
	groupId, _ := c.GetInt("GroupId")
	memberId, _ := c.GetInt("MemberId")
	page, _ := c.GetInt("Page")
	orderType, _ := c.GetInt("OrderType")
	orders, num, err := models.GetGroupOrder(groupId, memberId, orderState, orderType, (page - 1) * 20)
	if err == nil {
		c.Display("user", orders)
		c.apiResult.Code = constant.OK
		c.apiResult.Data = orders
		c.apiResult.Count = int(num)
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}

// @router /goodses [post]
func (c *SunOrderController) GetGoodesAndApprovers() {
	orderId, _ := c.GetInt("OrderId")
	groupId, _ := c.GetInt("GroupId")

	goodses, tags, err := models.GetOrderGoodsesAndTags(orderId, groupId)
	if err != nil {
		c.apiResult.Code = constant.FAIL
		c.ServeJSON()
		return
	}

	approvers, currentApproverId, approverIds, err := models.GetOrderApprovers(orderId)
	if err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = map[string]interface{}{
			"goodses": goodses,
			"tags":    tags,
			"approvers": approvers,
			"approversIds": approverIds,
			"currentApproverId": currentApproverId,
		}
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}

// @router /get [get]
func (c *SunOrderController) GroupOrders() {
	state, _ := c.GetInt("OrderState")
	groupId, _ := c.GetInt("GroupId")
	offset, _ := c.GetInt("offset")

	if groupId != 0 {
		orders, goodses, count, err := models.GetOrderByGroupId(
			groupId,
			state,
			offset,
		)
		if err != nil {
			c.apiResult.Code = constant.FAIL
		} else {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = map[string]interface{}{
				"orders":  orders,
				"goodses": goodses,
			}
			c.apiResult.Count = count
		}
	}
	c.ServeJSON()
}

// @Title Post
// @Description create SunOrder
// @Param	body		body 	models.SunOrder	true		"body for SunOrder content"
// @Success 201 {int} models.SunOrder
// @Failure 403 body is empty
// @router / [post]
func (c *SunOrderController) Post() {
	orders := []models.SunOrder{}
	goods := [][]models.SunOrderGoods{}
	if err := json.Unmarshal([]byte(c.GetString("Orders")), &orders); err != nil {
		c.apiResult.Code = constant.FAIL
		c.ServeJSON()
		return
	}
	if err := json.Unmarshal([]byte(c.GetString("Goods")), &goods); err != nil {
		c.apiResult.Code = constant.FAIL
		c.ServeJSON()
		return
	}

	//商品标签
	var tagIds []string
	tagIdsStr := c.GetString("TagIds")
	tagIds = strings.Split(tagIdsStr, "|")

	//组父级Id
	groupParentId, _ := c.GetInt64("GroupParentId")

	// c.Display("orders", orders)
	// c.Display("goods", goods)

	//是否是购物车订单
	isCart, _ := c.GetBool("IsCart")
	//审批级别
	approveLevel, _ := c.GetInt("ApproveLevel")
	//是否是自创的组
	isOwnerGroup, _ := c.GetBool("IsOwnerGroup")

	//运费
	areaId, _ := c.GetInt("AreaId")

	if models.SaveOrderWithGoods(
		orders,
		goods,
		tagIds,
		c.GetUserID(),
		approveLevel,
		uint(groupParentId),
		isCart,
		isOwnerGroup,
		areaId,
	) {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}

// @Title Get
// @Description get SunOrder by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunOrder
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunOrderController) GetOne() {
	//此处id为pay_sn的值
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.Atoi(idStr)
	// v, err := models.GetSunOrderByPaySN(id, c.GetCurrentUser().Id)
	// if err != nil {
	// 	c.apiResult.Code = constant.FindNoData
	// 	c.apiResult.Error = err
	// } else {
	// 	result := make(map[string]interface{})
	// 	result["order"] = v

	// 	orderCommons := make([]models.SunOrderCommon, len(v))
	// 	var orderGoods []models.SunOrderGoods
	// 	for index, order := range v {
	// 		orderID := (order.(models.SunOrder)).Id
	// 		//获取订单common
	// 		if value1, err := models.GetSunOrderCommonById(orderID); err == nil {
	// 			orderCommons[index] = *value1
	// 		}
	// 		//获取goods信息
	// 		if value2, err := models.GetSunOrderGoodsByOrderID(orderID); err == nil {
	// 			for _, value3 := range value2 {
	// 				orderGoods = append(orderGoods, value3.(models.SunOrderGoods))
	// 			}
	// 		}
	// 	}

	// 	result["orderCommon"] = orderCommons
	// 	result["orderGoods"] = orderGoods

	// 	c.apiResult.Code = constant.OK
	// 	c.apiResult.Data = result
	// }
	// c.ServeJSON()
}

// @Title Get All
// @Description get SunOrder
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunOrder
// @Failure 403
// @router / [get]
func (c *SunOrderController) GetAll() {
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
	query["buyer_id"] = strconv.Itoa(int(c.GetCurrentUser().Id))

	l, count, err := models.GetAllSunOrder(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = l
		c.apiResult.Count = int(count)
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunOrder
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunOrder	true		"body for SunOrder content"
// @Success 200 {object} models.SunOrder
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunOrderController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunOrder{Id: uint64(id)}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunOrderById(&v); err == nil {
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
// @Description delete the SunOrder
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunOrderController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunOrder(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title edit_shippingFee
// @Description 修改订单状态
// @Failure 403
// @router /change_state [get]
func (c *SunOrderController) ChangeState() {
	user := c.GetCurrentUser()
	if user.Id == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}
	//订单状态
	state_type := c.GetString("state_type")
	order_id, _ := c.GetInt("order_id")

	//获取order信息
	orderInfo, err := models.GetOrderById(order_id)
	//if int(orderInfo.BuyerId) != user.Id {
	//	c.HandleResult(err, "无权操作非本用户下的订单", constant.FAIL, nil, 0)
	//	return
	//}
	switch state_type {
	case "order_cancel":
		err = models.CancleOrder(orderInfo)
	case "order_affirm":
		err = models.AffirmOrder(orderInfo)
	default:
		c.HandleResult(nil, "订单操作类型错误！！！", constant.FAIL, nil, 0)
		return
	}

	if err == nil {
		c.HandleResult(err, "订单状态修改成功", constant.OK, true, 0)
	} else {
		c.HandleResult(err, "订单状态修改失败", constant.FAIL, nil, 0)
	}
}

// @Title get_settlement
// @Description 查询结算单
// @Failure 403
// @router /get_settlement [get]
func (c *SunOrderController) GetSettlement(){
	result := false
	count := 0
	c.HandleResult(nil, "结算单查询成功", constant.OK, result, count)
}

// @router /transportfee [post]
func (c *SunOrderController) GetTransportfee(){
	var goodseses []string
	var goodsNumses [][]int
	var amounts []float64
	areaId, _ := c.GetInt("AreaId")
	var transportfeeAmount float64
	err := json.Unmarshal([]byte(c.GetString("Goodseses")), &goodseses)
	err = json.Unmarshal([]byte(c.GetString("GoodsNumses")), &goodsNumses)
	err = json.Unmarshal([]byte(c.GetString("Amounts")), &amounts)
	if !tools.HasError(err, "json解析错误") {
		c.apiResult.Code = constant.FAIL
		return
	} else {
		for i, amount := range amounts  {
			beego.Info("GetTransportfee---", goodsNumses)
			transportfee:= models.GetTransportfee(goodseses[i], goodsNumses[i], amount, areaId )

				transportfeeAmount += transportfee
		}
	}

	c.HandleResult(nil, "运费查询成功", constant.OK, transportfeeAmount, 0)
}