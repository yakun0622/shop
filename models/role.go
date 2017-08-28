package models

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/yakun0622/shop/constant/permission"
)

type Role struct {
	Id             uint    `orm:"column(role_id);auto"`
	RoleName       string  `orm:"column(role_name);size(64)"`
	RoleDesc       string  `orm:"column(role_desc);size(255)"`
	GroupId        uint64  `orm:"column(group_id)"`
	RolePermission uint64  `orm:"column(role_permission)"`
	RoleCtime      uint    `orm:"column(role_ctime)"`
	ApproveLevel   uint8   `orm:"column(approve_level)"`
	ApproveOrder   float64 `orm:"column(approve_order);digits(10);decimals(2)"`
	ApproveMonth   float64 `orm:"column(approve_month);digits(10);decimals(2)"`
	ApprovedMonth  float64 `orm:"column(approved_month);digits(10);decimals(2)"`
	ApproveYear    float64 `orm:"column(approve_year);digits(10);decimals(2)"`
	ApprovedYear   float64 `orm:"column(approved_year);digits(10);decimals(2)"`
	RoleLock       uint8   `orm:"column(role_lock)"`
}

func (self *Role) TableName() string {
	return "sun_role"
}

func init() {
	orm.RegisterModel(new(Role))
}

func GetAllRole(query map[string]string, fields []string, sortby []string, order []string,
offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Role))
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

	var l []Role
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

func GetAllApproveRoleIdAndGroupId(groupId uint, groupParentId uint, approveLevel int, orderType int8, tagIds string) (firstRoleId uint, firstGroupId uint, allRoleAndGroup string, error error) {
	roleIds := GetApproveRoleIdsByGroupId(groupId, approveLevel, orderType, tagIds)
	Display("roleIds", roleIds)
	if roleIds != nil {
		firstGroupId = groupId
		firstRoleIdStr := roleIds[0]
		roleId, _ := strconv.Atoi(firstRoleIdStr)
		firstRoleId = uint(roleId)
		allRoleAndGroup += strconv.Itoa(int(groupId)) + ":" + strings.Join(roleIds, ",")

		if groupParentId != 0 {
			for {
				roleIds := GetApproveRoleIdsByGroupId(groupParentId, 0, orderType, tagIds)

				if roleIds != nil {
					allRoleAndGroup += ";" + strconv.Itoa(int(groupParentId)) + ":" + strings.Join(roleIds, ",")
				}

				o, q := GetQueryBuilder()
				sql := q.Select("group_parent").From("sun_group").Where("group_id=?").String()
				error = o.Raw(sql, groupParentId).QueryRow(&groupParentId)
				if groupParentId == 0 {
					break
				}
			}
		}
		allRoleAndGroup += "|" + firstRoleIdStr
		Display("allRoleAndGroup", allRoleAndGroup)
	}
	return
}

func GetApproveRoleIdsByGroupId(groupId uint, approveLevel int, orderType int8, tagIds string) (roleIDs []string) {
	o, q := GetQueryBuilder()
	var roleIds orm.ParamsList

	q = q.Select("r.role_id").From("sun_role as r")

	if tagIds != "" {
		q = q.InnerJoin("sun_tag_role as tr").
			On("r.role_id = tr.role_id")
	}

	q = q.Where("r.approve_level > ?").
		And("r.group_id=?").
		And("r.role_lock = 1")

	if tagIds != "" {
		q = q.And("tr.tag_id in (" + tagIds + ")")
		q = q.GroupBy("r.role_id")
		Display("role_tagIds", tagIds)
	}

	switch orderType {
	case 1:
		emergency := strconv.Itoa(permission.EMERGENCY)
		q = q.And("(r.role_permission & " + emergency + ") != 0")
	case 2:
		welfare := strconv.Itoa(permission.WELFARE)
		q = q.And("(r.role_permission & " + welfare + ") != 0")
	}

	sql := q.OrderBy("r.approve_level").String()
	_, err := o.Raw(sql, approveLevel, groupId).ValuesFlat(&roleIds)
	if err != nil {
		Display("role-approve", err, "sql", sql)
		return nil
	}

	for _, roleId := range roleIds {
		roleIDs = append(roleIDs, roleId.(string))
	}

	return
}

func GetApproveRoleIdAndGroupId(orderId string) (roleId string, groupId string, newApproves string, err error) {
	o, q := GetQueryBuilder()
	var sql string
	sql = q.Select("approvers").From("sun_order").Where("order_id=?").String()
	var approvers string
	err = o.Raw(sql, orderId).QueryRow(&approvers)
	Display("GetApproveRoleIdAndGroupId-order", approvers)
	if err != nil {
		return
	}

	//解析审批流字符串
	approversSlice := strings.Split(approvers, "|")
	currentRoleId := approversSlice[1]

	approversGroups := strings.Split(approversSlice[0], ";")
	for _, approversGroup := range approversGroups {
		groupAndRoles := strings.Split(approversGroup, ":")
		groupId = groupAndRoles[0]
		roleIds := strings.Split(groupAndRoles[1], ",")
		lastRoleIdIndex := len(roleIds) - 1
		for i, id := range roleIds {
			if currentRoleId == id {
				if lastRoleIdIndex == i {
					return
				} else {
					newApproves = approversSlice[0] + "|" + roleIds[i + 1]
					roleId = roleIds[i + 1]
					return
				}
			}
		}
	}

	return
}

//func GetApproveRoleIdAndGroupId(groupId uint, groupParentId uint, approveLevel int, orderType int8, tagIds string) (uint, uint, error) {
//	o, q := GetQueryBuilder()
//	var roleIds orm.ParamsList
//
//	q = q.Select("r.role_id").From("sun_role as r")
//
//	if tagIds != "" {
//		q = q.InnerJoin("sun_tag_role as tr").
//			On("r.role_id = tr.role_id")
//	}
//
//	q = q.Where("r.approve_level > ?").
//		And("r.group_id=?").
//		And("r.role_lock = 1")
//
//	if tagIds != "" {
//		q = q.And("tr.tag_id in (" + tagIds + ")")
//		Display("role_tagIds", tagIds)
//	}
//
//	switch orderType {
//	case 1:
//		emergency := strconv.Itoa(permission.EMERGENCY)
//		q = q.And("(r.role_permission & " + emergency + ") != 0")
//	case 2:
//		welfare := strconv.Itoa(permission.WELFARE)
//		q = q.And("(r.role_permission & " + welfare + ") != 0")
//	}
//
//	sql := q.OrderBy("r.approve_level").String()
//	num, err := o.Raw(sql, approveLevel, groupId).ValuesFlat(&roleIds)
//	if err != nil {
//		Display("role", err, "sql", sql)
//		return 0, 0, err
//	}
//	if num == 0 {
//		if groupParentId == 0 {
//			return 0, 0, nil
//		} else {
//			var groupParentIds orm.ParamsList
//			_, q := GetQueryBuilder()
//			sql := q.Select("group_parent").From("sun_group").Where("group_id = ?").String()
//			_, err := o.Raw(sql, groupParentId).ValuesFlat(&groupParentIds)
//			if err != nil {
//				return 0, 0, err
//			} else {
//				parentId, _ := strconv.Atoi(groupParentIds[0].(string))
//				return GetApproveRoleIdAndGroupId(groupParentId, uint(parentId), 0, orderType, tagIds)
//			}
//		}
//	}
//	roleId, _ := strconv.Atoi(roleIds[0].(string))
//	return uint(roleId), groupId, nil
//}
