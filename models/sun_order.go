package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"

	"time"

	"github.com/astaxie/beego"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/tools"
)

type SunOrder struct {
	Id              uint64  `orm:"column(order_id);auto"`
	OrderSn         string  `orm:"column(order_sn)"`
	PaySn           uint64  `orm:"column(pay_sn)"`
	StoreId         uint    `orm:"column(store_id)"`
	StoreName       string  `orm:"column(store_name);size(50)"`
	BuyerId         uint    `orm:"column(buyer_id)"`
	BuyerName       string  `orm:"column(buyer_name);size(50)"`
	BuyerEmail      string  `orm:"column(buyer_email);size(80)"`
	BuyerPhone      string  `orm:"column(buyer_phone);size(16)"`
	AddTime         uint    `orm:"column(add_time)"`
	PaymentCode     string  `orm:"column(payment_code);size(10)"`
	PaymentTime     uint    `orm:"column(payment_time);null"`
	FinnshedTime    uint    `orm:"column(finnshed_time)"`
	GoodsAmount     float64 `orm:"column(goods_amount);digits(10);decimals(2)"`
	OrderAmount     float64 `orm:"column(order_amount);digits(10);decimals(2)"`
	RcbAmount       float64 `orm:"column(rcb_amount);digits(10);decimals(2)"`
	PdAmount        float64 `orm:"column(pd_amount);digits(10);decimals(2)"`
	ShippingFee     float64 `orm:"column(shipping_fee);null;digits(10);decimals(2)"`
	EvaluationState int8    `orm:"column(evaluation_state);null"`
	OrderState      int8    `orm:"column(order_state)"`
	RefundState     uint8   `orm:"column(refund_state);null"`
	LockState       uint8   `orm:"column(lock_state);null"`
	DeleteState     int8    `orm:"column(delete_state)"`
	RefundAmount    float64 `orm:"column(refund_amount);null;digits(10);decimals(2)"`
	DelayTime       uint    `orm:"column(delay_time);null"`
	OrderFrom       int8    `orm:"column(order_from)"`
	ShippingCode    string  `orm:"column(shipping_code);size(50);null"`
	GroupId         uint    `orm:"column(group_id);null"`
	GroupName       string  `orm:"column(group_name);size(64);null"`
	ApproverId      int     `orm:"column(approver_id);null"`
	ApproverName    string  `orm:"column(approver_name);size(50);null"`
	Approvers       string  `orm:"column(approvers);size(255);null"`
	ApproveReason   string  `orm:"column(approve_reason);size(400);null"`
	ApproveTime     int     `orm:"column(approve_time);null"`
	OrderType       int8    `orm:"column(order_type)"`
	Personal        int8    `orm:"column(personal)"`
	AddressId       uint64  `orm:"column(address_id)"`
	BuyReason       string  `orm:"column(buy_reason);size(2000);null"`

	ShippingTime      uint   `orm:"column(shipping_time)"`
	ShippingExpressId int8   `orm:"column(shipping_express_id)"`
	EvalsellerState   int8   `orm:"column(evalseller_state)"`
	OrderMessage      string `orm:"column(order_message);size(300);null"`
	DeliverExplain    string `orm:"column(deliver_explain);null"`
	DaddressId        int32  `orm:"column(daddress_id)"`
	ReciverName       string `orm:"column(reciver_name);size(50)"`
	ReciverInfo       string `orm:"column(reciver_info);size(500)"`
	ReciverProvinceId uint32 `orm:"column(reciver_province_id)"`
	ReciverCityId     uint32 `orm:"column(reciver_city_id)"`
	InvoiceInfo       string `orm:"column(invoice_info);size(500);null"`
	InvoiceCode       string `orm:"column(invoice_code);size(50);null"`
	InvoiceNumber     string `orm:"column(invoice_number);size(20);null"`
	InvoiceImg        string `orm:"column(invoice_img);size(500);null"`
	DlyoPickupCode    string `orm:"column(dlyo_pickup_code);size(4);null"`
}

func (t *SunOrder) TableName() string {
	return "shop_order"
}

func init() {
	orm.RegisterModel(new(SunOrder))
}

