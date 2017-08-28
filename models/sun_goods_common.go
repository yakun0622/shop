package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

const (
	STATE1   = 1
	STATE0   = 0  // 下架
	STATE10  = 10 // 违规
	VERIFY1  = 1  // 审核通过
	VERIFY0  = 0  // 审核失败
	VERIFY10 = 10 // 等待审核
	LOCK0    = 0  // 锁定
	LOCK1    = 1  // 未锁定
)

type SunGoodsCommon struct {
	Id                   int     `orm:"column(goods_commonid);auto"`
	GoodsName            string  `orm:"column(goods_name);size(50)"`
	GoodsJingle          string  `orm:"column(goods_jingle);size(150)"`
	GcId                 uint    `orm:"column(gc_id)"`
	GcId1                uint    `orm:"column(gc_id_1)"`
	GcId2                uint    `orm:"column(gc_id_2)"`
	GcId3                uint    `orm:"column(gc_id_3)"`
	GcName               string  `orm:"column(gc_name);size(200)"`
	StoreId              uint    `orm:"column(store_id)"`
	StoreName            string  `orm:"column(store_name);size(50)"`
	SpecName             string  `orm:"column(spec_name)"`
	SpecValue            string  `orm:"column(spec_value)"`
	BrandId              uint    `orm:"column(brand_id)"`
	BrandName            string  `orm:"column(brand_name);size(100)"`
	TypeId               uint    `orm:"column(type_id)"`
	GoodsImage           string  `orm:"column(goods_image);size(100)"`
	GoodsImages          string  `orm:"column(goods_images);size(500)"`
	GoodsAttr            string  `orm:"column(goods_attr)"`
	GoodsBody            string  `orm:"column(goods_body)"`
	MobileBody           string  `orm:"column(mobile_body)"`
	Specs          		string  `orm:"column(specs)"`
	GoodsState           uint8   `orm:"column(goods_state)"`
	GoodsStateremark     string  `orm:"column(goods_stateremark);size(255);null"`
	GoodsVerify          uint8   `orm:"column(goods_verify)"`
	GoodsVerifyremark    string  `orm:"column(goods_verifyremark);size(255);null"`
	GoodsLock            uint8   `orm:"column(goods_lock)"`
	GoodsAddtime         uint    `orm:"column(goods_addtime)"`
	GoodsSelltime        uint    `orm:"column(goods_selltime)"`
	GoodsSpecname        string  `orm:"column(goods_specname)"`
	GoodsPrice           float64 `orm:"column(goods_price);digits(10);decimals(2)"`
	GoodsDiscount        float32 `orm:"column(goods_discount)"`
	GoodsSerial          string  `orm:"column(goods_serial);size(50)"`
	GoodsStorageAlarm    uint8   `orm:"column(goods_storage_alarm)"`
	TransportId          uint32  `orm:"column(transport_id)"`
	TransportTitle       string  `orm:"column(transport_title);size(60)"`
	GoodsCommend         uint8   `orm:"column(goods_commend)"`
	GoodsFreight         float64 `orm:"column(goods_freight);digits(10);decimals(2)"`
	GoodsVat             uint8   `orm:"column(goods_vat)"`
	Areaid1              uint    `orm:"column(areaid_1)"`
	Areaid2              uint    `orm:"column(areaid_2)"`
	GoodsStcids          string  `orm:"column(goods_stcids);size(255)"`
	PlateidTop           uint    `orm:"column(plateid_top);null"`
	PlateidBottom        uint    `orm:"column(plateid_bottom);null"`
	IsVirtual            uint8   `orm:"column(is_virtual)"`
	VirtualIndate        uint    `orm:"column(virtual_indate);null"`
	VirtualLimit         uint8   `orm:"column(virtual_limit);null"`
	VirtualInvalidRefund uint8   `orm:"column(virtual_invalid_refund)"`
	IsFcode              uint8   `orm:"column(is_fcode)"`
	IsAppoint            uint8   `orm:"column(is_appoint)"`
	AppointSatedate      uint    `orm:"column(appoint_satedate)"`
	IsPresell            uint8   `orm:"column(is_presell)"`
	PresellDeliverdate   uint    `orm:"column(presell_deliverdate)"`
	IsOwnShop            uint8   `orm:"column(is_own_shop)"`

	Goods []SunGoods `orm:"-"`
}

