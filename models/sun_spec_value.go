package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunSpecValue struct {
	Id            int    `orm:"column(sp_value_id);auto"`
	SpValueName   string `orm:"column(sp_value_name);size(100)"`
	SpId          uint   `orm:"column(sp_id)"`
	GcId          uint   `orm:"column(gc_id)"`
	StoreId       uint   `orm:"column(store_id)"`
	SpValueColor  string `orm:"column(sp_value_color);size(10);null"`
	SpValueSort   uint8  `orm:"column(sp_value_sort)"`
	GoodsCommonid int    `orm:"column(goods_commonid);null"`
}

func (t *SunSpecValue) TableName() string {
	return "sun_spec_value"
}

func init() {
	orm.RegisterModel(new(SunSpecValue))
}

// AddSunSpecValue insert a new SunSpecValue into database and returns
// last inserted Id on success.
func AddSunSpecValue(m *SunSpecValue) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunSpecValueById retrieves SunSpecValue by Id. Returns error if
// Id doesn't exist
func GetSunSpecValueById(id int) (v *SunSpecValue, err error) {
	o := orm.NewOrm()
	v = &SunSpecValue{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunSpecValue retrieves all SunSpecValue matches certain condition. Returns empty list if
// no records exist
func GetAllSunSpecValue(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunSpecValue))
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

	var l []SunSpecValue
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

// UpdateSunSpecValue updates SunSpecValue by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunSpecValueById(m *SunSpecValue) (err error) {
	o := orm.NewOrm()
	v := SunSpecValue{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunSpecValue deletes SunSpecValue by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunSpecValue(id int) (err error) {
	o := orm.NewOrm()
	v := SunSpecValue{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunSpecValue{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//GetAllSpecValueByGoodsCommonID 根据goodsCommonId获取所有specValue
func GetAllSpecValueByGoodsCommonID(ids []int, fields []string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunSpecValue))
	qs = qs.Filter("goods_commonid__in", ids)
	var l []SunSpecValue
	if _, err := qs.All(&l); err == nil {
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

//GetAllSpecValueByGoodsClassID 根据goodsClasId获取所有specValue
func GetAllSpecValueByGoodsClassID(ids []int, fields []string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunSpecValue))
	qs = qs.Filter("gc_id__in", ids)
	var l []SunSpecValue
	if _, err := qs.All(&l); err == nil {
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
