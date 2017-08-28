package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunGoodsClass struct {
	Id            int     `orm:"column(gc_id);auto"`
	GcName        string  `orm:"column(gc_name);size(100)"`
	TypeName      string  `orm:"column(type_name);size(100)"`
	GcParentId    uint    `orm:"column(gc_parent_id)"`
	CommisRate    float32 `orm:"column(commis_rate)"`
	GcSort        uint8   `orm:"column(gc_sort)"`
	GcVirtual     uint8   `orm:"column(gc_virtual)"`
	GcTitle       string  `orm:"column(gc_title);size(200)"`
	GcKeywords    string  `orm:"column(gc_keywords);size(255)"`
	GcDescription string  `orm:"column(gc_description);size(255)"`
}

func (t *SunGoodsClass) TableName() string {
	return "shop_goods_class"
}

func init() {
	orm.RegisterModel(new(SunGoodsClass))
}

// AddSunGoodsClass insert a new SunGoodsClass into database and returns
// last inserted Id on success.
func AddSunGoodsClass(m *SunGoodsClass) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunGoodsClassById retrieves SunGoodsClass by Id. Returns error if
// Id doesn't exist
func GetSunGoodsClassById(id int) (v *SunGoodsClass, err error) {
	o := orm.NewOrm()
	v = &SunGoodsClass{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunGoodsClass retrieves all SunGoodsClass matches certain condition. Returns empty list if
// no records exist
func GetAllSunGoodsClass(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsClass))
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

	var l []SunGoodsClass
	qs = qs.OrderBy(sortFields...)

	if limit == 0 {
		if _, err := qs.All(&l, fields...); err == nil {
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
	} else {
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
	}

	return nil, err
}

// UpdateSunGoodsClass updates SunGoodsClass by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunGoodsClassById(m *SunGoodsClass) (err error) {
	o := orm.NewOrm()
	v := SunGoodsClass{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunGoodsClass deletes SunGoodsClass by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunGoodsClass(id int) (err error) {
	o := orm.NewOrm()
	v := SunGoodsClass{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunGoodsClass{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//GetGoodsClassChildList 获取分类下的所有子类，包含自己，首位为分类自己的ID
func GetGoodsClassChildList(id int) ([]int, error) {
	var lists []orm.ParamsList
	o := orm.NewOrm()
	num, err := o.Raw("SELECT getGoosClassChild(?) as ids", id).ValuesList(&lists)
	if err == nil && num > 0 {
		idList := strings.Split(lists[0][0].(string), ",")
		ids := make([]int, len(idList)-1)
		for index, id := range idList {
			if id != "$" {
				ids[index-1], _ = strconv.Atoi(id)
			}
		}
		return ids, nil
	}
	return nil, err
}

func GetAllSunGoodsClassSimple() (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsClass))
	qs = qs.OrderBy("gc_sort")
	var l []SunGoodsClass
	if _, err := qs.All(&l); err == nil {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}
