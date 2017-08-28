package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunOrderBillGroup struct {
	Id                   int     `orm:"column(obg_id);auto"`
	ObgNo                int64   `orm:"column(obg_no)"`
	ObgStartDate         int     `orm:"column(obg_start_date)"`
	ObgEndDate           int     `orm:"column(obg_end_date)"`
	ObgOrderTotals       float64 `orm:"column(obg_order_totals);digits(10);decimals(2)"`
	ObgShippingTotals    float64 `orm:"column(obg_shipping_totals);digits(10);decimals(2)"`
	ObgOrderReturnTotals float64 `orm:"column(obg_order_return_totals);digits(10);decimals(2)"`
	ObgResultTotals      float64 `orm:"column(obg_result_totals);digits(10);decimals(2)"`
	ObgCreateDate        int     `orm:"column(obg_create_date);null"`
	ObgStoreId           int     `orm:"column(obg_store_id)"`
	ObgStoreName         string  `orm:"column(obg_store_name);size(50);null"`
	ObgGroupid           int     `orm:"column(obg_groupid)"`
	OsgId                int     `orm:"column(osg_id)"`
	OsgMonth             int32   `orm:"column(osg_month)"`
	ObgPeriodType        int     `orm:"column(obg_period_type);int(1)"`
}

func (t *SunOrderBillGroup) TableName() string {
	return "sun_order_bill_group"
}

func init() {
	orm.RegisterModel(new(SunOrderBillGroup))
}

// AddSunOrderBillGroup insert a new SunOrderBillGroup into database and returns
// last inserted Id on success.
func AddSunOrderBillGroup(m *SunOrderBillGroup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunOrderBillGroupById retrieves SunOrderBillGroup by Id. Returns error if
// Id doesn't exist
func GetSunOrderBillGroupById(id int) (v *SunOrderBillGroup, err error) {
	o := orm.NewOrm()
	v = &SunOrderBillGroup{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetSunOrderBillGroupByGroupId(ids []int, period_type int) (v []SunOrderBillGroup, count int64, err error) {
	if len(ids) <= 0 {
		return nil, 0, errors.New("no ids")
	}
	o := orm.NewOrm()
	sql := "SELECT * FROM sun_order_bill_group WHERE obg_groupid IN ("
	for index := range ids {
		if index != 0 {
			sql += ","
		}
		sql += "?"
	}
	sql += ") AND obg_period_type =? Order By obg_create_date DESC"
	if count, err := o.Raw(sql, ids, period_type).QueryRows(&v); err == nil {
		return v, count, nil
	}
	return nil, 0, err
}

// GetAllSunOrderBillGroup retrieves all SunOrderBillGroup matches certain condition. Returns empty list if
// no records exist
func GetAllSunOrderBillGroup(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunOrderBillGroup))
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

	var l []SunOrderBillGroup
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

// UpdateSunOrderBillGroup updates SunOrderBillGroup by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunOrderBillGroupById(m *SunOrderBillGroup) (err error) {
	o := orm.NewOrm()
	v := SunOrderBillGroup{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunOrderBillGroup deletes SunOrderBillGroup by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunOrderBillGroup(id int) (err error) {
	o := orm.NewOrm()
	v := SunOrderBillGroup{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunOrderBillGroup{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
