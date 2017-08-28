package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TransportExtend struct {
	Id             int     `orm:"column(id);auto"`
	AreaId         string  `orm:"column(area_id);null"`
	TopAreaId      string  `orm:"column(top_area_id);null"`
	AreaName       string  `orm:"column(area_name);null"`
	Snum           uint32  `orm:"column(snum);null"`
	Sprice         float64 `orm:"column(sprice);null;digits(10);decimals(2)"`
	Xnum           uint32  `orm:"column(xnum);null"`
	Xprice         float64 `orm:"column(xprice);null;digits(10);decimals(2)"`
	IsDefault      int     `orm:"column(is_default);null"`
	TransportId    int64   `orm:"column(transport_id)"`
	TransportTitle string  `orm:"column(transport_title);size(60);null"`
	DefalutPrice   float64 `orm:"column(defalut_price);null;digits(10);decimals(2)"`
	FreeLine       float64 `orm:"column(free_line);null;digits(10);decimals(2)"`
}

const TransportExtendTableName = "shop_transport_extend"

func (t *TransportExtend) TableName() string {
	return "shop_transport_extend"
}

func init() {
	orm.RegisterModel(new(TransportExtend))
}

// AddSunTransportExtend insert a new TransportExtend into database and returns
// last inserted Id on success.
func AddTransportExtend(m *TransportExtend) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunTransportExtendById retrieves TransportExtend by Id. Returns error if
// Id doesn't exist
func GetSunTransportExtendById(id int) (v *TransportExtend, err error) {
	o := orm.NewOrm()
	v = &TransportExtend{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunTransportExtend retrieves all TransportExtend matches certain condition. Returns empty list if
// no records exist
func GetAllSunTransportExtend(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TransportExtend))
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

	var l []TransportExtend
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

// UpdateSunTransportExtend updates TransportExtend by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunTransportExtendById(m *TransportExtend) (err error) {
	o := orm.NewOrm()
	v := TransportExtend{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunTransportExtend deletes TransportExtend by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunTransportExtend(id int) (err error) {
	v := TransportExtend{Id: id}
	return Remove(&v)
}

func DeleteTransportExtendByTransId(id int64) (err error) {
	orm, qb := GetQueryBuilder()
	qb = qb.Delete().From(TransportExtendTableName).Where("transport_id = " + strconv.Itoa(int(id)))
	_, err = orm.Raw(qb.String()).Exec()
	return err

}


func GetByTransportExtTransportId(transportId int64) (list []TransportExtend, err error) {
	orm, qb := GetQueryBuilder()
	qb = qb.Select("*").From(TransportExtendTableName).Where("transport_id = " + strconv.Itoa(int(transportId)))
	_, err = orm.Raw(qb.String()).QueryRows(&list)
	return list, err
}
