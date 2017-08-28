package models

import (

	"fmt"

	"github.com/astaxie/beego/orm"
)

type SunAlbumPic struct {
	Id         uint    `orm:"column(apic_id);auto"`
	ApicName   string `orm:"column(apic_name);size(100)"`
	ApicTag    string `orm:"column(apic_tag);size(255)"`
	AclassId   uint   `orm:"column(aclass_id)"`
	ApicCover  string `orm:"column(apic_cover);size(255)"`
	ApicSize   uint   `orm:"column(apic_size)"`
	ApicSpec   string `orm:"column(apic_spec);size(100)"`
	StoreId    uint   `orm:"column(store_id)"`
	UploadTime uint   `orm:"column(upload_time)"`
}

func (t *SunAlbumPic) TableName() string {
	return "sun_album_pic"
}

func init() {
	orm.RegisterModel(new(SunAlbumPic))
}

// AddSunAlbumPic insert a new SunAlbumPic into database and returns
// last inserted Id on success.
func AddSunAlbumPic(m *SunAlbumPic) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunAlbumPicById retrieves SunAlbumPic by Id. Returns error if
// Id doesn't exist
func GetSunAlbumPicById(id int) (v *SunAlbumPic, err error) {
	o := orm.NewOrm()
	v = &SunAlbumPic{Id: uint(id)}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunAlbumPic retrieves all SunAlbumPic matches certain condition. Returns empty list if
// no records exist
func GetAllSunAlbumPic( storeId uint, page int ) (albumPics []SunAlbumPic, count int, err error){
	o, q := GetQueryBuilder()
	q.Select("*").From("sun_album_pic").
		Where("store_id=?").OrderBy("upload_time").Desc().
	Limit(40).Offset((page) * 40)

	countq := QueryBuilder()
	countq.Select("count(apic_id)").From("sun_album_pic").
		Where("store_id=?")
	_, err = o.Raw(q.String(), storeId).QueryRows(&albumPics)
	if err!= nil {
		return
	}
	err = o.Raw(countq.String(), storeId).QueryRow(&count)
	return
}

// UpdateSunAlbumPic updates SunAlbumPic by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunAlbumPicById(m *SunAlbumPic) (err error) {
	o := orm.NewOrm()
	v := SunAlbumPic{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunAlbumPic deletes SunAlbumPic by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunAlbumPic(id int) (err error) {
	o := orm.NewOrm()
	v := SunAlbumPic{Id: uint(id)}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunAlbumPic{Id: uint(id)}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
