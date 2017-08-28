package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id       uint64 `orm:"column(tag_id);auto"`
	TagName  string `orm:"column(tag_name);size(64)"`
	MemberId uint   `orm:"column(member_id)"`
	GroupId  uint64 `orm:"column(group_id)"`
	TagLock  uint8  `orm:"column(tag_lock)"`
	TagCtime uint   `orm:"column(tag_ctime)"`
}

func (t *Tag) TableName() string {
	return "sun_tag"
}

func init() {
	orm.RegisterModel(new(Tag))
}

// GetAllSunAddress retrieves all SunAddress matches certain condition. Returns empty list if
// no records exist
func GetAllTag(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Tag))
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

	var l []Tag
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

	Display("dfdf", &l)

	return nil, err
}

type UserGroupTags struct {
	GroupId uint64 `orm:"column(group_id)"`
	TagId   uint64 `orm:"column(tag_id)"`
	TagName string `orm:"column(tag_name);size(64)"`
}

func GetGroupTags(groupIds string) (l []UserGroupTags, err error) {
	o, q := GetQueryBuilder()

	sql := q.Select("group_id, tag_id, tag_name").
		From("sun_tag").
		Where("group_id IN (" + groupIds + ")").
		And("tag_lock = 1").String()

	_, err = o.Raw(sql).QueryRows(&l)
	return
}

func GetShareGroupTags(groupIds string) (l []UserGroupTags, err error) {
	o, q := GetQueryBuilder()

	sql := q.Select("tg.group_id, tg.tag_id, t.tag_name").
		From("sun_tag_group as tg").
		InnerJoin("sun_tag as t").
		On("tg.tag_id = t.tag_id").
		Where("tg.group_id IN (" + groupIds + ")").
		And("t.tag_lock = 1").String()

	_, err = o.Raw(sql).QueryRows(&l)
	return
}

type TagGroup struct {
	GroupId uint64 `orm:"column(group_id)"`
	TagId   uint64 `orm:"column(tag_id)"`
}

func ShareTags(groupIds []uint64, tagId int64) bool {
	o := orm.NewOrm()

	r, _ := o.Raw("INSERT INTO sun_tag_group ( group_id, tag_id) value (?, ?)").Prepare()

	for _, groupId := range groupIds {
		_, err := r.Exec(groupId, tagId)
		if err != nil {
			return false
		}
	}
	r.Close()
	return true
}

func UnShareTags(groupIds []uint64, tagId int64) bool {
	o := orm.NewOrm()
	r, _ := o.Raw("DELETE FROM sun_tag_group WHERE group_id = ? AND tag_id = ?").Prepare()

	for _, groupId := range groupIds {
		_, err := r.Exec(groupId, tagId)
		if err != nil {
			return false
		}
	}
	r.Close()
	return true
}

func GetTagShare(groupId string, tagId string) ([]TagGroup, error) {
	o := orm.NewOrm()
	var tags []TagGroup

	var sql string

	if groupId == "" {
		sql = "SELECT * FROM sun_tag_group WHERE tag_id IN (" + tagId + ")"
	} else {
		sql = "SELECT * FROM sun_tag_group WHERE group_id IN (" + groupId + ")"
	}

	_, err := o.Raw(sql).QueryRows(&tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

type TagRole struct {
	RoleId uint64 `orm:"column(role_id)"`
	TagId  uint64 `orm:"column(tag_id)"`
}

func RoleTags(roleId int64, tagIds []uint64) bool {
	o := orm.NewOrm()

	r, _ := o.Raw("INSERT INTO sun_tag_role ( role_id, tag_id) value (?, ?)").Prepare()

	for _, tagId := range tagIds {
		_, err := r.Exec(roleId, tagId)
		if err != nil {
			return false
		}
	}
	r.Close()
	return true
}

func RemoveRoleTags(roleId int64, tagIds []uint64) bool {
	o := orm.NewOrm()
	r, _ := o.Raw("DELETE FROM sun_tag_role WHERE role_id = ? AND tag_id = ?").Prepare()

	for _, tagId := range tagIds {
		_, err := r.Exec(roleId, tagId)
		if err != nil {
			return false
		}
	}
	r.Close()
	return true
}

func GetRoleTag(roleId string, tagId string) ([]TagRole, error) {
	o := orm.NewOrm()
	var tags []TagRole

	var sql string

	if roleId == "" {
		sql = "SELECT * FROM sun_tag_role WHERE tag_id IN (" + tagId + ")"
	} else {
		sql = "SELECT * FROM sun_tag_role WHERE role_id IN (" + roleId + ")"
	}

	_, err := o.Raw(sql).QueryRows(&tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