// AddSunOrder insert a new SunOrder into database and returns
// last inserted Id on success.
func AddSunOrder(m *SunOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunOrderById retrieves SunOrder by Id. Returns error if
// Id doesn't exist
func GetSunOrderById(ID int) (v *SunOrder, err error) {
	o := orm.NewOrm()
	v = new(SunOrder)
	if err = o.QueryTable(v).Filter("ID", ID).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetOrderByGroupId(groupId int, state int, offset int) (orders []SunOrder, goodses [][]SunOrderGoods, count int, err error) {
	o, q := GetQueryBuilder()

	q = q.Select("count(*)").From("shop_order").Where("order_state=?").And("group_id = ?")

	err = o.Raw(q.String(), state, groupId).QueryRow(&count)
	if err != nil {
		Display("count", err)
		return
	}

	switch state {
	case 10:
		q = q.OrderBy("add_time")
	case 20, 30, 40:
		q = q.OrderBy("approve_time")
	case 0, 50:
		q = q.OrderBy("finnshed_time")
	}

	sql := strings.Replace(q.Desc().Limit(20).Offset(offset).String(), "count(*)", "*", 1)

	_, err = o.Raw(sql, state, groupId).QueryRows(&orders)

	if err != nil {
		Display("orders", err)
		return
	}

	var goods []SunOrderGoods
	for _, order := range orders {
		goodsQuery := QueryBuilder()
		_, err = o.Raw(goodsQuery.Select("*").From("shop_order_goods").Where("order_id=?").String(), order.Id).QueryRows(&goods)
		if err != nil {
			Display("goodses", err)
			return
		}
		goodses = append(goodses, goods)
	}

	return
}

func GetGroupOrder(groupId int, memberId int, orderState string, orderType int, offset int) (orders []SunOrder, num int64, err error) {
	o, q := GetQueryBuilder()

	q.Select("*").From("shop_order").Where("group_id=?").And("order_state=?")

	if orderType > -1 {
		q.And("order_type=?")
	}

	if memberId != 0 {
		q.And("buyer_id=?")
	}

	switch orderState {
	case "30":
		q.OrderBy("approve_time")
	case "40":
		q.OrderBy("shipping_time")
	case "50":
		q.OrderBy("finnshed_time")
	default:
		q.OrderBy("add_time")
	}

	sql := q.Desc().Limit(20).Offset(offset).String()
	switch {
	case memberId != 0 && orderType > -1:
		num, err = o.Raw(sql, groupId, orderState, orderType, memberId).QueryRows(&orders)
	case memberId != 0 && orderType == -1:
		num, err = o.Raw(sql, groupId, orderState, memberId).QueryRows(&orders)
	case orderType > -1 && memberId == 0:
		num, err = o.Raw(sql, groupId, orderState, orderType).QueryRows(&orders)
	case memberId == 0 && orderType == -1:
		num, err = o.Raw(sql, groupId, orderState).QueryRows(&orders)
	}
	return
}

func GetOrderByStoreId(storeId int, state int, limit int, offset int) (orders []SunOrder, count int, err error) {
	o, q := GetQueryBuilder()
	beego.Info("order state....", state)
	q = q.Select("count(*)").From("shop_order").Where("store_id = " + strconv.Itoa(storeId)).And("order_state != 20")
	if state > -1 {
		q.And("order_state=" + strconv.Itoa(state))
	}

	//err = o.Raw(q.String(), state, storeId).QueryRow(&count)
	err = o.Raw(q.String()).QueryRow(&count)
	if err != nil {
		Display("count", err)
		return
	}

	switch state {
	case 10:
		q = q.OrderBy("add_time")
	case 30:
		q = q.OrderBy("approve_time")
	case 40:
		q = q.OrderBy("shipping_time")
	case 50:
		q = q.OrderBy("finnshed_time")
	default:
		q = q.OrderBy("add_time")
	}

	sql := strings.Replace(q.Desc().Limit(limit).Offset(offset).String(), "count(*)", "*", 1)

	_, err = o.Raw(sql).QueryRows(&orders)

	if err != nil {
		Display("orders", err)
		return
	}

	return
}

// GetAllSunOrder retrieves all SunOrder matches certain condition. Returns empty list if
// no records exist
func GetAllSunOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []SunOrder, count int64, err error) {
	count = 0
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunOrder))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, count, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, count, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, count, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, count, errors.New("Error: unused 'order' fields")
		}
	}

	var orderList []SunOrder
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&orderList, fields...); err == nil {
		for _, v := range orderList {
			ml = append(ml, v)
		}
		if count, err = qs.Count(); err == nil {
			return ml, count, nil
		}
		return ml, count, err
	}
	return nil, count, err
}

