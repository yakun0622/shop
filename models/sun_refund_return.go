package models

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunRefundReturn struct {
	Id             int     `orm:"column(refund_id);auto"`
	OrderId        uint    `orm:"column(order_id)"`
	OrderSn        string  `orm:"column(order_sn);size(50)"`
	RefundSn       string  `orm:"column(refund_sn);size(50)"`
	StoreId        uint    `orm:"column(store_id)"`
	StoreName      string  `orm:"column(store_name);size(20)"`
	BuyerId        uint    `orm:"column(buyer_id)"`
	BuyerName      string  `orm:"column(buyer_name);size(50)"`
	//GoodsId        uint    `orm:"column(goods_id)"`
	Goods        	*SunGoods    `orm:"column(goods_id);rel(one);on_delete(do_nothing)"`
	OrderGoodsId   uint    `orm:"column(order_goods_id);null"`
	GoodsName      string  `orm:"column(goods_name);size(50)"`
	GoodsNum       uint    `orm:"column(goods_num);null"`
	RefundAmount   float64 `orm:"column(refund_amount);null;digits(10);decimals(2)"`
	GoodsImage     string  `orm:"column(goods_image);size(100);null"`
	OrderGoodsType uint8   `orm:"column(order_goods_type);null"`
	RefundType     uint8   `orm:"column(refund_type);null"`
	SellerState    uint8   `orm:"column(seller_state);null"`
	RefundState    uint8   `orm:"column(refund_state);null"`
	ReturnType     uint8   `orm:"column(return_type);null"`
	OrderLock      uint8   `orm:"column(order_lock);null"`
	GoodsState     uint8   `orm:"column(goods_state);null"`
	AddTime        uint    `orm:"column(add_time)"`
	SellerTime     uint    `orm:"column(seller_time);null"`
	AdminTime      uint    `orm:"column(admin_time);null"`
	ReasonId       uint    `orm:"column(reason_id);null"`
	ReasonInfo     string  `orm:"column(reason_info);size(300);null"`
	PicInfo        string  `orm:"column(pic_info);size(300);null"`
	BuyerMessage   string  `orm:"column(buyer_message);size(300);null"`
	SellerMessage  string  `orm:"column(seller_message);size(300);null"`
	AdminMessage   string  `orm:"column(admin_message);size(300);null"`
	ExpressId      uint8   `orm:"column(express_id);null"`
	InvoiceNo      string  `orm:"column(invoice_no);size(50);null"`
	ShipTime       uint    `orm:"column(ship_time);null"`
	DelayTime      uint    `orm:"column(delay_time);null"`
	ReceiveTime    uint    `orm:"column(receive_time);null"`
	ReceiveMessage string  `orm:"column(receive_message);size(300);null"`
	CommisRate     int16   `orm:"column(commis_rate);null"`
	GroupId        int     `orm:"column(group_id);null"`
	GroupName      string  `orm:"column(group_name);size(255);null"`
}

func (t *SunRefundReturn) TableName() string {
	return "shop_refund_return"
}

func init() {
	orm.RegisterModel(new(SunRefundReturn))
}

// AddSunRefundReturn insert a new SunRefundReturn into database and returns
// last inserted Id on success.
func AddSunRefundReturn(m *SunRefundReturn) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunRefundReturnById retrieves SunRefundReturn by Id. Returns error if
// Id doesn't exist
func GetSunRefundReturnById(id int) (v *SunRefundReturn, err error) {
	o := orm.NewOrm()
	v = &SunRefundReturn{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunRefundReturn retrieves all SunRefundReturn matches certain condition. Returns empty list if
// no records exist
func GetAllSunRefundReturn(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []SunRefundReturn,count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunRefundReturn))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	sortFields, err = GetSortFields(sortby, order)

	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).RelatedSel().All(&ml, fields...); err == nil {
		if count, err = qs.RelatedSel().Count(); err == nil {
			return ml, count, nil
		}
		return ml, count, nil
	}
	return nil, count, err
}

// UpdateSunRefundReturn updates SunRefundReturn by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunRefundReturnById(m *SunRefundReturn) (err error) {
	o := orm.NewOrm()
	v := SunRefundReturn{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunRefundReturn deletes SunRefundReturn by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunRefundReturn(id int) (err error) {
	o := orm.NewOrm()
	v := SunRefundReturn{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunRefundReturn{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
