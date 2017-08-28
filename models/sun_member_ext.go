package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"strconv"
)

type SunMemberExt struct {
	Id         int        `orm:"column(member_ext_id);auto"`
	MemberUser *SunMember `orm:"column(member_userId);rel(one)"`
	// MemberUserId  uint64  `orm:"column(member_userId);"`
	MemberGroupId uint64  `orm:"column(member_groupId)"`
	MemberType    string  `orm:"column(member_type);size(64)"`
	MemberStatus  uint8   `orm:"column(member_status)"`
	MemberLevels  uint8   `orm:"column(member_levels)"`
	ApproveLevels uint8   `orm:"column(approve_levels);null"`
	ApproveOrder  float64 `orm:"column(approve_order);digits(10);decimals(2)"`
	ApproveMonth  float64 `orm:"column(approve_month);digits(10);decimals(2)"`
	ApproveMonths float64 `orm:"column(approve_months);digits(10);decimals(2)"`
	ApproveYear   float64 `orm:"column(approve_year);digits(10);decimals(2)"`
	ApproveYears  float64 `orm:"column(approve_years);digits(10);decimals(2)"`
	MemberCtime   string  `orm:"column(member_ctime);size(45);null"`
}

func (t *SunMemberExt) TableName() string {
	return "sun_member_ext"
}

func init() {
	orm.RegisterModel(new(SunMemberExt))
}

// AddSunMemberExt insert a new SunMemberExt into database and returns
// last inserted Id on success.
func AddSunMemberExt(m *SunMemberExt) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunMemberExtById retrieves SunMemberExt by Id. Returns error if
// Id doesn't exist
func GetSunMemberExtById(id int) (v *SunMemberExt, err error) {

	o := orm.NewOrm()
	v = &SunMemberExt{Id: id}
	if err = o.Read(v); err == nil {
		if v.MemberUser != nil {
			o.Read(v.MemberUser)
		}
		return v, nil
	}
	return nil, err

}

// GetAllSunMemberExt retrieves all SunMemberExt matches certain condition. Returns empty list if
// no records exist
func GetAllSunMemberExt(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()

	qs := o.QueryTable(new(SunMemberExt))

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
	var l []SunMemberExt
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).RelatedSel().All(&l, fields...); err == nil {
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
		//MemberExt := &SunMemberExt{}
		//qs.Filter("Id", 1).RelatedSel().One(MemberExt)

		return ml, nil
	}
	return nil, err
}

// UpdateSunMemberExt updates SunMemberExt by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunMemberExtById(m *SunMemberExt) (err error) {
	o := orm.NewOrm()
	v := SunMemberExt{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunMemberExt deletes SunMemberExt by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunMemberExt(id int) (err error) {
	o := orm.NewOrm()
	v := SunMemberExt{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunMemberExt{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetAllSunMemberExt retrieves all SunMemberExt matches certain condition. Returns empty list if
// no records exist
func GetAllMemberTree(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()

	qs := o.QueryTable(new(SunMemberExt))

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		if k == "MemberGroupId" {
			memberGroupId, _ := strconv.Atoi(v)
			groupChildList, _ := GetGroupChildList(memberGroupId)
			groupChildIds := buildChildIds(groupChildList)
			qs = qs.Filter("MemberGroupId__in", append(groupChildIds, v))
			fmt.Println(qs)

		} else {
			k = strings.Replace(k, ".", "__", -1)
			qs = qs.Filter(k, v)
		}

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

	var l []SunMemberExt
	qs = qs.OrderBy(sortFields...)
	fmt.Println(qs)
	if _, err := qs.Limit(limit, offset).RelatedSel().All(&l, fields...); err == nil {
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
		//MemberExt := &SunMemberExt{}
		//qs.Filter("Id", 1).RelatedSel().One(MemberExt)

		return ml, nil
	}
	return nil, err
}
