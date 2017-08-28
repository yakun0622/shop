package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunStore struct {
	Id                        int     `orm:"column(store_id);auto"`
	StoreName                 string  `orm:"column(store_name);size(50)"`
	GradeId                   int     `orm:"column(grade_id)"`
	MemberId                  int     `orm:"column(member_id)"`
	MemberName                string  `orm:"column(member_name);size(50)"`
	SellerName                string  `orm:"column(seller_name);size(50);null"`
	ScId                      int     `orm:"column(sc_id)"`
	StoreCompanyName          string  `orm:"column(store_company_name);size(50);null"`
	ProvinceId                uint32  `orm:"column(province_id)"`
	AreaInfo                  string  `orm:"column(area_info);size(100)"`
	StoreAddress              string  `orm:"column(store_address);size(100)"`
	StoreZip                  string  `orm:"column(store_zip);size(10)"`
	StoreState                int8    `orm:"column(store_state)"`
	StoreCloseInfo            string  `orm:"column(store_close_info);size(255);null"`
	StoreSort                 int     `orm:"column(store_sort)"`
	StoreTime                 string  `orm:"column(store_time);size(10)"`
	StoreEndTime              string  `orm:"column(store_end_time);size(10);null"`
	StoreLabel                string  `orm:"column(store_label);size(255);null"`
	StoreBanner               string  `orm:"column(store_banner);size(255);null"`
	StoreAvatar               string  `orm:"column(store_avatar);size(150);null"`
	StoreKeywords             string  `orm:"column(store_keywords);size(255)"`
	StoreDescription          string  `orm:"column(store_description);size(255)"`
	StoreQq                   string  `orm:"column(store_qq);size(50);null"`
	StoreWw                   string  `orm:"column(store_ww);size(50);null"`
	StorePhone                string  `orm:"column(store_phone);size(20);null"`
	StoreZy                   string  `orm:"column(store_zy);null"`
	StoreDomain               string  `orm:"column(store_domain);size(50);null"`
	StoreDomainTimes          uint8   `orm:"column(store_domain_times);null"`
	StoreRecommend            int8    `orm:"column(store_recommend)"`
	StoreTheme                string  `orm:"column(store_theme);size(50)"`
	StoreCredit               int     `orm:"column(store_credit)"`
	StoreDesccredit           float32 `orm:"column(store_desccredit)"`
	StoreServicecredit        float32 `orm:"column(store_servicecredit)"`
	StoreDeliverycredit       float32 `orm:"column(store_deliverycredit)"`
	StoreCollect              uint    `orm:"column(store_collect)"`
	StoreSlide                string  `orm:"column(store_slide);null"`
	StoreSlideUrl             string  `orm:"column(store_slide_url);null"`
	StoreStamp                string  `orm:"column(store_stamp);size(200);null"`
	StorePrintdesc            string  `orm:"column(store_printdesc);size(500);null"`
	StoreSales                uint    `orm:"column(store_sales)"`
	StorePresales             string  `orm:"column(store_presales);null"`
	StoreAftersales           string  `orm:"column(store_aftersales);null"`
	StoreWorkingtime          string  `orm:"column(store_workingtime);size(100);null"`
	StoreFreePrice            float64 `orm:"column(store_free_price);digits(10);decimals(2)" json:"StoreFreePrice,string"`
	StoreDecorationSwitch     uint    `orm:"column(store_decoration_switch)"`
	StoreDecorationOnly       uint8   `orm:"column(store_decoration_only)"`
	StoreDecorationImageCount uint    `orm:"column(store_decoration_image_count)"`
	LiveStoreName             string  `orm:"column(live_store_name);size(255);null"`
	LiveStoreAddress          string  `orm:"column(live_store_address);size(255);null"`
	LiveStoreTel              string  `orm:"column(live_store_tel);size(255);null"`
	LiveStoreBus              string  `orm:"column(live_store_bus);size(255);null"`
	IsOwnShop                 uint8   `orm:"column(is_own_shop)"`
	BindAllGc                 uint8   `orm:"column(bind_all_gc)"`
	StoreVrcodePrefix         string  `orm:"column(store_vrcode_prefix);size(3);null"`
	StoreBaozh                int8    `orm:"column(store_baozh);null"`
	StoreBaozhopen            int8    `orm:"column(store_baozhopen);null"`
	StoreBaozhrmb             string  `orm:"column(store_baozhrmb);size(10);null"`
	StoreQtian                int8    `orm:"column(store_qtian);null"`
	StoreZhping               int8    `orm:"column(store_zhping);null"`
	StoreErxiaoshi            int8    `orm:"column(store_erxiaoshi);null"`
	StoreTuihuo               int8    `orm:"column(store_tuihuo);null"`
	StoreShiyong              int8    `orm:"column(store_shiyong);null"`
	StoreShiti                int8    `orm:"column(store_shiti);null"`
	StoreXiaoxie              int8    `orm:"column(store_xiaoxie);null"`
	StoreHuodaofk             int8    `orm:"column(store_huodaofk);null"`
}

