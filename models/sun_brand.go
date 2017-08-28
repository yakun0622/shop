package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunBrand struct {
	Id             int    `orm:"column(brand_id);auto"`
	BrandName      string `orm:"column(brand_name);size(100);null"`
	BrandInitial   string `orm:"column(brand_initial);size(1)"`
	BrandClass     string `orm:"column(brand_class);size(50);null"`
	BrandPic       string `orm:"column(brand_pic);size(100);null"`
	BrandSort      uint8  `orm:"column(brand_sort);null"`
	BrandRecommend int8   `orm:"column(brand_recommend);null"`
	StoreId        uint   `orm:"column(store_id)"`
	BrandApply     int8   `orm:"column(brand_apply)"`
	ClassId        uint   `orm:"column(class_id);null"`
	ShowType       int8   `orm:"column(show_type)"`
}

func (t *SunBrand) TableName() string {
	return "shop_brand"
}

func init() {
	orm.RegisterModel(new(SunBrand))
}

// AddSunBrand insert a new SunBrand into database and returns
// last inserted Id on success.
func AddSunBrand(m *SunBrand) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunBrandById retrieves SunBrand by Id. Returns error if
// Id doesn't exist
func GetSunBrandById(id int) (v *SunBrand, err error) {
	o := orm.NewOrm()
	v = &SunBrand{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunBrand retrieves all SunBrand matches certain condition. Returns empty list if
// no records exist
func GetAllSunBrand(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunBrand))
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

	var l []SunBrand
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

// UpdateSunBrand updates SunBrand by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunBrandById(m *SunBrand) (err error) {
	o := orm.NewOrm()
	v := SunBrand{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunBrand deletes SunBrand by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunBrand(id int) (err error) {
	o := orm.NewOrm()
	v := SunBrand{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunBrand{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//GetAllSunBrandByClassID 获取某个分类下所有的品牌
func GetAllSunBrandByClassID(ids []int, fields ...string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunBrand))
	qs = qs.Filter("class_id__in", ids)
	var l []SunBrand
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