func (t *SunGoodsCommon) TableName() string {
	return "shop_goods_common"
}

func init() {
	orm.RegisterModel(new(SunGoodsCommon))
}

// AddSunGoodsCommon insert a new SunGoodsCommon into database and returns
// last inserted Id on success.
func AddSunGoodsCommon(m *SunGoodsCommon) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunGoodsCommonById retrieves SunGoodsCommon by Id. Returns error if
// Id doesn't exist
func GetSunGoodsCommonById(id int) (v *SunGoodsCommon, err error) {
	o := orm.NewOrm()
	v = &SunGoodsCommon{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunGoodsCommon retrieves all SunGoodsCommon matches certain condition. Returns empty list if
// no records exist
func GetAllSunGoodsCommon(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsCommon))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		if k == "keywords" {
			keywordSlices := strings.Split(v, " ")
			cond := orm.NewCondition()
			for _, keyword := range keywordSlices {
				if len(keyword) > 0 {
					cond = cond.Or("goods_name__contains", keyword )
				}
			}
			qs = qs.SetCond(cond)
		}else {
			k = strings.Replace(k, ".", "__", -1)
			qs = qs.Filter(k, v)
		}

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
					return nil, 0,  errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SunGoodsCommon
	qs = qs.OrderBy(sortFields...)

	if count, err = qs.Count(); err != nil {
		return nil, 0, errors.New("Error: faild to get count")
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
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
		return ml, count, nil
	}
	return nil, 0, err
}

func SaveGoodsesAndCommon(goodsCommon *SunGoodsCommon, goodses []SunGoods) (err error) {
	o := Orm()
	o.Begin()

	commonId, err := o.Insert(goodsCommon)
	if err != nil {
		Display("goodscoom", err)
		o.Rollback()
		return
	}
	for _, goods := range goodses {
		goods.GoodsCommonid = uint(commonId)
		goods.StoreId = goodsCommon.StoreId
		_, err = o.Insert(&goods)
		if err != nil {
			Display("goods", err)
			o.Rollback()
			return
		}
	}

	o.Commit()
	return
}

//GetAllSunGoodsCommonCount return totalCount
func GetAllSunGoodsCommonCount(query map[string]string) (totalCount int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsCommon))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	//qs = qs.Filter("goods_state", 1)
	//qs = qs.Filter("goods_verify", 1)
	//qs = qs.Filter("goods_lock", 0)
	if totalCount, err = qs.Count(); err == nil {
		return totalCount, nil
	}
	return 0, err
}

