package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunOrderLog struct {
	Id            int    `orm:"column(log_id);auto"`
	OrderId       int    `orm:"column(order_id)"`
	LogMsg        string `orm:"column(log_msg);size(150);null"`
	LogTime       uint   `orm:"column(log_time)"`
	LogRole       string `orm:"column(log_role);size(2)"`
	LogUser       string `orm:"column(log_user);size(30);null"`
	LogOrderstate string `orm:"column(log_orderstate);null"`
}

func (t *SunOrderLog) TableName() string {
	return "shop_order_log"
}

func init() {
	orm.RegisterModel(new(SunOrderLog))
}

// AddSunOrderLog insert a new SunOrderLog into database and returns
// last inserted Id on success.
func AddSunOrderLog(m *SunOrderLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunOrderLogById retrieves SunOrderLog by Id. Returns error if
// Id doesn't exist
func GetSunOrderLogById(id int) (v *SunOrderLog, err error) {
	o := orm.NewOrm()
	v = &SunOrderLog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunOrderLog retrieves all SunOrderLog matches certain condition. Returns empty list if
// no records exist
func GetAllSunOrderLog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunOrderLog))
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

	var l []SunOrderLog
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

// UpdateSunOrderLog updates SunOrderLog by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunOrderLogById(m *SunOrderLog) (err error) {
	o := orm.NewOrm()
	v := SunOrderLog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunOrderLog deletes SunOrderLog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunOrderLog(id int) (err error) {
	o := orm.NewOrm()
	v := SunOrderLog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunOrderLog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
