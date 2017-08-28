package models

import (
	//"strconv"
	"strings"

	"github.com/yakun0622/shop/tools"
)

type SunApproveOrder struct {
	Id                 int     `orm:"column(approve_order_id);auto"`
	RoleId             uint    `orm:"column(role_id);auto"`
	GroupId            uint    `orm:"column(group_id);auto"`
	StoreId            uint    `orm:"column(store_id)"`
	StoreName          string  `orm:"column(store_name);size(50)"`
	OrderType          int8    `orm:"column(order_type)"`
	OrderId            uint    `orm:"column(order_id)"`
	OrderSn            string  `orm:"column(order_sn)"`
	OrderAmount        float64 `orm:"column(order_amount);digits(10);decimals(2)"`
	ApproveOrderState  int8    `orm:"column(approve_order_state)"`
	ApproveOrderTime   int     `orm:"column(approve_order_time);null"`
	ApproveOrderReason string  `orm:"column(approve_order_reason);size(500);null"`
	Ctime              int     `orm:"column(ctime)"`
	BuyerId            uint    `orm:"column(buyer_id)"`
	TagIds             string  `orm:"column(tag_ids)"`
	BuyerName          string  `orm:"column(buyer_name);size(50)"`
}

func (t *SunApproveOrder) TableName() string {
	return "sun_approve_order"
}

func GetAllApproveOrders(roleId string, groupId string, approveState string, orderType string, orderSn string, offset int) (int, []SunApproveOrder, error) {
	o, q := GetQueryBuilder()

	q = q.Select("count(*)").From("sun_approve_order as ao").
		InnerJoin("sun_order as o").On("ao.order_id = o.order_id").
		Where("o.order_state!=0")

	if roleId != "" {
		q = q.And("ao.role_id = " + roleId)
	} else {
		q = q.And("ao.group_id = " + groupId)
	}

	if orderType != "" {
		q = q.And("o.order_type =" + orderType)
	}

	if orderSn != "" {
		q = q.And("o.order_sn in (" + orderSn + ")")
	}

	if approveState != "" {
		q = q.And("ao.approve_order_state=" + approveState)
	}

	var num int
	// Display("order", q.String())
	o.Raw(q.String()).QueryRow(&num)

	q = q.OrderBy("ao.ctime").Desc().Limit(20).Offset(offset * 20)

	var approveOrders []SunApproveOrder

	_, err := o.Raw(strings.Replace(q.String(), "count(*)", "*", 1)).QueryRows(&approveOrders)
	if err != nil {
		// Display("err", q.String())
		return 0, nil, err
	}
	return num, approveOrders, nil
}

func ApproveOrders(ids string, state int, orderIds []string, reason string, tagIds []string) error {
	o := Orm()
	o.Begin()
	Display("approve-state", tagIds)
	_, err := o.Raw("UPDATE sun_approve_order SET approve_order_state=?, approve_order_reason=?, approve_order_time=? WHERE approve_order_id in (" + ids + ")", state, reason, tools.GetTime()).Exec()
	if err != nil {
		Display("approve-state11", err)
		o.Rollback()
		return err
	}

	orderUpdate, _ := o.Raw("UPDATE sun_order SET order_state=?, approve_time=? WHERE order_id=?").Prepare()
	approveInsert, _ := o.Raw("INSERT INTO sun_approve_order (order_id, role_id, group_id, tag_ids, ctime) VALUES (?, ?, ?, ?, ?)").Prepare()

	Display("approve-state22", state)
	for i, orderId := range orderIds {
		if state == 1 {
			roleId, groupId, approvers, err := GetApproveRoleIdAndGroupId(orderId)
			if err != nil {
				Display("approve-GetApproveRoleIdAndGroupId", err)
				o.Rollback()
				return err
			}

			if roleId == "" || groupId == "" {
				_, err := orderUpdate.Exec(30, tools.GetTime(), orderId)
				if err != nil {
					Display("approve-orderUpdate", err)
					o.Rollback()
					return err
				}
			} else {
				_, err := approveInsert.Exec(orderId, roleId, groupId, tagIds[i], tools.GetTime())
				if err != nil {
					Display("approve-approveInsert", err)
					o.Rollback()
					return err
				}

				_, err = o.Raw("UPDATE sun_order SET approvers=? WHERE order_id=?", approvers, orderId).Exec()
				if err != nil {
					Display("approve-approversUpdate", err)
					o.Rollback()
					return err
				}
			}
		} else {
			_, err := orderUpdate.Exec(20, tools.GetTime(), orderId)
			if err != nil {

				Display("orderUpdate", err)
				o.Rollback()
				return err
			}
		}
	}
	Display("approve-end", state)
	approveInsert.Close()
	orderUpdate.Close()
	o.Commit()
	return nil
}

type approveGoodsAndTags struct {
	OrderId    uint    `orm:"column(order_id)"`
	GoodsId    uint    `orm:"column(goods_id)"`
	GoodsName  string  `orm:"column(goods_name);size(50)"`
	GoodsPrice float64 `orm:"column(goods_price);digits(10);decimals(2)"`
	StoreId    uint    `orm:"column(store_id)"`
	StoreName  string  `orm:"column(store_name);size(50)"`
	GoodsNum   uint16  `orm:"column(goods_num)"`
	GoodsImage string  `orm:"column(goods_image);size(100);null"`
	TagId      uint64  `orm:"column(tag_id)"`
	TagName    string  `orm:"column(tag_name)"`
	GroupId    uint64  `orm:"column(group_id)"`
}

func GetGoodsAndTagsByApproveOrder(orderIds string, groupId uint) (goodses []SunOrderGoods, tags [][]Tag, err error) {
	o, q := GetQueryBuilder()
	sql := q.Select("*").From("sun_order_goods").Where("order_id in(" + orderIds + ")").String()

	_, err = o.Raw(sql).QueryRows(&goodses)

	tagQuery := QueryBuilder()
	sql = tagQuery.Select("*").
		From("sun_goods_tag as tg").
		InnerJoin("sun_tag as t").On("tg.tag_id = t.tag_id").
		Where("t.group_id = ?").
		And("tg.goods_id = ?").
		String()
	for _, goods := range goodses {
		var tag []Tag
		_, err = o.Raw(sql, groupId, goods.GoodsId).QueryRows(&tag)
		if err == nil {
			tags = append(tags, tag)
		} else {
			return
			Display("GetGoodsAndTagsByApproveOrder", "获取tag失败")
		}
	}
	return
}
