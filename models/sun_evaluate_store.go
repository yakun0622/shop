package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunEvaluateStore struct {
	Id                  int    `orm:"column(seval_id);auto"`
	SevalOrderid        uint   `orm:"column(seval_orderid)"`
	SevalOrderno        string `orm:"column(seval_orderno);size(100)"`
	SevalAddtime        uint   `orm:"column(seval_addtime)"`
	SevalStoreid        uint   `orm:"column(seval_storeid)"`
	SevalStorename      string `orm:"column(seval_storename);size(100)"`
	SevalMemberid       uint   `orm:"column(seval_memberid)"`
	SevalMembername     string `orm:"column(seval_membername);size(100)"`
	SevalDesccredit     uint8  `orm:"column(seval_desccredit)"`
	SevalServicecredit  uint8  `orm:"column(seval_servicecredit)"`
	SevalDeliverycredit uint8  `orm:"column(seval_deliverycredit)"`
}

func (t *SunEvaluateStore) TableName() string {
	return "shop_evaluate_store"
}

func init() {
	orm.RegisterModel(new(SunEvaluateStore))
}

// AddSunEvaluateStore insert a new SunEvaluateStore into database and returns
// last inserted Id on success.
func AddSunEvaluateStore(m *SunEvaluateStore) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunEvaluateStoreById retrieves SunEvaluateStore by Id. Returns error if
// Id doesn't exist
func GetSunEvaluateStoreById(id int) (v *SunEvaluateStore, err error) {
	o := orm.NewOrm()
	v = &SunEvaluateStore{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunEvaluateStore retrieves all SunEvaluateStore matches certain condition. Returns empty list if
// no records exist
func GetAllSunEvaluateStore(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunEvaluateStore))
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

	var l []SunEvaluateStore
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

// UpdateSunEvaluateStore updates SunEvaluateStore by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunEvaluateStoreById(m *SunEvaluateStore) (err error) {
	o := orm.NewOrm()
	v := SunEvaluateStore{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunEvaluateStore deletes SunEvaluateStore by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunEvaluateStore(id int) (err error) {
	o := orm.NewOrm()
	v := SunEvaluateStore{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunEvaluateStore{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