// UpdateSunOrder updates SunOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunOrderById(m *SunOrder) (err error) {
	o := orm.NewOrm()
	o.Begin()
	v := SunOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
			o.Commit()
		} else {
			o.Rollback()
			return err
		}
	}
	return
}

// DeleteSunOrder deletes SunOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunOrder(id int) (err error) {
	o := orm.NewOrm()
	v := SunOrder{Id: uint64(id)}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunOrder{Id: uint64(id)}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

/**
 * 创建订单及运费
 */
func SaveOrderWithGoods(
orders []SunOrder,
goods [][]SunOrderGoods,
tagIds []string,
userId uint,
approveLevel int,
groupParentId uint,
isCart bool,
isOwnerGroup bool,
areaId int,
) bool {
	o := Orm()
	if err := o.Begin(); err == nil {
		ordersCp := len(orders) - 1
		for i, v := range orders {
			Display("order", v)

			v.BuyerId = userId
			v.AddTime = tools.GetTime()

			//创建订单
			if id, err := o.Insert(&v); err == nil {
				goodsCp := len(goods[i]) - 1
				Display("order", len(tagIds))

				//如果不是购物车订单及不是自建组，则设置审批
				if !isCart && !isOwnerGroup && v.Personal == 0 {

					//获取角色Id
					roleId, groupId, approves, err := GetAllApproveRoleIdAndGroupId(v.GroupId, groupParentId, approveLevel, v.OrderType, tagIds[i])
					Display("approve", roleId)
					if err != nil {
						Display("approve-err", err)
						o.Rollback()
						return false
					}
					//如果roleId等于0，则审批完成
					if roleId != 0 {
						_, err := o.Raw("INSERT INTO shop_approve_order (order_id, role_id, group_id, tag_ids, ctime) VALUES (?, ?, ?, ?, ?)", id, roleId, groupId, tagIds[i], tools.GetTime()).Exec()
						v.OrderState = 10
						if err != nil {
							o.Rollback()
							return false
						}
						v.Approvers = approves
					} else {
						v.OrderState = 30
					}
				} else {
					v.OrderState = 30
					v.ApproveTime = int(tools.GetTime())
				}

				//设置OrderSn
				v.OrderSn = tools.GetTimeString(2) + strconv.Itoa(int(id))
				if _, err := o.Update(&v, "order_sn", "order_state", "approvers"); err != nil {
					o.Rollback()
					return false
				}

				//设置goods,计算运费
				var goodsIds []string
				var goodsNums []int
				var amount float64
				for gi, gv := range goods[i] {

					gv.BuyerId = userId
					gv.OrderId = uint64(id)

					goodsIds = append(goodsIds, strconv.Itoa(gv.GoodsId))
					goodsNums = append(goodsNums, int(gv.GoodsNum))
					amount += gv.GoodsPrice * float64(gv.GoodsNum )

					if _, err := o.Insert(&gv); err == nil {
						Display("goods", "22222222")
						if i == ordersCp && gi == goodsCp {
							v.ShippingFee = GetTransportfee(strings.Join(goodsIds, ","), goodsNums, amount, areaId)
							v.OrderAmount = v.GoodsAmount + v.ShippingFee
							if _, err := o.Update(&v, "shipping_fee", "order_amount"); err != nil {
								o.Rollback()
								return false
							}

							if isCart {
								RemoveBy(new(SunCart), "BuyerId", userId, "GoodsId__in", goodsIds)
							}
							o.Commit()
						}
					} else {
						Display("goods", err)
						o.Rollback()
						return false
					}
				}
			} else {
				Display("order_sn", err)
				o.Rollback()
				return false
			}
		}
	} else {
		Display("order", err)
		return false
	}
	return true
}

