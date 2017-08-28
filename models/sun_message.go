package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunMessage struct {
	Id                int    `orm:"column(message_id);auto"`
	MessageParentId   int    `orm:"column(message_parent_id)"`
	FromMemberId      int    `orm:"column(from_member_id)"`
	ToMemberId        string `orm:"column(to_member_id);size(1000)"`
	MessageTitle      string `orm:"column(message_title);size(50);null"`
	MessageBody       string `orm:"column(message_body);size(255)"`
	MessageTime       string `orm:"column(message_time);size(10)"`
	MessageUpdateTime string `orm:"column(message_update_time);size(10);null"`
	MessageOpen       int8   `orm:"column(message_open)"`
	MessageState      int8   `orm:"column(message_state)"`
	MessageType       int8   `orm:"column(message_type)"`
	ReadMemberId      string `orm:"column(read_member_id);size(1000);null"`
	DelMemberId       string `orm:"column(del_member_id);size(1000);null"`
	MessageIsmore     int8   `orm:"column(message_ismore)"`
	FromMemberName    string `orm:"column(from_member_name);size(100);null"`
	ToMemberName      string `orm:"column(to_member_name);size(100);null"`
}

func (t *SunMessage) TableName() string {
	return "shop_message"
}

func init() {
	orm.RegisterModel(new(SunMessage))
}

// AddSunMessage insert a new SunMessage into database and returns
// last inserted Id on success.
func AddSunMessage(m *SunMessage) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunMessageById retrieves SunMessage by Id. Returns error if
// Id doesn't exist
func GetSunMessageById(id int) (v *SunMessage, err error) {
	o := orm.NewOrm()
	v = &SunMessage{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunMessage retrieves all SunMessage matches certain condition. Returns empty list if
// no records exist
func GetAllSunMessage(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunMessage))
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

	var l []SunMessage
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

// UpdateSunMessage updates SunMessage by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunMessageById(m *SunMessage) (err error) {
	o := orm.NewOrm()
	v := SunMessage{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunMessage deletes SunMessage by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunMessage(id int) (err error) {
	o := orm.NewOrm()
	v := SunMessage{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunMessage{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
