package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunFavorites struct {
	Id       int    `orm:"column(id);auto"`
	MemberId uint   `orm:"column(member_id)"`
	StoreId  uint   `orm:"column(store_id)"`
	GoodsId  uint   `orm:"column(goods_id)"`
	GoodsNum uint16 `orm:"column(goods_num)"`
	FavTime  uint   `orm:"column(fav_time)"`
	FolderId uint   `orm:"column(folder_id)"`
}

func (t *SunFavorites) TableName() string {
	return "sun_favorites"
}

func init() {
	orm.RegisterModel(new(SunFavorites))
}

// AddSunFavorites insert a new SunFavorites into database and returns
// last inserted Id on success.
func AddSunFavorites(m *SunFavorites) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunFavoritesById retrieves SunFavorites by Id. Returns error if
// Id doesn't exist
func GetSunFavoritesById(id int) (v *SunFavorites, err error) {
	o := orm.NewOrm()
	v = &SunFavorites{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunFavorites retrieves all SunFavorites matches certain condition. Returns empty list if
// no records exist
func GetAllSunFavorites(query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunFavorites))
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

	var l []SunFavorites
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

func GetAllSunFavoritesBuyFolderId(folderId int, userId uint) (favs []SunFavorites, err error) {
	o, q := GetQueryBuilder()
	sql := q.Select("*").
		From("sun_favorites as f").
		InnerJoin("sun_goods as g").
		On("f.goods_id = g.goods_id").
		Where("folder_id=?").
		And("member_id=?").String()
	_, err = o.Raw(sql, folderId, userId).QueryRows(&favs)
	return
}

// UpdateSunFavorites updates SunFavorites by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunFavoritesById(m *SunFavorites) (err error) {
	o := orm.NewOrm()
	v := SunFavorites{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunFavorites deletes SunFavorites by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunFavorites(id int) (err error) {
	o := orm.NewOrm()
	v := SunFavorites{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunFavorites{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//FindFavoritesByFolderId 校验商品或者商店是否已经在收藏夹里面
func FindFavoritesByFolderId(userId uint, folderId int) (favs []SunFavorites, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(SunFavorites{}).Filter("FolderId", folderId).Filter("MemberId", userId).All(&favs)
	return
}

//RemoveFavoritesByGoodsId 根据GoodsId删除当前用户的收藏
func RemoveFavoritesByGoodsId(goodsId uint, memberId uint) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(SunFavorites{}).Filter("GoodsId", goodsId).Filter("MemberId", memberId).Delete()
	return
}
