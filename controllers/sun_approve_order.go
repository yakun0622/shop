package controllers

import (
	"strconv"
	"strings"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunApproveOrder
type SunApproveOrderController struct {
	BaseController
}

func (c *SunApproveOrderController) URLMapping() {
	c.Mapping("Approve", c.Approve)
	// c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

// @router /approve [post]
func (c *SunApproveOrderController) Approve() {
	//审批订单ids
	idsStr := c.GetString("ApproveOrderIds")
	//审批状态，是否通过
	state, _ := c.GetInt("ApproveState")
	//审批备注
	reason := c.GetString("ApproveOrderReason")
	//组Id
	//groupId, _ := c.GetInt("GroupId")
	////组父级Id
	//groupParentId, _ := c.GetInt("GroupParentId")
	////当前角色审批级别
	//approveLevel, _ := c.GetInt("ApproveLevel")
	//
	////订单类型
	//orderTypes := strings.Split(c.GetString("OrderTypes"), ",")
	//标签
	tagIds := strings.Split(c.GetString("TagIds"), "|")
	//订单
	orderIds := strings.Split(c.GetString("OrderIds"), ",")

	if err := models.ApproveOrders(
		idsStr,
		state,
		orderIds,
		reason,
		tagIds,
	); err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunApproveOrder
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunApproveOrder
// @Failure 403
// @router / [get]
func (c *SunApproveOrderController) GetAll() {
	page, _ := c.GetInt("Page")

	num, l, err := models.GetAllApproveOrders(c.GetString("RoleId"),
		c.GetString("GroupId"),
		c.GetString("ApproveState"),
		c.GetString("OrderType"),
		c.GetString("OrderSn"),
		page-1)
	if err == nil {
		if num > 0 {
			groupId, _ := c.GetInt("GroupBelong")
			if groupId == 0 {
				groupId, _ = c.GetInt("GroupId")
			}

			var orderIds []string
			for _, approveOrder := range l {
				orderId := strconv.Itoa(int(approveOrder.OrderId))
				orderIds = append(orderIds, orderId)
			}
			goodses, tags, err := models.GetGoodsAndTagsByApproveOrder(strings.Join(orderIds, ","), uint(groupId))
			if err == nil {
				c.apiResult.Code = constant.OK
				c.apiResult.Data = map[string]interface{}{
					"approveOrders": l,
					"goodses":       goodses,
					"tags":          tags,
				}
				c.apiResult.Count = num
			} else {
				c.apiResult.Code = constant.FAIL
			}
		} else {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = nil
		}
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}
