package models

type SunTypeBrand struct {
	TypeId  uint `orm:"column(type_id)"`
	BrandId uint `orm:"column(brand_id)"`
}
