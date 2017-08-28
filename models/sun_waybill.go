package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunWaybill struct {
	Id            int    `orm:"column(waybill_id);auto"`
	WaybillName   string `orm:"column(waybill_name);size(50)"`
	WaybillImage  string `orm:"column(waybill_image);size(50)"`
	WaybillWidth  uint   `orm:"column(waybill_width)"`
	WaybillHeight uint   `orm:"column(waybill_height)"`
	WaybillData   string `orm:"column(waybill_data);size(2000);null"`
	WaybillUsable uint8  `orm:"column(waybill_usable)"`
	WaybillTop    int    `orm:"column(waybill_top)"`
	WaybillLeft   int    `orm:"column(waybill_left)"`
	ExpressId     uint8  `orm:"column(express_id)"`
	ExpressName   string `orm:"column(express_name);size(50)"`
	StoreId       uint   `orm:"column(store_id)"`
}

func (t *SunWaybill) TableName() string {
	return "sun_waybill"
}

func init() {
	orm.RegisterModel(new(SunWaybill))
}

// AddSunWaybill insert a new SunWaybill into database and returns
// last inserted Id on success.
func AddSunWaybill(m *SunWaybill) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunWaybillById retrieves SunWaybill by Id. Returns error if
// Id doesn't exist
func GetSunWaybillById(id int) (v *SunWaybill, err error) {
	o := orm.NewOrm()
	v = &SunWaybill{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunWaybill retrieves all SunWaybill matches certain condition. Returns empty list if
// no records exist
func GetAllSunWaybill(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunWaybill))
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

	var l []SunWaybill
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

// UpdateSunWaybill updates SunWaybill by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunWaybillById(m *SunWaybill) (err error) {
	o := orm.NewOrm()
	v := SunWaybill{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunWaybill deletes SunWaybill by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunWaybill(id int) (err error) {
	o := orm.NewOrm()
	v := SunWaybill{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunWaybill{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
