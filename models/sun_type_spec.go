package models

type SunTypeSpec struct {
	TypeId uint `orm:"column(type_id)"`
	SpId   uint `orm:"column(sp_id)"`
}
