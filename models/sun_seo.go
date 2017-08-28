package models

type SunSeo struct {
	Id_RENAME   int32  `orm:"column(id)"`
	Title       string `orm:"column(title);size(255)"`
	Keywords    string `orm:"column(keywords);size(255)"`
	Description string `orm:"column(description)"`
	Type        string `orm:"column(type);size(20)"`
}
