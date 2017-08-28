package models

type SunStoreMsgRead struct {
	SmId     int `orm:"column(sm_id)"`
	SellerId int `orm:"column(seller_id)"`
	ReadTime int `orm:"column(read_time)"`
}
