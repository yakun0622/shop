package models

type SunGoodsAttrIndex struct {
	GoodsId       uint `orm:"column(goods_id)"`
	GoodsCommonid uint `orm:"column(goods_commonid)"`
	GcId          uint `orm:"column(gc_id)"`
	TypeId        uint `orm:"column(type_id)"`
	AttrId        uint `orm:"column(attr_id)"`
	AttrValueId   uint `orm:"column(attr_value_id)"`
}
