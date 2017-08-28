package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunGoodsClassTag struct {
	Id         int    `orm:"column(gc_tag_id);auto"`
	GcId1      uint   `orm:"column(gc_id_1)"`
	GcId2      uint   `orm:"column(gc_id_2)"`
	GcId3      uint   `orm:"column(gc_id_3)"`
	GcTagName  string `orm:"column(gc_tag_name);size(255)"`
	GcTagValue string `orm:"column(gc_tag_value)"`
	GcId       uint   `orm:"column(gc_id)"`
	TypeId     uint   `orm:"column(type_id)"`
}

func (t *SunGoodsClassTag) TableName() string {
	return "shop_goods_class_tag"
}

func init() {
	orm.RegisterModel(new(SunGoodsClassTag))
}

// AddSunGoodsClassTag insert a new SunGoodsClassTag into database and returns
// last inserted Id on success.
func AddSunGoodsClassTag(m *SunGoodsClassTag) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunGoodsClassTagById retrieves SunGoodsClassTag by Id. Returns error if
// Id doesn't exist
func GetSunGoodsClassTagById(id int) (v *SunGoodsClassTag, err error) {
	o := orm.NewOrm()
	v = &SunGoodsClassTag{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunGoodsClassTag retrieves all SunGoodsClassTag matches certain condition. Returns empty list if
// no records exist
func GetAllSunGoodsClassTag(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsClassTag))
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

	var l []SunGoodsClassTag
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

// UpdateSunGoodsClassTag updates SunGoodsClassTag by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunGoodsClassTagById(m *SunGoodsClassTag) (err error) {
	o := orm.NewOrm()
	v := SunGoodsClassTag{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunGoodsClassTag deletes SunGoodsClassTag by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunGoodsClassTag(id int) (err error) {
	o := orm.NewOrm()
	v := SunGoodsClassTag{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunGoodsClassTag{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
