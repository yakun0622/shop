package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunStoreGrade struct {
	Id               int     `orm:"column(sg_id);auto"`
	SgName           string  `orm:"column(sg_name);size(50);null"`
	SgGoodsLimit     uint32  `orm:"column(sg_goods_limit)"`
	SgAlbumLimit     uint32  `orm:"column(sg_album_limit)"`
	SgSpaceLimit     uint    `orm:"column(sg_space_limit)"`
	SgTemplateNumber uint8   `orm:"column(sg_template_number)"`
	SgTemplate       string  `orm:"column(sg_template);size(255);null"`
	SgPrice          float64 `orm:"column(sg_price);digits(10);decimals(2)"`
	SgDescription    string  `orm:"column(sg_description);null"`
	SgFunction       string  `orm:"column(sg_function);size(255);null"`
	SgSort           uint8   `orm:"column(sg_sort)"`
}

func (t *SunStoreGrade) TableName() string {
	return "shop_store_grade"
}

func init() {
	orm.RegisterModel(new(SunStoreGrade))
}

// AddSunStoreGrade insert a new SunStoreGrade into database and returns
// last inserted Id on success.
func AddSunStoreGrade(m *SunStoreGrade) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunStoreGradeById retrieves SunStoreGrade by Id. Returns error if
// Id doesn't exist
func GetSunStoreGradeById(id int) (v *SunStoreGrade, err error) {
	o := orm.NewOrm()
	v = &SunStoreGrade{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunStoreGrade retrieves all SunStoreGrade matches certain condition. Returns empty list if
// no records exist
func GetAllSunStoreGrade(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunStoreGrade))
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

	var l []SunStoreGrade
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

// UpdateSunStoreGrade updates SunStoreGrade by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunStoreGradeById(m *SunStoreGrade) (err error) {
	o := orm.NewOrm()
	v := SunStoreGrade{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunStoreGrade deletes SunStoreGrade by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunStoreGrade(id int) (err error) {
	o := orm.NewOrm()
	v := SunStoreGrade{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunStoreGrade{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
