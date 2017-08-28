package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunEvaluateGoods struct {
	Id                  int     `orm:"column(geval_id);auto"`
	GevalOrderid        int     `orm:"column(geval_orderid)"`
	GevalOrderno        uint64  `orm:"column(geval_orderno)"`
	GevalOrdergoodsid   int     `orm:"column(geval_ordergoodsid)"`
	GevalGoodsid        int     `orm:"column(geval_goodsid)"`
	GevalGoodsname      string  `orm:"column(geval_goodsname);size(100)"`
	GevalGoodsprice     float64 `orm:"column(geval_goodsprice);null;digits(10);decimals(2)"`
	GevalGoodsimage     string  `orm:"column(geval_goodsimage);size(255);null"`
	GevalScores         int8    `orm:"column(geval_scores)"`
	GevalContent        string  `orm:"column(geval_content);size(255);null"`
	GevalIsanonymous    int8    `orm:"column(geval_isanonymous)"`
	GevalAddtime        int     `orm:"column(geval_addtime)"`
	GevalStoreid        int     `orm:"column(geval_storeid)"`
	GevalStorename      string  `orm:"column(geval_storename);size(100)"`
	GevalFrommemberid   int     `orm:"column(geval_frommemberid)"`
	GevalFrommembername string  `orm:"column(geval_frommembername);size(100)"`
	GevalState          int8    `orm:"column(geval_state)"`
	GevalRemark         string  `orm:"column(geval_remark);size(255);null"`
	GevalExplain        string  `orm:"column(geval_explain);size(255);null"`
	GevalImage          string  `orm:"column(geval_image);size(255);null"`
}

type evaluateGoodsAndMember struct {
	Id                  int     `orm:"column(geval_id);auto"`
	GevalOrderid        int     `orm:"column(geval_orderid)"`
	GevalOrderno        uint64  `orm:"column(geval_orderno)"`
	GevalOrdergoodsid   int     `orm:"column(geval_ordergoodsid)"`
	GevalGoodsid        int     `orm:"column(geval_goodsid)"`
	GevalGoodsname      string  `orm:"column(geval_goodsname);size(100)"`
	GevalGoodsprice     float64 `orm:"column(geval_goodsprice);null;digits(10);decimals(2)"`
	GevalGoodsimage     string  `orm:"column(geval_goodsimage);size(255);null"`
	GevalScores         int8    `orm:"column(geval_scores)"`
	GevalContent        string  `orm:"column(geval_content);size(255);null"`
	GevalIsanonymous    int8    `orm:"column(geval_isanonymous)"`
	GevalAddtime        int     `orm:"column(geval_addtime)"`
	GevalStoreid        int     `orm:"column(geval_storeid)"`
	GevalStorename      string  `orm:"column(geval_storename);size(100)"`
	GevalFrommemberid   int     `orm:"column(geval_frommemberid)"`
	GevalFrommembername string  `orm:"column(geval_frommembername);size(100)"`
	GevalState          int8    `orm:"column(geval_state)"`
	GevalRemark         string  `orm:"column(geval_remark);size(255);null"`
	GevalExplain        string  `orm:"column(geval_explain);size(255);null"`
	GevalImage          string  `orm:"column(geval_image);size(255);null"`
	MemberAvatar        string  `orm:"column(member_avatar);size(50);null"`
}

func (t *SunEvaluateGoods) TableName() string {
	return "shop_evaluate_goods"
}

func init() {
	orm.RegisterModel(new(SunEvaluateGoods))
}

// AddSunEvaluateGoods insert a new SunEvaluateGoods into database and returns
// last inserted Id on success.
func AddSunEvaluateGoods(m *SunEvaluateGoods) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunEvaluateGoodsById retrieves SunEvaluateGoods by Id. Returns error if
// Id doesn't exist
func GetSunEvaluateGoodsById(id int) (v *SunEvaluateGoods, err error) {
	o := orm.NewOrm()
	v = &SunEvaluateGoods{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunEvaluateGoods retrieves all SunEvaluateGoods matches certain condition. Returns empty list if
// no records exist
func GetAllSunEvaluateGoods(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunEvaluateGoods))
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

	var l []SunEvaluateGoods
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

// UpdateSunEvaluateGoods updates SunEvaluateGoods by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunEvaluateGoodsById(m *SunEvaluateGoods) (err error) {
	o := orm.NewOrm()
	v := SunEvaluateGoods{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunEvaluateGoods deletes SunEvaluateGoods by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunEvaluateGoods(id int) (err error) {
	o := orm.NewOrm()
	v := SunEvaluateGoods{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunEvaluateGoods{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return

}

func GetSunEvaluateGoods(goodsIds string, goodsCommonId string) (v []evaluateGoodsAndMember, err error) {
	o, q := GetQueryBuilder()

	if goodsIds != "" {
		evaluateSql := q.Select("*").From("shop_evaluate_goods as eg").InnerJoin("shop_member as m").On("eg.geval_frommemberid = m.member_id").Where("eg.geval_goodsid IN(?)").String()
		_, err = o.Raw(evaluateSql, goodsIds).QueryRows(&v)
	} else {
		evaluateSql := q.Select("*").From("shop_evaluate_goods as eg").InnerJoin("shop_member as m, shop_goods as g").On("eg.geval_goodsid = g.goods_id and eg.geval_frommemberid = m.member_id").Where("g.goods_commonid=?").String()
		_, err = o.Raw(evaluateSql, goodsCommonId).QueryRows(&v)
	}
	return
}