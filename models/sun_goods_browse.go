package models

type SunGoodsBrowse struct {
	GoodsId    int `orm:"column(goods_id)"`
	MemberId   int `orm:"column(member_id)"`
	Browsetime int `orm:"column(browsetime)"`
	GcId       int `orm:"column(gc_id)"`
	GcId1      int `orm:"column(gc_id_1)"`
	GcId2      int `orm:"column(gc_id_2)"`
	GcId3      int `orm:"column(gc_id_3)"`
}
