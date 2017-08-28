package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunInvoice struct {
	Id             int    `orm:"column(inv_id);auto"`
	MemberId       uint   `orm:"column(member_id)"`
	InvState       string `orm:"column(inv_state);null"`
	InvTitle       string `orm:"column(inv_title);size(50);null"`
	InvContent     string `orm:"column(inv_content);size(10);null"`
	InvCompany     string `orm:"column(inv_company);size(50);null"`
	InvCode        string `orm:"column(inv_code);size(50);null"`
	InvRegAddr     string `orm:"column(inv_reg_addr);size(50);null"`
	InvRegPhone    string `orm:"column(inv_reg_phone);size(30);null"`
	InvRegBname    string `orm:"column(inv_reg_bname);size(30);null"`
	InvRegBaccount string `orm:"column(inv_reg_baccount);size(30);null"`
	InvRecName     string `orm:"column(inv_rec_name);size(20);null"`
	InvRecMobphone string `orm:"column(inv_rec_mobphone);size(15);null"`
	InvRecProvince string `orm:"column(inv_rec_province);size(30);null"`
	InvGotoAddr    string `orm:"column(inv_goto_addr);size(50);null"`
}

func (t *SunInvoice) TableName() string {
	return "shop_invoice"
}

func init() {
	orm.RegisterModel(new(SunInvoice))
}

// AddSunInvoice insert a new SunInvoice into database and returns
// last inserted Id on success.
func AddSunInvoice(m *SunInvoice) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunInvoiceById retrieves SunInvoice by Id. Returns error if
// Id doesn't exist
func GetSunInvoiceById(id int) (v *SunInvoice, err error) {
	o := orm.NewOrm()
	v = &SunInvoice{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunInvoice retrieves all SunInvoice matches certain condition. Returns empty list if
// no records exist
func GetAllSunInvoice(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunInvoice))
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

	var l []SunInvoice
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

// UpdateSunInvoice updates SunInvoice by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunInvoiceById(m *SunInvoice) (err error) {
	o := orm.NewOrm()
	v := SunInvoice{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunInvoice deletes SunInvoice by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunInvoice(id int) (err error) {
	o := orm.NewOrm()
	v := SunInvoice{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunInvoice{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
