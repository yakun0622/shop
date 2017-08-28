package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SunGoods struct {
	Id                   int     `orm:"column(goods_id);auto"`
	GoodsCommonid        uint    `orm:"column(goods_commonid)"`
	GoodsName            string  `orm:"column(goods_name);size(50)"`
	GoodsJingle          string  `orm:"column(goods_jingle);size(150)"`
	StoreId              uint    `orm:"column(store_id)"`
	StoreName            string  `orm:"column(store_name);size(50)"`
	GcId                 uint    `orm:"column(gc_id)"`
	GcId1                uint    `orm:"column(gc_id_1)"`
	GcId2                uint    `orm:"column(gc_id_2)"`
	GcId3                uint    `orm:"column(gc_id_3)"`
	BrandId              uint    `orm:"column(brand_id)"`
	GoodsPrice           float64 `orm:"column(goods_price);digits(10);decimals(2)"`
	GoodsPromotionPrice  float64 `orm:"column(goods_promotion_price);digits(10);decimals(2)"`
	GoodsPromotionType   uint8   `orm:"column(goods_promotion_type)"`
	GoodsMarketprice     float64 `orm:"column(goods_marketprice);digits(10);decimals(2)"`
	GoodsSerial          string  `orm:"column(goods_serial);size(50)"`
	GoodsStorageAlarm    uint8   `orm:"column(goods_storage_alarm)"`
	GoodsClick           uint    `orm:"column(goods_click)"`
	GoodsSalenum         uint    `orm:"column(goods_salenum)"`
	GoodsCollect         uint    `orm:"column(goods_collect)"`
	GoodsSpec            string  `orm:"column(goods_spec)"`
	GoodsStorage         uint    `orm:"column(goods_storage)"`
	GoodsImage           string  `orm:"column(goods_image);size(100)"`
	GoodsImages          string  `orm:"column(goods_images);size(500)"`
	Specs          		string  `orm:"column(specs);size(500)"`
	GoodsState           uint8   `orm:"column(goods_state)"`
	GoodsVerify          uint8   `orm:"column(goods_verify)"`
	GoodsAddtime         uint    `orm:"column(goods_addtime)"`
	GoodsEdittime        uint    `orm:"column(goods_edittime)"`
	Areaid1              uint    `orm:"column(areaid_1)"`
	Areaid2              uint    `orm:"column(areaid_2)"`
	TransportId          uint32  `orm:"column(transport_id)"`
	GoodsFreight         float64 `orm:"column(goods_freight);digits(10);decimals(2)"`
	GoodsVat             uint8   `orm:"column(goods_vat)"`
	GoodsCommend         uint8   `orm:"column(goods_commend)"`
	GoodsStcids          string  `orm:"column(goods_stcids);size(255)"`
	EvaluationGoodStar   uint8   `orm:"column(evaluation_good_star)"`
	EvaluationCount      uint    `orm:"column(evaluation_count)"`
	IsVirtual            uint8   `orm:"column(is_virtual)"`
	VirtualIndate        uint    `orm:"column(virtual_indate)"`
	VirtualLimit         uint8   `orm:"column(virtual_limit)"`
	VirtualInvalidRefund uint8   `orm:"column(virtual_invalid_refund)"`
	IsFcode              int8    `orm:"column(is_fcode)"`
	IsAppoint            uint8   `orm:"column(is_appoint)"`
	IsPresell            uint8   `orm:"column(is_presell)"`
	HaveGift             uint8   `orm:"column(have_gift)"`
	IsOwnShop            uint8   `orm:"column(is_own_shop)"`
}

func (t *SunGoods) TableName() string {
	return "sun_goods"
}

func init() {
	orm.RegisterModel(new(SunGoods))
}

