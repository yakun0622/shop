package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunDaddress struct {
	Id         int    `orm:"column(address_id);auto"`
	StoreId    uint32 `orm:"column(store_id)"`
	SellerName string `orm:"column(seller_name);size(50)"`
	//AreaId     *SunArea `orm:"column(area_id);rel(one)"`
	//CityId     *SunArea `orm:"column(city_id);rel(one)"`
	ProvinceId int32  `orm:"column(province_id);"`
	CityId     int32  `orm:"column(city_id);null"`
	AreaId     uint32 `orm:"column(area_id)"`

	AreaInfo  string `orm:"column(area_info);size(100);null"`
	Address   string `orm:"column(address);size(100)"`
	Telphone  string `orm:"column(telphone);size(40);null"`
	Company   string `orm:"column(company);size(50)"`
	IsDefault int `orm:"column(is_default)"`
}

func (t *SunDaddress) TableName() string {
	return "sun_daddress"
}

func init() {
	orm.RegisterModel(new(SunDaddress))
}

// AddSunDaddress insert a new SunDaddress into database and returns
// last inserted Id on success.
func AddSunDaddress(m *SunDaddress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunDaddressById retrieves SunDaddress by Id. Returns error if
// Id doesn't exist
func GetSunDaddressById(id int) (v *SunDaddress, err error) {
	o := orm.NewOrm()
	v = &SunDaddress{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunDaddress retrieves all SunDaddress matches certain condition. Returns empty list if
// no records exist
func GetAllSunDaddress(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*SunDaddress, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunDaddress))
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

	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&ml, fields...); err == nil {
		return ml, nil
	}
	return nil, err
}

// UpdateSunDaddress updates SunDaddress by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunDaddressById(m *SunDaddress) (err error) {
	o := orm.NewOrm()
	v := SunDaddress{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunDaddress deletes SunDaddress by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunDaddress(id int) (err error) {
	o := orm.NewOrm()
	v := SunDaddress{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunDaddress{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteSunDaddressByStore 根据店铺ID及ID删除发货地址，直接跳过店铺ID与发货地址的存在校验
func DeleteSunDaddressByStore(id int, stroreID uint) (result bool, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunDaddress))
	if num, err := qs.Filter("Id", id).Filter("StoreId", stroreID).Delete(); err == nil {
		return num > 0, nil
	}
	return false, err
}

// SetDefaultDAddress 设置默认收货地址
func SetDefaultDAddress(id int, stroreID uint) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunDaddress))
	if count, err := qs.Filter("Id", id).Filter("StoreId", stroreID).Count(); err != nil || count <= 0 {
		return false
	}
	o.Begin()
	_, err1 := qs.Filter("Id", id).Update(orm.Params{"IsDefault": 1})
	_, err2 := qs.Filter("StoreId", stroreID).Exclude("Id", id).Update(orm.Params{"IsDefault": 2})

	if err1 != nil || err2 != nil {
		o.Rollback()
		return false
	}
	o.Commit()
	return true

}
