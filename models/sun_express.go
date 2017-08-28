package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/yakun0622/shop/redis"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type SunExpress struct {
	Id       int    `orm:"column(id)"`
	EName    string `orm:"column(e_name);size(50)"`
	EState   string `orm:"column(e_state)"`
	ECode    string `orm:"column(e_code);size(50)"`
	ELetter  string `orm:"column(e_letter);size(1)"`
	EOrder   string `orm:"column(e_order)"`
	EUrl     string `orm:"column(e_url);size(100)"`
	EZtState int    `orm:"column(e_zt_state);null"`
}

func (t *SunExpress) TableName() string {
	return "shop_express"
}

const ExpressTableName = "shop_express"

func init() {
	orm.RegisterModel(new(SunExpress))
}

// AddSunExpress insert a new SunExpress into database and returns
// last inserted Id on success.
func AddSunExpress(m *SunExpress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunExpressById retrieves SunExpress by Id. Returns error if
// Id doesn't exist
func GetSunExpressById(id int) (v *SunExpress, err error) {
	o := orm.NewOrm()
	v = &SunExpress{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunExpress retrieves all SunExpress matches certain condition. Returns empty list if
// no records exist
func GetAllSunExpress() (expressList []SunExpress, err error) {
	// 缓存中有数据则返回
	cacheExpressList := redis.Instance().Get("express")
	if cacheExpressList != nil {
		json.Unmarshal(cacheExpressList.([]byte), &expressList)
		beego.Info("从缓存获取快递列表.....")
		return expressList, nil
	}
	o, q := GetQueryBuilder()
	q = q.Select("*").From(ExpressTableName)
	_, err = o.Raw(q.String()).QueryRows(&expressList)
	expressListJson, err := json.Marshal(expressList)
	redis.Instance().Put("express", expressListJson, 3600*time.Minute)
	return expressList, err
}

// UpdateSunExpressById updates SunExpress by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunExpressById(m *SunExpress) (err error) {
	o := orm.NewOrm()
	v := SunExpress{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunExpress deletes SunExpress by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunExpress(id int) (err error) {
	o := orm.NewOrm()
	v := SunExpress{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunExpress{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