func GetOrderById(ID int) (v *SunOrder, err error) {
	o := orm.NewOrm()
	v = new(SunOrder)
	if err = o.QueryTable(v).Filter("ID", ID).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

/**
订单发货
*/
func ChangeOrderSend(orderId int, data map[string]interface{}) bool {
	//获取订单信息
	orderInfo, err := GetSunOrderById(orderId)
	//更新order表
	orderInfo.ShippingCode = data["shippingCode"].(string)
	orderInfo.ShippingExpressId = data["shippingExpressId"].(int8)
	orderInfo.InvoiceNumber = data["invoiceNumber"].(string)
	orderInfo.InvoiceCode = data["invoiceCode"].(string)
	//orderInfo.ReciverInfo = data["receiverInfo"].(string)
	//orderInfo.ReciverName = data["reciverName"].(string)
	orderInfo.DeliverExplain = data["deliverExplain"].(string)
	orderInfo.DaddressId = data["daddressId"].(int32)
	orderInfo.ShippingTime = uint(time.Now().Unix())
	orderInfo.OrderState = constant.ORDER_STATE_SEND
	beego.Info(orderInfo)
	UpdateSunOrderById(orderInfo)
	if err != nil {
		return false
	}
	return true
}

/**
订单发货
*/
func EditShippingFee(data map[string]interface{}) bool {
	beego.Info(data)
	//获取订单信息
	orderId := int(data["OrderId"].(float64))
	orderInfo, err := GetSunOrderById(orderId)
	//更新order表
	shippingFee := data["ShippingFee"].(float64)
	orderInfo.ShippingFee = shippingFee
	orderInfo.OrderAmount = orderInfo.GoodsAmount + shippingFee
	beego.Info(orderInfo)
	UpdateSunOrderById(orderInfo)
	if err != nil {
		return false
	}
	return true
}

func CancleOrder(order *SunOrder) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	//更新订单状态
	order.OrderState = 0
	_, err = o.Update(order, "OrderState")
	if err != nil {
		o.Rollback()
		return err
	}
	//更新库存
	orderGoodsList, err := GetGoodsByOrderId(order.Id)
	for _, orderGoods := range orderGoodsList {
		if _, err = UpdateGoodsStorageByID(int(orderGoods.GoodsId), uint(orderGoods.GoodsNum)); err == nil {
			o.Commit()
		} else {
			o.Rollback()
			return err
		}
	}
	return err
}

func AffirmOrder(order *SunOrder) (err error) {
	o := Orm()
	order.OrderState = 50
	order.FinnshedTime = tools.GetTime()
	_, err = o.Update(order, "OrderState")
	return
}

type orderApprovers struct {
	Id             int    `orm:"column(member_id);auto"`
	MemberName     string `orm:"column(member_name);size(50)"`
	MemberTruename string `orm:"column(member_truename);size(20);null"`
	MemberAvatar   string `orm:"column(member_avatar);size(50);null"`
	RoleId         int    `orm:"column(role_id);size(50);null"`
}

func GetOrderApprovers(orderId int) (approverList []orderApprovers, currentApproverId string, approverIds string, err error) {
	o, q := GetQueryBuilder()
	sql := q.Select("approvers").From("shop_order").Where("order_id=?").String()
	var approvers string
	err = o.Raw(sql, orderId).QueryRow(&approvers)
	if err == nil && approvers != "" {
		approver := strings.Split(approvers, "|")
		currentApproverId = approver[1]
		groupAndApprovers := strings.Split(approver[0], ";")
		for i, groupAndApprover := range groupAndApprovers {
			if i > 0 {
				approverIds += ","
			}
			groupAndApproverSplice := strings.Split(groupAndApprover, ":")
			approverIds += groupAndApproverSplice[1]
		}
		if approverIds != "" {
			q := QueryBuilder()
			q.Select("*").From("shop_member as m").InnerJoin("shop_member_group as mg").
				On("m.member_id=mg.member_id").Where("mg.role_id").In(approverIds).And("mg.status=3")

			_, err = o.Raw(q.String()).QueryRows(&approverList)
			if err != nil {
				return
			}
		}
	}
	return
}

func GetOrderGoodsesAndTags(orderId int, groupId int) (goodses []SunOrderGoods, tags [][]Tag, err error) {
	o, q := GetQueryBuilder()
	sql := q.Select("*").From("shop_order_goods").Where("order_id=?").String()

	_, err = o.Raw(sql, orderId).QueryRows(&goodses)

	tagQuery := QueryBuilder()
	sql = tagQuery.Select("*").
		From("shop_goods_tag as tg").
		InnerJoin("shop_tag as t").On("tg.tag_id = t.tag_id").
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
