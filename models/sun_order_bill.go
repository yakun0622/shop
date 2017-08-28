package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunOrderBill struct {
	Id                   int     `orm:"column(ob_no);auto"`
	ObStartDate          int     `orm:"column(ob_start_date)"`
	ObEndDate            int     `orm:"column(ob_end_date)"`
	ObOrderTotals        float64 `orm:"column(ob_order_totals);digits(10);decimals(2)"`
	ObShippingTotals     float64 `orm:"column(ob_shipping_totals);digits(10);decimals(2)"`
	ObOrderReturnTotals  float64 `orm:"column(ob_order_return_totals);digits(10);decimals(2)"`
	ObCommisTotals       float64 `orm:"column(ob_commis_totals);digits(10);decimals(2)"`
	ObCommisReturnTotals float64 `orm:"column(ob_commis_return_totals);digits(10);decimals(2)"`
	ObStoreCostTotals    float64 `orm:"column(ob_store_cost_totals);digits(10);decimals(2)"`
	ObResultTotals       float64 `orm:"column(ob_result_totals);digits(10);decimals(2)"`
	ObCreateDate         int     `orm:"column(ob_create_date);null"`
	OsMonth              uint32  `orm:"column(os_month)"`
	ObState              string  `orm:"column(ob_state);null"`
	ObPayDate            int     `orm:"column(ob_pay_date);null"`
	ObPayContent         string  `orm:"column(ob_pay_content);size(200);null"`
	ObStoreId            int     `orm:"column(ob_store_id)"`
	ObStoreName          string  `orm:"column(ob_store_name);size(50);null"`
}

func (t *SunOrderBill) TableName() string {
	return "sun_order_bill"
}

func init() {
	orm.RegisterModel(new(SunOrderBill))
}

// AddSunOrderBill insert a new SunOrderBill into database and returns
// last inserted Id on success.
func AddSunOrderBill(m *SunOrderBill) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunOrderBillById retrieves SunOrderBill by Id. Returns error if
// Id doesn't exist
func GetSunOrderBillById(id int) (v *SunOrderBill, err error) {
	o := orm.NewOrm()
	v = &SunOrderBill{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunOrderBill retrieves all SunOrderBill matches certain condition. Returns empty list if
// no records exist
func GetAllSunOrderBill(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunOrderBill))
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

	var l []SunOrderBill
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

// UpdateSunOrderBill updates SunOrderBill by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunOrderBillById(m *SunOrderBill) (err error) {
	o := orm.NewOrm()
	v := SunOrderBill{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunOrderBill deletes SunOrderBill by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunOrderBill(id int) (err error) {
	o := orm.NewOrm()
	v := SunOrderBill{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunOrderBill{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