// UpdateSunGoodsCommon updates SunGoodsCommon by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunGoodsCommonById(m *SunGoodsCommon) (err error) {
	o := orm.NewOrm()
	v := SunGoodsCommon{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunGoodsCommon deletes SunGoodsCommon by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunGoodsCommon(id int) (err error) {
	o := orm.NewOrm()
	v := SunGoodsCommon{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunGoodsCommon{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteGoodsCommon 可以根据置入属性进行删除
func DeleteGoodsCommon(goodsCommon *SunGoodsCommon) (err error) {
	o := orm.NewOrm()
	// ascertain id exists in the database
	if err = o.Read(&goodsCommon); err == nil {
		var num int64
		if num, err = o.Delete(goodsCommon); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//GetAllSunGoodsCommonByIDs 根据commonID批量查询商品
func GetAllSunGoodsCommonByIDs(ids []int, fields []string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsCommon))
	qs = qs.Filter("goods_commonid__in", ids)
	qs = qs.Filter("goods_state", 1)
	qs = qs.Filter("goods_verify", 1)
	qs = qs.Filter("goods_lock", 0)
	var l []SunGoodsCommon
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

//GetUnlockGoodsCommonByIDs 根据commonID批量查询商品
func GetUnlockGoodsCommonByIDs(ids []int) (ml []SunGoodsCommon, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsCommon))
	qs = qs.Filter("goods_commonid__in", ids)
	qs = qs.Filter("goods_lock", 0)
	if _, err := qs.All(&ml); err == nil {
		return ml, nil
	}

	return nil, err
}

// FilterSunGoodsCommon 根据条件筛选GoodsCommon
func FilterSunGoodsCommon(keyword string, query map[string]string, specValues map[string]string, brands []int, goodsPrice string,
	sortby []string, order []string, offset int64, limit int64) (ml []interface{}, err error) {
	cond := GetCond(keyword, query, specValues, brands, goodsPrice)
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsCommon)).SetCond(cond)
	//合规
	qs = qs.Filter("goods_state", 1)
	qs = qs.Filter("goods_verify", 1)
	qs = qs.Filter("goods_lock", 0)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	//排序
	sortFields, err := GetSortFields(sortby, order)
	if err != nil {
		return nil, err
	}

	if len(sortFields) > 0 {
		qs = qs.OrderBy(sortFields...)
	}

	var l []SunGoodsCommon
	if _, err := qs.Limit(limit, offset).All(&l); err == nil {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}

// FilterSunGoodsCommonAllCount 根据条件筛选计算总量
func FilterSunGoodsCommonAllCount(keyword string, query map[string]string, specValues map[string]string, brands []int, goodsPrice string) (int64, error) {
	cond := GetCond(keyword, query, specValues, brands, goodsPrice)
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGoodsCommon)).SetCond(cond)
	qs = qs.Filter("goods_state", 1)
	qs = qs.Filter("goods_verify", 1)
	qs = qs.Filter("goods_lock", 0)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	return qs.Count()
}

//GetCond 封装商品筛选条件处理
func GetCond(keyword string, query map[string]string, specValues map[string]string, brands []int, goodsPrice string) *orm.Condition {
	cond := orm.NewCondition()
	if keyword != "" {
		keywordSlices := strings.Split(keyword, " ")
		condKeyword := orm.NewCondition()
		for _, keyword := range keywordSlices {
			if len(keyword) > 0 {
				condKeyword = condKeyword.Or("goods_name__contains", keyword )
			}
		}
		cond = cond.AndCond(condKeyword)
	}



	//多规格筛选，599:32组/箱,598:3卷/组，，，规格值ID：规格值
	if len(specValues) > 0 {
		cond1 := orm.NewCondition()

		for key, value := range specValues {
			cond1 = cond1.Or("spec_value__contains", `"`+key+`":"`+value+`"`)
		}
		cond = cond.AndCond(cond1)
	}

	//多品牌筛选
	if len(brands) > 0 {
		cond = cond.And("brand_id__in", brands)
	}

	//价格筛选，50~100， ~100，100~
	if goodsPrice != "" {
		prices := strings.Split(goodsPrice, ",")
		condPrice := orm.NewCondition()
		for _, priceStr := range prices{
			price := strings.Split(priceStr, "~")
			//|出现位置
			splitStringIndex := strings.Index(goodsPrice, "~")

			//情况： ~100
			if splitStringIndex == 0 {
				lowPrice, _ := strconv.Atoi(price[0])
				condPrice = condPrice.Or("goods_price__lte", lowPrice)
			} else if splitStringIndex > 0 && len(price) == 2 {
				lowPrice, _ := strconv.Atoi(price[0])
				highPrice, _ := strconv.Atoi(price[1])
				tempCond := orm.NewCondition()
				tempCond = tempCond.And("goods_price__gte", lowPrice)
				tempCond = tempCond.And("goods_price__lte", highPrice)
				condPrice = condPrice.OrCond(tempCond)

			} else {
				highPrice, _ := strconv.Atoi(price[1])
				condPrice = condPrice.Or("goods_price__gte", highPrice)
			}
		}
		cond = cond.AndCond(condPrice)

	}
	return cond
}

