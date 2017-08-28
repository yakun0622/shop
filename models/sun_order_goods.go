package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type SunOrderGoods struct {
	Id           uint64  `orm:"column(rec_id);auto"`
	OrderId      uint64  `orm:"column(order_id)"`
	GoodsId      int     `orm:"column(goods_id)"`
	GoodsName    string  `orm:"column(goods_name);size(50)"`
	GoodsPrice   float64 `orm:"column(goods_price);digits(10);decimals(2)"`
	GoodsNum     uint16  `orm:"column(goods_num)"`
	GoodsImage   string  `orm:"column(goods_image);size(100);null"`
	StoreId      uint    `orm:"column(store_id)"`
	BuyerId      uint    `orm:"column(buyer_id)"`
	GoodsType    uint    `orm:"column(goods_type)"`
	PromotionsId uint32  `orm:"column(promotions_id)"`
	CommisRate   uint16  `orm:"column(commis_rate)"`
	GcId         uint32  `orm:"column(gc_id)"`
}

const OrderGoodsTableName  = "sun_order_goods"

func init() {
	orm.RegisterModel(new(SunOrderGoods))
}

func GetGoodsByOrderId(orderId uint64) (list []SunOrderGoods, err error)  {
	o, query := GetQueryBuilder()
	_, err = o.Raw(query.Select("*").From(OrderGoodsTableName).Where("order_id=?").String(), orderId).QueryRows(&list)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return list, nil
}

// AddSunOrderGoods insert a new SunOrderGoods into database and returns
// last inserted Id on success.
func AddSunOrderGoods(m *SunOrderGoods) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunOrderGoodsById retrieves SunOrderGoods by Id. Returns error if
// Id doesn't exist
func GetSunOrderGoodsById(id int) (v *SunOrderGoods, err error) {
	o := orm.NewOrm()
	v = &SunOrderGoods{Id: uint64(id)}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetSunOrderGoodsByOrderID 根据订单ID获取该订单所有商品的信息
func GetSunOrderGoodsByOrderID(orderID int) (list []SunOrderGoods, err error) {
	o := orm.NewOrm()
	if _, err := o.QueryTable(SunOrderGoods{}).Filter("OrderId", orderID).All(&list); err == nil {
		return list, nil
	}
	return nil, err
}

//GetSingeSunOrderGoodsByOrderID  查询订单中单个商品信息
func GetSingeSunOrderGoodsByOrderID(orderID int, GoodsID int) (v SunOrderGoods, err error) {
	o := orm.NewOrm()
	//v = &SunOrderGoods{OrderId:uint64(orderID), GoodsId:GoodsID}
	//if err = o.Read(v); err == nil{
	var list []SunOrderGoods
	if _, err := o.QueryTable(SunOrderGoods{}).Filter("OrderId", orderID).Filter("GoodsId", GoodsID).All(&list); err == nil {
		if len(list) > 0 {
			v = list[0]
		}
	}
	return v, err
}

// GetAllSunOrderGoods retrieves all SunOrderGoods matches certain condition. Returns empty list if
// no records exist
func GetAllSunOrderGoods(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunOrderGoods))
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SunOrderGoods
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSunOrderGoods updates SunOrderGoods by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunOrderGoodsById(m *SunOrderGoods) (err error) {
	o := orm.NewOrm()
	v := SunOrderGoods{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunOrderGoods deletes SunOrderGoods by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunOrderGoods(id int) (err error) {
	o := orm.NewOrm()
	v := SunOrderGoods{Id: uint64(id)}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunOrderGoods{Id: uint64(id)}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
