package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunStoreSnsComment struct {
	Id                int    `orm:"column(scomm_id);auto"`
	StraceId          int    `orm:"column(strace_id)"`
	ScommContent      string `orm:"column(scomm_content);size(150);null"`
	ScommMemberid     int    `orm:"column(scomm_memberid);null"`
	ScommMembername   string `orm:"column(scomm_membername);size(45);null"`
	ScommMemberavatar string `orm:"column(scomm_memberavatar);size(50);null"`
	ScommTime         string `orm:"column(scomm_time);size(11);null"`
	ScommState        int8   `orm:"column(scomm_state)"`
}

func (t *SunStoreSnsComment) TableName() string {
	return "sun_store_sns_comment"
}

func init() {
	orm.RegisterModel(new(SunStoreSnsComment))
}

// AddSunStoreSnsComment insert a new SunStoreSnsComment into database and returns
// last inserted Id on success.
func AddSunStoreSnsComment(m *SunStoreSnsComment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunStoreSnsCommentById retrieves SunStoreSnsComment by Id. Returns error if
// Id doesn't exist
func GetSunStoreSnsCommentById(id int) (v *SunStoreSnsComment, err error) {
	o := orm.NewOrm()
	v = &SunStoreSnsComment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunStoreSnsComment retrieves all SunStoreSnsComment matches certain condition. Returns empty list if
// no records exist
func GetAllSunStoreSnsComment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunStoreSnsComment))
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

	var l []SunStoreSnsComment
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

// UpdateSunStoreSnsComment updates SunStoreSnsComment by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunStoreSnsCommentById(m *SunStoreSnsComment) (err error) {
	o := orm.NewOrm()
	v := SunStoreSnsComment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunStoreSnsComment deletes SunStoreSnsComment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunStoreSnsComment(id int) (err error) {
	o := orm.NewOrm()
	v := SunStoreSnsComment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunStoreSnsComment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