func (t *SunStore) TableName() string {
	return "sun_store"
}

func init() {
	orm.RegisterModel(new(SunStore))
}

// AddSunStore insert a new SunStore into database and returns
// last inserted Id on success.
func AddSunStore(m *SunStore) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunStoreById retrieves SunStore by Id. Returns error if
// Id doesn't exist
func GetSunStoreById(id int) (v *SunStore, err error) {
	o := orm.NewOrm()
	v = &SunStore{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunStore retrieves all SunStore matches certain condition. Returns empty list if
// no records exist
func GetAllSunStore(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunStore))
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

	var l []SunStore
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

// UpdateSunStore updates SunStore by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunStoreById(m *SunStore) (err error) {
	o := orm.NewOrm()
	v := SunStore{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunStore deletes SunStore by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunStore(id int) (err error) {
	o := orm.NewOrm()
	v := SunStore{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunStore{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//GetGoodsCommonCountByStatus 获取各类状态的商品数目
//1:出售中
//2:已发布待审核
//3:审核失败
//4:仓库中已审核(待上架商品)
//5:待回复咨询
//6:违规下架
func GetGoodsCommonCountByStatus(storeID uint, status int) (count int64, err error) {
	o := orm.NewOrm()
	var qs orm.QuerySeter

	// 待回复咨询
	if status == 5 {
		qs = o.QueryTable(new(SunGoodsCommon))
		// qs = qs.Filter("store_id", storeID).Filter("consult_reply", "")
	} else {
		qs = o.QueryTable(new(SunGoodsCommon))
		qs = qs.Filter("store_id", storeID)
	}

	switch status {
	case 1:
		//出售中
		qs = qs.Filter("goods_state", 1).Filter("goods_verify", 1)
	case 2:
		//已发布待审核
		qs = qs.Filter("goods_verify", 10)
	case 3:
		//审核失败
		qs = qs.Filter("goods_verify", 0)
	case 4:
		//仓库中已审核(待上架商品)
		qs = qs.Filter("goods_state", 0).Filter("goods_verify", 1)
	case 5:
		//待回复咨询
	case 6:
		//违规下架
		qs = qs.Filter("goods_state", 10).Filter("goods_verify", 1)
	}

	count, err = qs.Count()
	if err == nil {
		return count, nil
	}
	return 0, err
}

//GetOrderCountByStatus 获取各类状态的商品数目
//1:待付款
//2:待发货
//3:售前退款
//4:售后退款
//5:售前退货
//6:售后退货
//7:待确认订单
func GetOrderCountByStatus(storeID uint, status int) (count int64, err error) {
	o := orm.NewOrm()
	var qs orm.QuerySeter

	if status == 1 || status == 2 {
		qs = o.QueryTable(new(SunOrder))
		qs = qs.Filter("store_id", storeID)
	} else if status == 3 || status == 4 || status == 5 || status == 6 {
		//退款退货相关
		qs = o.QueryTable(new(SunRefundReturn))
		qs = qs.Filter("store_id", storeID)
	} else if status == 7 {
		//待确认订单
		qs = o.QueryTable(new(SunOrderBill))
	}

	switch status {
	case 1:
		//审批中/待付款
		qs = qs.Filter("order_state", 10)
	case 2:
		//待发货
		qs = qs.Filter("order_state", 30)
	case 3:
		//售前退款
		qs = qs.Filter("refund_type", 1).Filter("order_lock", 2).Filter("refund_state__lt", 3)
	case 4:
		//售后退款
		qs = qs.Filter("refund_type", 1).Filter("order_lock", 1).Filter("refund_state__lt", 3)
	case 5:
		//售前退货
		qs = qs.Filter("refund_type", 2).Filter("order_lock", 2).Filter("refund_state__lt", 3)
	case 6:
		//售后退货
		qs = qs.Filter("refund_type", 2).Filter("order_lock", 1).Filter("refund_state__lt", 3)
	case 7:
		//待确认订单
		qs = qs.Filter("ob_store_id", storeID).Filter("ob_state", 1)
	}

	count, err = qs.Count()
	if err == nil {
		return count, nil
	}
	return 0, err
}

//GetStoreByMemberID 根据账号ID获取商店ID
func GetStoreByMemberID(memberID uint) (SunStore, error) {
	o := orm.NewOrm()
	v := SunStore{}
	qs := o.QueryTable(new(SunStore))
	qs = qs.Filter("member_id", memberID)
	if err := qs.One(&v); err != nil {
		return v, err
	}
	return v, nil
}

//SetStoreFreeFreightByStroreID 根据商户ID更新免运费额度
func SetStoreFreeFreightByStroreID(storeID uint, price float64) (SunStore, error) {
	o := orm.NewOrm()
	v := SunStore{Id: int(storeID)}

	if o.Read(&v) == nil {
		v.StoreFreePrice = price
		if _, err := o.Update(&v); err == nil {
			return v, nil
		}
	}
	return v, nil
}
