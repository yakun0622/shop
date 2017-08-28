package models

type SunTypeGc struct {
	TypeId uint `orm:"column(type_id)"`
	GcId   uint `orm:"column(gc_id)"`
}
