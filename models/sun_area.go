package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/yakun0622/shop/redis"
	"strconv"
	"time"
	"github.com/astaxie/beego"
	"github.com/yakun0622/shop/tools"
)

type SunArea struct {
	Id           int    `orm:"column(area_id);auto"`
	AreaName     string `orm:"column(area_name);size(50)"`
	AreaParentId uint   `orm:"column(area_parent_id)"`
	AreaSort     uint8  `orm:"column(area_sort)"`
	AreaDeep     uint8  `orm:"column(area_deep)"`
	AreaRegion   string `orm:"column(area_region);size(3);null"`

	CacheData interface{} `orm:"-"`
}

const AreaTableName = "sun_area"

func (t *SunArea) TableName() string {
	return "sun_area"
}

func init() {
	orm.RegisterModel(new(SunArea))
}

// AddSunArea insert a new SunArea into database and returns
// last inserted Id on success.
func AddSunArea(m *SunArea) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunAreaById retrieves SunArea by Id. Returns error if
// Id doesn't exist
func GetSunAreaById(id int) (v *SunArea, err error) {
	o := orm.NewOrm()
	v = &SunArea{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunArea retrieves all SunArea matches certain condition. Returns empty list if
// no records exist
func GetAllSunArea(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunArea))
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

	var l []SunArea
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

// UpdateSunArea updates SunArea by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunAreaById(m *SunArea) (err error) {
	o := orm.NewOrm()
	v := SunArea{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunArea deletes SunArea by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunArea(id int) (err error) {
	o := orm.NewOrm()
	v := SunArea{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunArea{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetAllArea() (result map[string]map[string]interface{}, area_children_data map[string][]interface{}, area_region_data map[string][]int, err error) {
	return GetCache()
}

func GetCache() (result map[string]map[string]interface{}, area_children_data map[string][]interface{}, area_region_data map[string][]int, err error) {
	// 缓存中有数据则返回
	cache_area_data := redis.Instance().Get("area")
	cache_area_children := redis.Instance().Get("area_children")
	cache_area_region := redis.Instance().Get("area_region")
	if cache_area_data != nil && cache_area_children != nil && cache_area_region != nil {
		var result_area_data map[string]map[string]interface{}
		var result_area_children map[string][]interface{}
		var result_area_region map[string][]int
		json.Unmarshal(cache_area_data.([]byte), &result_area_data)
		json.Unmarshal(cache_area_children.([]byte), &result_area_children)
		json.Unmarshal(cache_area_region.([]byte), &result_area_region)
		beego.Info("从缓存获取地区数据.....")
		return result_area_data, result_area_children, result_area_region, nil
	}

	//查库
	var area_all []SunArea
	data := make(map[string]map[string]interface{})
	o, q := GetQueryBuilder()
	q = q.Select("*").From(AreaTableName)
	beego.Info("查询开始.....")
	_, err = o.Raw(q.String()).QueryRows(&area_all)
	beego.Info("查询结束.....")
	area_name := make(map[string]interface{})
	parent_ids := make(map[string]interface{})
	area_children := make(map[string][]interface{})
	area_region := make(map[string][]int)
	for _, area := range area_all {
		area_name[strconv.Itoa(area.Id)] = area.AreaName
		parent_ids[strconv.Itoa(area.Id)] = area.AreaParentId
		children_slice := area_children[strconv.Itoa(int(area.AreaParentId))]
		children_slice = append(children_slice, area.Id)
		area_children[strconv.Itoa(int(area.AreaParentId))] = children_slice
		if area.AreaDeep == 1 && len(area.AreaRegion) > 0 {
			region_slice := area_region[area.AreaRegion]
			if !tools.InIntSlice(area.Id, region_slice) {
				area_region[area.AreaRegion] = append(region_slice, area.Id)
			}

			//area_region[area.AreaRegion] = append(tools.Slice_unique(region_slice), area.Id)
			//fmt.Println(area_region)
		}
	}
	//area_region 去重
	data["name"] = area_name
	data["parent"] = parent_ids
	area_data_json, err := json.Marshal(data)
	area_children_json, err := json.Marshal(area_children)
	area_region_json, err := json.Marshal(area_region)
	redis.Instance().Put("area", area_data_json, 3600*time.Minute)
	redis.Instance().Put("area_children", area_children_json, 3600*time.Minute)
	redis.Instance().Put("area_region", area_region_json, 3600*time.Minute)
	return data, area_children, area_region, err
}
