package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunOrderCommon struct {
	Id                int    `orm:"column(order_id);auto"`
	StoreId           uint   `orm:"column(store_id)"`
	ShippingTime      uint   `orm:"column(shipping_time)"`
	ShippingExpressId int8   `orm:"column(shipping_express_id)"`
	EvaluationTime    uint   `orm:"column(evaluation_time)"`
	EvalsellerState   string `orm:"column(evalseller_state)"`
	EvalsellerTime    uint   `orm:"column(evalseller_time)"`
	OrderMessage      string `orm:"column(order_message);size(300);null"`
	OrderPointscount  int    `orm:"column(order_pointscount)"`
	VoucherPrice      int    `orm:"column(voucher_price);null"`
	VoucherCode       string `orm:"column(voucher_code);size(32);null"`
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
	PromotionInfo     string `orm:"column(promotion_info);size(500);null"`
	DlyoPickupCode    string `orm:"column(dlyo_pickup_code);size(4);null"`
}

func (t *SunOrderCommon) TableName() string {
	return "shop_order_common"
}

func init() {
	orm.RegisterModel(new(SunOrderCommon))
}

// AddSunOrderCommon insert a new SunOrderCommon into database and returns
// last inserted Id on success.
func AddSunOrderCommon(m *SunOrderCommon) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunOrderCommonById retrieves SunOrderCommon by Id. Returns error if
// Id doesn't exist
func GetSunOrderCommonById(id int) (v *SunOrderCommon, err error) {
	o := orm.NewOrm()
	v = &SunOrderCommon{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunOrderCommon retrieves all SunOrderCommon matches certain condition. Returns empty list if
// no records exist
func GetAllSunOrderCommon(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunOrderCommon))
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

	var l []SunOrderCommon
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

// UpdateSunOrderCommon updates SunOrderCommon by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunOrderCommonById(m *SunOrderCommon) (err error) {
	o := orm.NewOrm()
	v := SunOrderCommon{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunOrderCommon deletes SunOrderCommon by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunOrderCommon(id int) (err error) {
	o := orm.NewOrm()
	v := SunOrderCommon{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunOrderCommon{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
