package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunOrderStatisGroup struct {
	Id                   int     `orm:"column(osg_id);auto"`
	OsgMonth             uint32  `orm:"column(osg_month)"`
	OsgYear              int16   `orm:"column(osg_year);null"`
	OsgStartDate         int     `orm:"column(osg_start_date)"`
	OsgEndDate           int     `orm:"column(osg_end_date)"`
	OsgOrderTotals       float64 `orm:"column(osg_order_totals);digits(10);decimals(2)"`
	OsgShippingTotals    float64 `orm:"column(osg_shipping_totals);digits(10);decimals(2)"`
	OsgOrderReturnTotals float64 `orm:"column(osg_order_return_totals);digits(10);decimals(2)"`
	OsgResultTotals      float64 `orm:"column(osg_result_totals);digits(10);decimals(2)"`
	OsgCreateDate        int     `orm:"column(osg_create_date);null"`
	OsgGroupid           int     `orm:"column(osg_groupid)"`
	OsgState             string  `orm:"column(osg_state);null"`
}

func (t *SunOrderStatisGroup) TableName() string {
	return "sun_order_statis_group"
}

func init() {
	orm.RegisterModel(new(SunOrderStatisGroup))
}

// AddSunOrderStatisGroup insert a new SunOrderStatisGroup into database and returns
// last inserted Id on success.
func AddSunOrderStatisGroup(m *SunOrderStatisGroup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunOrderStatisGroupById retrieves SunOrderStatisGroup by Id. Returns error if
// Id doesn't exist
func GetSunOrderStatisGroupById(id int) (v *SunOrderStatisGroup, err error) {
	o := orm.NewOrm()
	v = &SunOrderStatisGroup{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunOrderStatisGroup retrieves all SunOrderStatisGroup matches certain condition. Returns empty list if
// no records exist
func GetAllSunOrderStatisGroup(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunOrderStatisGroup))
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

	var l []SunOrderStatisGroup
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

// UpdateSunOrderStatisGroup updates SunOrderStatisGroup by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunOrderStatisGroupById(m *SunOrderStatisGroup) (err error) {
	o := orm.NewOrm()
	v := SunOrderStatisGroup{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunOrderStatisGroup deletes SunOrderStatisGroup by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunOrderStatisGroup(id int) (err error) {
	o := orm.NewOrm()
	v := SunOrderStatisGroup{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunOrderStatisGroup{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
