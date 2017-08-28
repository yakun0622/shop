package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunAddress struct {
	Id        int    `orm:"column(address_id);auto"`
	MemberId  uint32 `orm:"column(member_id)"`
	TrueName  string `orm:"column(true_name);size(50)"`
	AreaId    uint32 `orm:"column(area_id)"`
	CityId    int32  `orm:"column(city_id);null"`
	AreaInfo  string `orm:"column(area_info);size(255)"`
	Address   string `orm:"column(address);size(255)"`
	TelPhone  string `orm:"column(tel_phone);size(20);null"`
	MobPhone  string `orm:"column(mob_phone);size(15);null"`
	IsDefault int    `orm:"column(is_default)"`
	DlypId    int    `orm:"column(dlyp_id);null"`
}

func (t *SunAddress) TableName() string {
	return "shop_address"
}

func init() {
	orm.RegisterModel(new(SunAddress))
}

// AddSunAddress insert a new SunAddress into database and returns
// last inserted Id on success.
func AddSunAddress(m *SunAddress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunAddressById retrieves SunAddress by Id. Returns error if
// Id doesn't exist
func GetSunAddressById(id int) (v *SunAddress, err error) {
	o := orm.NewOrm()
	v = &SunAddress{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunAddress retrieves all SunAddress matches certain condition. Returns empty list if
// no records exist
func GetAllSunAddress(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunAddress))
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

	var l []SunAddress
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		fmt.Println("addressList: ", l)
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
	fmt.Println("err: ", err)
	return nil, err
}

// UpdateSunAddress updates SunAddress by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunAddressById(m *SunAddress) (err error) {
	o := orm.NewOrm()
	v := SunAddress{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunAddress deletes SunAddress by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunAddress(id int) (err error) {
	o := orm.NewOrm()
	v := SunAddress{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunAddress{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
