package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunGoodsClassStaple struct {
	Id         int    `orm:"column(staple_id);auto"`
	StapleName string `orm:"column(staple_name);size(255)"`
	GcId1      uint   `orm:"column(gc_id_1)"`
	GcId2      uint   `orm:"column(gc_id_2)"`
	GcId3      uint   `orm:"column(gc_id_3)"`
	TypeId     uint   `orm:"column(type_id)"`
	MemberId   uint   `orm:"column(member_id)"`
	Counter    uint   `orm:"column(counter)"`
}

func (t *SunGoodsClassStaple) TableName() string {
	return "sun_goods_class_staple"
}

func init() {
	orm.RegisterModel(new(SunGoodsClassStaple))
}

// AddSunGoodsClassStaple insert a new SunGoodsClassStaple into database and returns
// last inserted Id on success.
func AddSunGoodsClassStaple(m *SunGoodsClassStaple) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunGoodsClassStapleById retrieves SunGoodsClassStaple by Id. Returns error if
// Id doesn't exist
func GetSunGoodsClassStapleById(id int) (v *SunGoodsClassStaple, err error) {
	o := orm.NewOrm()
	v = &SunGoodsClassStaple{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunGoodsClassStaple retrieves all SunGoodsClassStaple matches certain condition. Returns empty list if
// no records exist
func GetAllSunGoodsClassStaple(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsClassStaple))
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

	var l []SunGoodsClassStaple
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

// UpdateSunGoodsClassStaple updates SunGoodsClassStaple by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunGoodsClassStapleById(m *SunGoodsClassStaple) (err error) {
	o := orm.NewOrm()
	v := SunGoodsClassStaple{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunGoodsClassStaple deletes SunGoodsClassStaple by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunGoodsClassStaple(id int) (err error) {
	o := orm.NewOrm()
	v := SunGoodsClassStaple{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunGoodsClassStaple{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
