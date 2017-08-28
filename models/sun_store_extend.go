package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunStoreExtend struct {
	Id              int    `orm:"column(store_id);pk"`
	Express         string `orm:"column(express);null"`
	Pricerange      string `orm:"column(pricerange);null"`
	Orderpricerange string `orm:"column(orderpricerange);null"`
}

func (t *SunStoreExtend) TableName() string {
	return "shop_store_extend"
}

func init() {
	orm.RegisterModel(new(SunStoreExtend))
}

// AddSunStoreExtend insert a new SunStoreExtend into database and returns
// last inserted Id on success.
func AddSunStoreExtend(m *SunStoreExtend) (id int64, err error) {
	o := orm.NewOrm()
	v := SunStoreExtend{Id: m.Id}
	//存在则更新，否则添加
	if err = o.Read(&v); err == nil {
		v.Express = m.Express
		_, err = o.Update(&v)
	} else {
		id, err = o.Insert(&m)
	}
	return
}

// GetSunStoreExtendById retrieves SunStoreExtend by Id. Returns error if
// Id doesn't exist
func GetSunStoreExtendById(id int) (v *SunStoreExtend, err error) {
	o := orm.NewOrm()
	v = &SunStoreExtend{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunStoreExtend retrieves all SunStoreExtend matches certain condition. Returns empty list if
// no records exist
func GetAllSunStoreExtend(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []SunStoreExtend, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunStoreExtend))
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

	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&ml, fields...); err == nil {
		return ml, nil
	}
	return nil, err
}

// UpdateSunStoreExtend updates SunStoreExtend by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunStoreExtendById(m *SunStoreExtend) (err error) {
	o := orm.NewOrm()
	v := SunStoreExtend{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunStoreExtend deletes SunStoreExtend by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunStoreExtend(id int) (err error) {
	o := orm.NewOrm()
	v := SunStoreExtend{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunStoreExtend{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetStoreExtendByStoreId(storeId uint, fields []string) (ml []SunStoreExtend, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunStoreExtend))
	qs = qs.Filter("Id", storeId)

	if _, err := qs.All(&ml, fields...); err == nil {
		return ml, nil
	}
	return nil, err
}
