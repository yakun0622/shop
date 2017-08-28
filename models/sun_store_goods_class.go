package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunStoreGoodsClass struct {
	Id          int    `orm:"column(stc_id);auto"`
	StcName     string `orm:"column(stc_name);size(50)"`
	StcParentId int    `orm:"column(stc_parent_id)"`
	StcState    int8   `orm:"column(stc_state)"`
	StoreId     int    `orm:"column(store_id)"`
	StcSort     int    `orm:"column(stc_sort)"`
}

func (t *SunStoreGoodsClass) TableName() string {
	return "sun_store_goods_class"
}

func init() {
	orm.RegisterModel(new(SunStoreGoodsClass))
}

// AddSunStoreGoodsClass insert a new SunStoreGoodsClass into database and returns
// last inserted Id on success.
func AddSunStoreGoodsClass(m *SunStoreGoodsClass) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunStoreGoodsClassById retrieves SunStoreGoodsClass by Id. Returns error if
// Id doesn't exist
func GetSunStoreGoodsClassById(id int) (v *SunStoreGoodsClass, err error) {
	o := orm.NewOrm()
	v = &SunStoreGoodsClass{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunStoreGoodsClass retrieves all SunStoreGoodsClass matches certain condition. Returns empty list if
// no records exist
func GetAllSunStoreGoodsClass(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunStoreGoodsClass))
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

	var l []SunStoreGoodsClass
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

// UpdateSunStoreGoodsClass updates SunStoreGoodsClass by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunStoreGoodsClassById(m *SunStoreGoodsClass) (err error) {
	o := orm.NewOrm()
	v := SunStoreGoodsClass{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunStoreGoodsClass deletes SunStoreGoodsClass by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunStoreGoodsClass(id int) (err error) {
	o := orm.NewOrm()
	v := SunStoreGoodsClass{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunStoreGoodsClass{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
