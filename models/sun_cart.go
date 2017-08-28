package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type SunCart struct {
	Id         int     `orm:"column(cart_id);auto" gorm:"column:cart_id;primary_key"`
	BuyerId    uint    `orm:"column(buyer_id)"`
	StoreId    uint    `orm:"column(store_id)"`
	StoreName  string  `orm:"column(store_name);size(50)"`
	GoodsId    uint    `orm:"column(goods_id)"`
	GoodsName  string  `orm:"column(goods_name);size(100)"`
	GoodsPrice float64 `orm:"column(goods_price);digits(10);decimals(2)"`
	GoodsNum   uint16  `orm:"column(goods_num)"`
	GoodsImage string  `orm:"column(goods_image);size(100)"`
	BlId       uint32  `orm:"column(bl_id)"`
	CartType   int8    `orm:"column(cart_type)"`
}

func (t *SunCart) TableName() string {
	return "sun_cart"
}

func init() {
	orm.RegisterModel(new(SunCart))
}

// AddSunCart insert a new SunCart into database and returns
// last inserted Id on success.
func AddSunCart(m *SunCart) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunCartById retrieves SunCart by Id. Returns error if
// Id doesn't exist
func GetSunCartById(id int) (v *SunCart, err error) {
	o := orm.NewOrm()
	v = &SunCart{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunCart retrieves all SunCart matches certain condition. Returns empty list if
// no records exist
func GetAllSunCart(userId uint) (ml []SunCart, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(SunCart)).Filter("buyer_id", userId).All(&ml)
	return
}

// UpdateSunCart updates SunCart by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunCartById(m *SunCart) (err error) {
	o := orm.NewOrm()
	v := SunCart{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunCart deletes SunCart by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunCart(id int) (err error) {
	o := orm.NewOrm()
	v := SunCart{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunCart{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func IsExistInCart(m *SunCart) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunCart))
	qs = qs.Filter("buyer_id", m.BuyerId).Filter("store_id", m.StoreId).Filter("goods_id", m.GoodsId).Filter("cart_type", m.CartType)
	err := qs.One(m)
	if err == orm.ErrNoRows {
		return false
	}
	return true
}