func GetGoodsCommonAndGoodses(goodsId string, goodsCommonId string) (common SunGoodsCommon, goodses []SunGoods, err error) {
	o, q := GetQueryBuilder()
	goodsSql := q.Select("*").From("sun_goods").Where("goods_commonid=?").String()
	commonQuery := QueryBuilder()

	if goodsCommonId != "" {
		goodsCommonSql := commonQuery.Select("*").From("sun_goods_common").Where("goods_commonid=?").String()

		err = o.Raw(goodsCommonSql, goodsCommonId).QueryRow(&common)
		if err != nil {
			Display("goodsCommonId--goodscommon", err, "sql", goodsCommonSql)
			return
		}
		_, err = o.Raw(goodsSql, goodsCommonId).QueryRows(&goodses)
		return
	}

	goodsCommonSql := commonQuery.Select("*").From("sun_goods as g").
		InnerJoin("sun_goods_common as gc").On("g.goods_commonid = gc.goods_commonid").
		Where("g.goods_id=?").String()

	err = o.Raw(goodsCommonSql, goodsId).QueryRow(&common)
	if err != nil {
		Display("goodsId--goodscommon", err, "sql", goodsCommonSql)
		return
	}
	commonId := common.Id
	_, err = o.Raw(goodsSql, commonId).QueryRows(&goodses)
	return
}

// AddSunGoods insert a new SunGoods into database and returns
// last inserted Id on success.
func AddSunGoods(m *SunGoods) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunGoodsById retrieves SunGoods by Id. Returns error if
// Id doesn't exist
func GetSunGoodsById(id int) (v *SunGoods, err error) {
	o := orm.NewOrm()
	v = &SunGoods{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunGoods retrieves all SunGoods matches certain condition. Returns empty list if
// no records exist
func GetAllSunGoods(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoods))
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

	var l []SunGoods
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

// UpdateSunGoods updates SunGoods by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunGoodsById(m *SunGoods) (err error) {
	o := orm.NewOrm()
	v := SunGoods{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunGoods deletes SunGoods by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunGoods(id int) (err error) {
	o := orm.NewOrm()
	v := SunGoods{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunGoods{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//GetAllSunGoodsByCommonID 根据CommonID批量查询商品
func GetAllSunGoodsByCommonID(ids []int, fields []string) (ml []SunGoods, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoods))
	qs = qs.Filter("goods_commonid__in", ids)
	if _, err := qs.All(&ml); err == nil {
		return ml, nil
	}

	return nil, err
}

func GetGoodsByCommon(id int) (ml []SunGoods, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoods))
	qs = qs.Filter("goods_commonid", id)
	if _, err := qs.All(&ml); err == nil {
		return ml, nil
	}
	return nil, err
}

//GetAllSunGoodsByIDs 根据ID批量查询商品
func GetAllSunGoodsByIDs(ids []int, fields []string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoods))
	qs = qs.Filter("goods_id__in", ids)
	var l []SunGoods
	if _, err := qs.All(&l); err == nil {
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

//GetAllRealtedGoodsByIDs 根据产品ID，获取同GoodsCommonID的产品
func GetAllRealtedGoodsByIDs(ids []int) (ml []interface{}, err error) {
	if len(ids) <= 0 {
		return nil, errors.New("no ids")
	}
	o := orm.NewOrm()
	var l []SunGoods
	sql := "SELECT * FROM sun_goods WHERE goods_commonid IN (SELECT DISTINCT goods_commonid FROM sun_goods WHERE goods_id IN ("
	for index := range ids {
		if index != 0 {
			sql += ","
		}
		sql += "?"
	}
	sql += "))"
	if _, err := o.Raw(sql, ids).QueryRows(&l); err == nil {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}

//根据GoodsID更新库存
func UpdateGoodsStorageByID(id int, num uint) (SunGoods, error) {
	o := orm.NewOrm()
	var v SunGoods
	sql := "UPDATE `sun_goods` SET `goods_storage`=`goods_storage`+? WHERE `goods_id`=?"
	res, err := o.Raw(sql, num, id).Exec()
	if err == nil{
		updaterow, _ := res.RowsAffected()
		if updaterow == 0 {
			beego.Info("无数据更新")
		} else {
			beego.Info("库存更新  商品ID:", id, " 数量变化：", num)
		}
	} else {
		return v, err
	}
	return v,nil
}
