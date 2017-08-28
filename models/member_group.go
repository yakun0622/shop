package models

// "strconv"

type MemberGroup struct {
	MgId             uint64 `orm:"column(mg_id);auto"`
	GroupId          uint64 `orm:"column(group_id)"`
	MemberId         uint   `orm:"column(member_id)"`
	MemberName       string `orm:"column(member_name)"`
	MemberTruename   string `orm:"column(member_truename)"`
	MemberAvatar     string `orm:"column(member_avatar);size(50);null"`
	MemberMobile     string `orm:"column(member_mobile);size(11);null"`
	RoleId           uint64 `orm:"column(role_id)"`
	RoleName         string `orm:"column(role_name);size(64)"`
	RoleDesc         string `orm:"column(role_desc);size(255)"`
	RolePermission   uint64 `orm:"column(role_permission)"`
	ApproveLevel     uint8  `orm:"column(approve_level)"`
	//0 申请失败  1申请状态  2通过申请但未分配工作 3 工作状态 4 转换角色  5自己离开 6被组开除
	Status           uint8 `orm:"column(status)"`
	Levels           uint8 `orm:"column(levels)"`
	ApplyTime        uint  `orm:"column(apply_time)"`
	JoinTime         uint  `orm:"column(join_time)"`
	StartTime        uint  `orm:"column(start_time)"`
	LeaveTime        uint  `orm:"column(leave_time)"`
	LeaveReason      uint  `orm:"column(leave_reason)"`

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

func (m *MemberGroup) TableName() string {
	return "sun_member_group"
}

func ApplyJoinGroup(groupId uint64, memberId uint, applyTime uint) int64 {
	o := Orm()

	r, err := o.Raw("INSERT INTO sun_member_group (group_id, member_id, apply_time) VALUE (?, ?, ?)", groupId, memberId, applyTime).Exec()

	if err == nil {
		if id, err := r.LastInsertId(); err == nil {
			return id
		}
		return 0
	}
	return 0
}


func ImportMemberGroup(groupId uint64, memberId uint, applyTime uint, status int, joinTime uint) int64 {
	o := Orm()

	r, err := o.Raw("INSERT INTO sun_member_group (group_id, member_id, apply_time, status, join_time) VALUE (?, ?, ?, ?, ?)", groupId, memberId, applyTime, status, joinTime).Exec()

	if err == nil {
		if id, err := r.LastInsertId(); err == nil {
			return id
		}
		return 0
	}
	return 0
}

/**
 * 如果memberId为空则是根据groupId找出组下的成员
 * 如果groupId为空则是根据memberId找出成员的组下
 * 如果都不为空则查出相关成员和组数据
 */
func GetJoinGroup(memberId string, groupId string, status string) ([]MemberGroup, error) {
	o, q := GetQueryBuilder()
	var joins []MemberGroup

	q = q.Select("*").
		From("sun_role as r").
		RightJoin("sun_member_group as mg").On("mg.role_id = r.role_id")

	if groupId != "" {
		q = q.InnerJoin("sun_member as m").On("m.member_id = mg.member_id")
	}

	if memberId != "" {
		q = q.InnerJoin("sun_group as g").On("g.group_id = mg.group_id")
	}

	if groupId != "" {
		q = q.Where("mg.group_id=" + groupId)
	}

	if memberId != "" {
		if groupId == "" {
			q = q.Where("mg.member_id=" + memberId)
		} else {
			q = q.And("mg.member_id=" + memberId)
		}
	}

	if status != "" {
		q = q.And("mg.status in(" + status + ")")
	} else {
		q = q.And("mg.status > 0")
	}

	// q = q.OrderBy("mg.apply_time DESC,mg.join_time DESC,mg.start_time DESC")

	sql := q.String()

	_, err := o.Raw(sql).QueryRows(&joins)
	// Display("sql", sql)
	return joins, err
}

func PassJoinGroup(id int64, groupId int64, roleId int64, joinTime uint) bool {
	o := Orm()
	var startTime uint
	var status int = 2

	if roleId != 0 {
		startTime = joinTime
		status = 3
	}

	_, err := o.Raw("UPDATE sun_member_group SET group_id=?, role_id=?, join_time=?, start_time=?, status=? WHERE mg_id=?",
		groupId, roleId, joinTime, startTime, status, id).Exec()

	if err == nil {
		return true
	}
	return false
}

func AddJoinGroupRole(id int64, roleId int64, groupId int64, startTime uint) bool {
	o := Orm()

	_, err := o.Raw("UPDATE sun_member_group SET role_id=?, group_id=?, start_time=?, status=3 WHERE mg_id=?",
		roleId, groupId, startTime, id).Exec()

	if err == nil {
		return true
	}
	return false
}

func ChangeJoinGroupRole(id int64, roleId int64, groupId int64, memberId int, applyTime int, joinTime int, time uint) (newId int64, err error) {
	o := Orm()

	o.Begin()

	r, rawErr := o.Raw("INSERT INTO sun_member_group (group_id, member_id, role_id, status, apply_time, join_time, start_time) VALUE (?, ?, ?, 3, ?, ?, ?)",
		groupId, memberId, roleId, applyTime, joinTime, time).Exec()

	if rawErr == nil {
		if newId, err = r.LastInsertId(); err != nil {
			o.Rollback()
			return
		}
	} else {
		err = rawErr
		o.Rollback()
		return
	}

	_, err = o.Raw("UPDATE sun_member_group SET status=4, leave_time=? WHERE mg_id=?",
		time, id).Exec()

	if err != nil {
		o.Rollback()
		return
	}

	o.Commit()
	return
}

func RemoveJoinGroupById(id int64, status int, leaveReason string, leaveTime uint) bool {
	o := Orm()

	_, err := o.Raw("UPDATE sun_member_group SET leave_time=?, status=?, leave_reason=? WHERE mg_id=?",
		leaveTime, status, leaveReason, id).Exec()

	if err == nil {
		return true
	}
	return false
}
