package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunStoreMsg struct {
	Id        int    `orm:"column(sm_id);auto"`
	SmtCode   string `orm:"column(smt_code);size(100)"`
	StoreId   uint   `orm:"column(store_id)"`
	SmContent string `orm:"column(sm_content);size(255)"`
	SmAddtime uint   `orm:"column(sm_addtime)"`
	SmReadids string `orm:"column(sm_readids);size(255)"`
}

func (t *SunStoreMsg) TableName() string {
	return "shop_store_msg"
}

func init() {
	orm.RegisterModel(new(SunStoreMsg))
}

// AddSunStoreMsg insert a new SunStoreMsg into database and returns
// last inserted Id on success.
func AddSunStoreMsg(m *SunStoreMsg) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunStoreMsgById retrieves SunStoreMsg by Id. Returns error if
// Id doesn't exist
func GetSunStoreMsgById(id int) (v *SunStoreMsg, err error) {
	o := orm.NewOrm()
	v = &SunStoreMsg{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunStoreMsg retrieves all SunStoreMsg matches certain condition. Returns empty list if
// no records exist
func GetAllSunStoreMsg(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunStoreMsg))
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

	var l []SunStoreMsg
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

// UpdateSunStoreMsg updates SunStoreMsg by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunStoreMsgById(m *SunStoreMsg) (err error) {
	o := orm.NewOrm()
	v := SunStoreMsg{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunStoreMsg deletes SunStoreMsg by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunStoreMsg(id int) (err error) {
	o := orm.NewOrm()
	v := SunStoreMsg{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunStoreMsg{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
