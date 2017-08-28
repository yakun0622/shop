package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunFavoritesFolder struct {
	Id         int    `orm:"column(folder_id);auto"`
	FolderName string `orm:"column(folder_name);size(50)"`
	FolderType int    `orm:"column(folder_type)"`
	MemberId   uint   `orm:"column(member_id)"`
	CreatedAt  uint   `orm:"column(created_at)"`
}

func (t *SunFavoritesFolder) TableName() string {
	return "shop_favorites_folder"
}

func init() {
	orm.RegisterModel(new(SunFavoritesFolder))
}

// AddSunFavoritesFolder insert a new SunFavoritesFolder into database and returns
// last inserted Id on success.
func AddSunFavoritesFolder(m *SunFavoritesFolder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunFavoritesFolderById retrieves SunFavoritesFolder by Id. Returns error if
// Id doesn't exist
func GetSunFavoritesFolderById(id int) (v *SunFavoritesFolder, err error) {
	o := orm.NewOrm()
	v = &SunFavoritesFolder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunFavoritesFolder retrieves all SunFavoritesFolder matches certain condition. Returns empty list if
// no records exist
func GetAllSunFavoritesFolder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunFavoritesFolder))
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

	var l []SunFavoritesFolder
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

// UpdateSunFavoritesFolder updates SunFavoritesFolder by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunFavoritesFolderById(m *SunFavoritesFolder) (err error) {
	o := orm.NewOrm()
	v := SunFavoritesFolder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunFavoritesFolder deletes SunFavoritesFolder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunFavoritesFolder(id int) (err error) {
	o := orm.NewOrm()
	v := SunFavoritesFolder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunFavoritesFolder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//DeleteSunFavoritesFolderByUser 删除特定用户的特定收藏文件夹
func DeleteSunFavoritesFolderByUser(id int, memberID int) (err error) {
	o := orm.NewOrm()
	v := SunFavoritesFolder{Id: id, MemberId: uint(memberID)}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunFavoritesFolder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//RemoveFavoritesFolderWithAllFavorites
func RemoveFavoritesFolderWithAllFavorites(id int) (err error) {
	o := orm.NewOrm()
	folder := SunFavoritesFolder{Id: id}
	if err := o.Begin(); err == nil {
		if _, err := o.Delete(&folder); err != nil {
			o.Rollback()
			return err
		} else if _, err := o.QueryTable(SunFavorites{}).Filter("FolderId", uint(id)).Delete(); err != nil {
			o.Rollback()
			return err
		} else {
			o.Commit()
		}
	}
	return err
}

type favoritesFolder struct {
	FolderId   int    `orm:"column(folder_id);auto"`
	FolderName string `orm:"column(folder_name);size(50)"`
	FolderType int    `orm:"column(folder_type)"`
	MemberId   uint   `orm:"column(member_id)"`
	Id         int    `orm:"column(id)"`
	GoodsId    uint   `orm:"column(goods_id)"`
	CreatedAt  uint   `orm:"column(created_at)"`

	StoreId    uint    `orm:"column(store_id)"`
	StoreName  string  `orm:"column(store_name);size(50)"`
	GoodsName  string  `orm:"column(goods_name)"`
	GoodsImage string  `orm:"column(goods_image)"`
	GoodsPrice float64 `orm:"column(goods_price);digits(10);decimals(2)"`
	GoodsSpec  string  `orm:"column(goods_spec)"`
	GoodsNum   uint16  `orm:"column(goods_num)"`
	GcId       uint    `orm:"column(gc_id)"`
	FavTime    uint    `orm:"column(fav_time)"`
}

//GetAllSunFavoritesFolderByUserID 获取用户的所有收藏文件夹信息
func GetAllSunFavoritesFolderByUserID(memberID uint) (ml []favoritesFolder, err error) {
	o, q := GetQueryBuilder()

	sql := q.Select("ff.folder_id, ff.member_id, ff.folder_name,f.fav_time, ff.folder_type, ff.created_at, f.goods_num, f.goods_id, g.gc_id, f.id, f.store_id, g.store_name, g.goods_spec, g.goods_image, g.goods_name, g.goods_price").
		From("shop_favorites_folder as ff").
		LeftJoin("shop_favorites as f").
		On("ff.folder_id=f.folder_id").
		LeftJoin("shop_goods as g").
		On("f.goods_id = g.goods_id").
		Where("ff.member_id = ?").
		String()

	_, err = o.Raw(sql, memberID).QueryRows(&ml)
	return
}
