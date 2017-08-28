package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SunAdmin struct {
	Id             int    `orm:"column(admin_id);auto"`
	AdminName      string `orm:"column(admin_name);size(20)" json:"admin_name"`
	AdminPassword  string `orm:"column(admin_password);size(32)" json:"-"`
	AdminLoginTime int    `orm:"column(admin_login_time)"`
	AdminLoginNum  int    `orm:"column(admin_login_num)"`
	AdminIsSuper   int8   `orm:"column(admin_is_super)"`
	AdminGid       int16  `orm:"column(admin_gid);null"`
}

func (t *SunAdmin) TableName() string {
	return "sun_admin"
}

func init() {
	orm.RegisterModel(new(SunAdmin))
}

// AddSunAdmin insert a new SunAdmin into database and returns
// last inserted Id on success.
func AddSunAdmin(m *SunAdmin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunAdminById retrieves SunAdmin by Id. Returns error if
// Id doesn't exist
func GetSunAdminById(id int) (v *SunAdmin, err error) {
	o := orm.NewOrm()
	v = &SunAdmin{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunAdmin retrieves all SunAdmin matches certain condition. Returns empty list if
// no records exist
func GetAllSunAdmin(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunAdmin))
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

	var l []SunAdmin
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

// UpdateSunAdmin updates SunAdmin by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunAdminById(m *SunAdmin) (err error) {
	o := orm.NewOrm()
	v := SunAdmin{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunAdmin deletes SunAdmin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunAdmin(id int) (err error) {
	o := orm.NewOrm()
	v := SunAdmin{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunAdmin{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func CheckAdminLogin(username, password string) (SunAdmin, error) {
	o := orm.NewOrm()
	//TODO:登录成功，添加登录日志
	v := SunAdmin{AdminName: username, AdminPassword: password}
	err := o.QueryTable(v).Filter("admin_name", username).Filter("admin_password", password).One(&v)
	if err == orm.ErrNoRows {
		return v, err
	}

	v.AdminLoginNum++
	v.AdminLoginTime = int(time.Now().Unix())
	o = orm.NewOrm()
	o.Update(&v, "admin_login_time", "admin_login_num")

	return v, nil
}
