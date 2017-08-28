package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"strings"
	"github.com/yakun0622/shop/tools"
)

func Orm() orm.Ormer {
	return orm.NewOrm()
}

//SaveByID 根据id更新数据
func SaveByID(model interface{}, fields ...string) (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(model, fields...); err != nil {
		return err
	}
	return
}

//Save 新增数据
func Save(m interface{}) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//Remove 删除数据
func Remove(m interface{}) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(m)
	return
}

/**
 * 根据字段值删除
 */
func RemoveBy(table interface{}, query ...interface{}) bool {
	if len(query) < 2 {
		return false
	}

	o := orm.NewOrm()
	q := o.QueryTable(table)

	for i, key := range query {
		if i % 2 == 0 {
			q = q.Filter(key.(string), query[i + 1])
		}
	}

	if _, err := q.Delete(); err != nil {
		return false
	}
	return true
}

func Display(value ...interface{}) {
	utils.Display(value...)
}

//RegisterModel 模型注册
func RegisterModel(models ...interface{}) {
	orm.RegisterModel(models...)
}

//获取QueryBuilder
func GetQueryBuilder() (orm.Ormer, orm.QueryBuilder) {
	o := orm.NewOrm()

	if qb, error := orm.NewQueryBuilder("mysql"); error != nil {
		Display("QueryBuilder", "Get QueryBuilder Fail!!")
		return o, nil
	} else {
		return o, qb
	}
}

func QueryBuilder() orm.QueryBuilder {
	if qb, error := orm.NewQueryBuilder("mysql"); error != nil {
		Display("QueryBuilder", "Get QueryBuilder Fail!!")
		return nil
	} else {
		return qb
	}
}

//判断是否存在
func IsExist(model interface{}, key string, value interface{}) bool {
	o := orm.NewOrm()
	return o.QueryTable(model).Filter(key, value).Exist()
}

//GetSortFields 根据sortby及order，输出最终排序
func GetSortFields(sortby []string, order []string) ([]string, error) {
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
	return sortFields, nil
}

//数据count
func GetCount(query map[string]string, table interface{}) (totalCount int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	if totalCount, err = qs.Count(); err == nil {
		return totalCount, nil
	}
	return 0, err
}

func buildFieldStr(fields []string) string {
	fildsStr := ""
	if fields == nil {
		fildsStr = "*"
	} else {
		for _, fileld := range fields {
			fildsStr += fileld + ","
		}
		tools.Substr(fildsStr, 0, len(fildsStr) - 1)
	}
	return fildsStr
}