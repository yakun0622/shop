package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SunGroup struct {
	Id               int    `orm:"column(group_id);auto"`
	GroupName        string `orm:"column(group_name);size(64)"`
	GroupLogo        string `orm:"column(group_logo);size(45);null"`
	GroupType        uint8  `orm:"column(group_type)"`
	GroupDesc        string `orm:"column(group_desc);size(500)"`
	GroupTel         string `orm:"column(group_tel);size(25)"`
	GroupMemberCount uint   `orm:"column(group_memberCount)"`
	GroupParent      uint   `orm:"column(group_parent)"`
	GroupAddress     string `orm:"column(group_address);size(200);null"`
	GroupEmail       string `orm:"column(group_email);size(45);null"`
	GroupAreaId      uint   `orm:"column(group_area_id);null"`
	GroupCityId      uint   `orm:"column(group_city_id);null"`
	GroupDistrictId  uint   `orm:"column(group_district_id);null"`
	GroupAreaInfo    string `orm:"column(group_Area_info);size(128)null"`
	GroupOwnerId     uint   `orm:"column(group_ownerId);null"`
	GroupOwnerName   string `orm:"column(group_ownerName);size(45);null"`
	GroupLevels      uint8  `orm:"column(group_levels)"`
	GroupCtime       uint   `orm:"column(group_ctime)"`
	GroupState       int8   `orm:"column(group_state)"`
	GroupBelong      uint64 `orm:"column(group_belong)"`
}

func (t *SunGroup) TableName() string {
	return "shop_group"
}

func init() {
	orm.RegisterModel(new(SunGroup))
}

// AddSunGroup insert a new SunGroup into database and returns
// last inserted Id on success.
func AddSunGroup(m *SunGroup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunGroupById retrieves SunGroup by Id. Returns error if
// Id doesn't exist
func GetSunGroupById(id int) (v *SunGroup, err error) {
	o := orm.NewOrm()
	v = &SunGroup{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func SearchGroup(groupName string, userId uint) ([]SunGroup, int64, error) {
	var group []SunGroup
	o := Orm()

	query := o.QueryTable(new(SunGroup)).
		Filter("group_name__contains", groupName).Filter("group_parent", 0).
		Exclude("group_ownerId", userId)

	count, err := query.Count()

	if err != nil {
		return nil, 0, err
	}

	_, e := query.All(&group)

	if e != nil {
		return nil, 0, e
	}
	return group, count, nil
}

// GetAllSunGroup retrieves all SunGroup matches certain condition. Returns empty list if
// no records exist
func GetAllSunGroup(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunGroup))
	limit = 9999
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

	var l []SunGroup
	qs = qs.OrderBy(sortFields...)

	if _, err := qs.Limit(limit, offset).RelatedSel().All(&l, fields...); err == nil {
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

type joinGroup struct {
	Id               int    `orm:"column(group_id);auto"`
	GroupName        string `orm:"column(group_name);size(64)"`
	GroupLogo        string `orm:"column(group_logo);size(45);null"`
	GroupType        uint8  `orm:"column(group_type)"`
	GroupDesc        string `orm:"column(group_desc);size(500)"`
	GroupTel         string `orm:"column(group_tel);size(25)"`
	GroupMemberCount uint   `orm:"column(group_memberCount)"`
	GroupParent      uint   `orm:"column(group_parent)"`
	GroupAddress     string `orm:"column(group_address);size(200);null"`
	GroupEmail       string `orm:"column(group_email);size(45);null"`
	GroupAreaId      uint   `orm:"column(group_area_id);null"`
	GroupCityId      uint   `orm:"column(group_city_id);null"`
	GroupDistrictId  uint   `orm:"column(group_district_id);null"`
	GroupAreaInfo    string `orm:"column(group_Area_info);size(128)null"`
	GroupOwnerId     uint   `orm:"column(group_ownerId);null"`
	GroupOwnerName   string `orm:"column(group_ownerName);size(45);null"`
	GroupLevels      uint8  `orm:"column(group_levels)"`
	GroupBelong      uint64 `orm:"column(group_belong)"`

	RoleId         uint    `orm:"column(role_id);"`
	RoleName       string  `orm:"column(role_name);size(64)"`
	RoleDesc       string  `orm:"column(role_desc);size(255)"`
	RolePermission uint64  `orm:"column(role_permission)"`
	ApproveLevel   uint8   `orm:"column(approve_level)"`
	ApproveOrder   float64 `orm:"column(approve_order);digits(10);decimals(2)"`
	ApproveMonth   float64 `orm:"column(approve_month);digits(10);decimals(2)"`
	ApprovedMonth  float64 `orm:"column(approved_month);digits(10);decimals(2)"`
	ApproveYear    float64 `orm:"column(approve_year);digits(10);decimals(2)"`
	ApprovedYear   float64 `orm:"column(approved_year);digits(10);decimals(2)"`
	RoleLock       uint8   `orm:"column(role_lock)"`
}

func (s *SunGroup) GetJoinGroup(memberId uint) []joinGroup {
	o, q := GetQueryBuilder()
	var joins []joinGroup
	sql := q.Select("*").
		From(s.TableName() + " as g").
		InnerJoin("shop_member_group as mg").On("g.group_id = mg.group_id").
		InnerJoin("shop_role as r").On("mg.roleid = r.role_id").
		Where("mg.member_id=" + strconv.Itoa(int(memberId))).
		String()

	num, _ := o.Raw(sql).QueryRows(&joins)
	Display("sql", sql, "num", num)
	return joins
}

// UpdateSunGroup updates SunGroup by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunGroupById(m *SunGroup) (err error) {
	o := orm.NewOrm()
	v := SunGroup{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunGroup deletes SunGroup by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunGroup(id int) (err error) {
	o := orm.NewOrm()
	v := SunGroup{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunGroup{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//GetGroupChildList 获取某公司下所有子公司，包含自己，首位为该公司的ID
func GetGroupChildList(id int) (result []interface{}, err error) {
	groupListQuery := make(map[string]string)
	groupListQuery["GroupParent"] = strconv.Itoa(id)
	groupList, err := GetAllSunGroup(groupListQuery, []string{}, []string{}, []string{}, 0, 9999)
	if err == nil {
		if len(groupList) > 0 {
			for _, group := range groupList {
				result = append(result, group)
				temp, _ := group.(SunGroup)
				fmt.Println(group)

				groupChildList, err := GetGroupChildList(temp.Id)
				if err != nil {
					return nil, err
				}
				if len(groupChildList) > 0 {
					//result = append(result, groupChildList)
					result = mergeChildGroup(groupChildList, result)
				}
			}
		}
		return result, err
	}
	return nil, err
}

func mergeChildGroup(list []interface{}, oldResult []interface{}) (result []interface{}) {
	result = oldResult
	for _, item := range list {
		result = append(result, item)
	}
	return result
}

func buildChildIds(list []interface{}) (result []string) {
	if len(list) > 0 {
		for _, item := range list {
			group, _ := item.(SunGroup)
			result = append(result, strconv.Itoa(group.Id))
		}
		fmt.Println("lit_ids......")
		fmt.Println(result)
		return
	}
	return nil
}

func GetRootGroup(groupId uint) (rootGroup *SunGroup, err error) {
	rootGroup, err = GetSunGroupById(int(groupId))
	if rootGroup.GroupParent != 0 {
		rootGroup, err = GetRootGroup(rootGroup.GroupParent)
	}
	return rootGroup, err
}