/**
获取goodsCommon列表
statusType 1-online  2-offline  3-lockUp  4-waitVerify
*/
func GetGoodsCommonByStatus(statusType int, storeID uint, goodsName string, limit int64, offset int64) (ml []SunGoodsCommon, totalCount int64, err error) {
	orm := Orm()
	qs := orm.QueryTable(new(SunGoodsCommon))
	qs = qs.Filter("goods_lock", LOCK0).Filter("store_id", int(storeID)).OrderBy("-GoodsAddtime")
	fmt.Println("goodsname....", goodsName)
	if len(goodsName) > 0 {
		qs = qs.Filter("goods_name__contains", goodsName)
	}
	switch statusType {
	case 1:
		qs = qs.Filter("goods_state", STATE1).Filter("goods_verify", VERIFY1)
		break
	case 2:
		qs = qs.Filter("goods_state", STATE0).Filter("goods_verify", VERIFY1)
		break
	case 3:
		qs = qs.Filter("goods_state", STATE10).Filter("goods_verify", VERIFY1)
		break
	case 4:
		qs = qs.Exclude("goods_verify", VERIFY1)
		break
	default:
		break
	}
	qs.Limit(limit).Offset(offset).All(&ml)
	if len(ml) > 0 {
		for index, goodsCommon := range ml {
			goodsList, err := GetGoodsByCommon(goodsCommon.Id)
			if err != nil {
				return nil, 0, err
			}
			ml[index].Goods = goodsList
		}
	}
	totalCount, err = qs.Count()
	return ml, totalCount, err
}

/**
* 删除未锁定商品
*
 */
func DelGoodsNoLock(storeId uint, goodsCommonIds []int) (result bool, err error) {
	if len(goodsCommonIds) <= 0 {
		return false, errors.New("no goodsCommonIds")
	}
	if storeId <= 0 {
		return false, errors.New("no such store....")
	}

	// 删除商品表数据
	result = RemoveBy(new(SunGoods), "GoodsCommonid__in", goodsCommonIds)
	// 删除商品公共表数据
	result = RemoveBy(new(SunGoodsCommon), "Id__in", goodsCommonIds)
	// 删除商品图片表数据
	result = RemoveBy(new(SunGoodsImages), "GoodsCommonid__in", goodsCommonIds)

	return true, err

}

/**
* 商品下架
*
 */
func EditProducesOffline(storeId uint, goodsCommonIds []int) (result bool, err error) {
	if storeId <= 0 {
		return false, errors.New("no such store....")
	}
	return EditProducesState(goodsCommonIds, STATE0)
}

/**
* 商品上架
*
 */
func EditProducesOnline(storeId uint, goodsCommonIds []int) (result bool, err error) {
	if storeId <= 0 {
		return false, errors.New("no such store....")
	}
	return EditProducesState(goodsCommonIds, STATE1)
}

func EditProducesState(goodsCommonIds []int, state uint8) (result bool, err error) {
	if len(goodsCommonIds) <= 0 {
		return false, errors.New("no goodsCommonIds")
	}

	o := orm.NewOrm()
	for _, id := range goodsCommonIds {
		goodsCommon, err := GetSunGoodsCommonById(id)
		if err != nil {
			return false, err
		}
		goodsCommon.GoodsState = state
		var num int64
		if num, err = o.Update(goodsCommon); err == nil {
			fmt.Println("Number of records updated in database:", num)
		} else {
			return false, err
		}
	}

	return true, err

}
