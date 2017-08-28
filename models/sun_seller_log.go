package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunSellerLog struct {
	Id            int    `orm:"column(log_id);auto"`
	LogContent    string `orm:"column(log_content);size(50)"`
	LogTime       uint   `orm:"column(log_time)"`
	LogSellerId   uint   `orm:"column(log_seller_id)"`
	LogSellerName string `orm:"column(log_seller_name);size(50)"`
	LogStoreId    uint   `orm:"column(log_store_id)"`
	LogSellerIp   string `orm:"column(log_seller_ip);size(50)"`
	LogUrl        string `orm:"column(log_url);size(50)"`
	LogState      uint8  `orm:"column(log_state)"`
}

func (t *SunSellerLog) TableName() string {
	return "sun_seller_log"
}

func init() {
	orm.RegisterModel(new(SunSellerLog))
}

// AddSunSellerLog insert a new SunSellerLog into database and returns
// last inserted Id on success.
func AddSunSellerLog(m *SunSellerLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunSellerLogById retrieves SunSellerLog by Id. Returns error if
// Id doesn't exist
func GetSunSellerLogById(id int) (v *SunSellerLog, err error) {
	o := orm.NewOrm()
	v = &SunSellerLog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunSellerLog retrieves all SunSellerLog matches certain condition. Returns empty list if
// no records exist
func GetAllSunSellerLog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunSellerLog))
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

	var l []SunSellerLog
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

// UpdateSunSellerLog updates SunSellerLog by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunSellerLogById(m *SunSellerLog) (err error) {
	o := orm.NewOrm()
	v := SunSellerLog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunSellerLog deletes SunSellerLog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunSellerLog(id int) (err error) {
	o := orm.NewOrm()
	v := SunSellerLog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunSellerLog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
