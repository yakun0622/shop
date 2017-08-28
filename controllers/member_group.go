package controllers

import (
	// "encoding/json"
	// "strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

type MemberGroup struct {
	BaseController
}

func (c *MemberGroup) URLMapping() {
	c.Mapping("Apple", c.Apple)
	// c.Mapping("GetOne", c.GetOne)
	c.Mapping("Get", c.Get)
	// c.Mapping("Search", c.Search)
	c.Mapping("Pass", c.Pass)
	c.Mapping("AddRole", c.AddRole)
	c.Mapping("ChangeRole", c.ChangeRole)
	c.Mapping("Delete", c.Delete)
	// c.Mapping("GetGoup", c.GetGroupByUserID)
}

// @router / [get]
func (c *MemberGroup) Get() {
	groups, err := models.GetJoinGroup(c.GetString("MemberId"), c.GetString("GroupId"), c.GetString("Status"))
	if err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = groups
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}

// @Title Post
// @Description create MemberGroup
// @router / [post]
func (c *MemberGroup) Apple() {
	groupId, _ := c.GetInt64("GroupId")

	id := models.ApplyJoinGroup(uint64(groupId), c.GetUserID(), c.GetCurrentTime())
	if id != 0 {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = id
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @router / [put]
func (c *MemberGroup) Pass() {
	id, _ := c.GetInt64("Id")
	groupId, _ := c.GetInt64("GroupId")
	roleId, _ := c.GetInt64("RoleId")
	if models.PassJoinGroup(id, groupId, roleId, c.GetCurrentTime()) {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @router /addrole [put]
func (c *MemberGroup) AddRole() {
	id, _ := c.GetInt64("Id")
	roleId, _ := c.GetInt64("RoleId")
	groupId, _ := c.GetInt64("GroupId")
	if models.AddJoinGroupRole(id, roleId, groupId, c.GetCurrentTime()) {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @router /changerole [put]
func (c *MemberGroup) ChangeRole() {
	id, _ := c.GetInt64("Id")
	roleId, _ := c.GetInt64("RoleId")
	groupId, _ := c.GetInt64("GroupId")
	memberId, _ := c.GetInt("MemberId")
	applyTime, _ := c.GetInt("ApplyTime")
	joinTime, _ := c.GetInt("JoinTime")
	if id, err := models.ChangeJoinGroupRole(id, roleId, groupId, memberId, applyTime, joinTime, c.GetCurrentTime()); err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = id
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @router / [delete]
func (c *MemberGroup) Delete() {
	id, _ := c.GetInt64("Id")
	status, _ := c.GetInt("Status")
	if models.RemoveJoinGroupById(id, status, c.GetString("LeaveReason"), c.GetCurrentTime()) {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @router /import [get]
func (c *MemberGroup) Import() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 9999
	var offset int64 = 0
	query, fields, sortby, order, offset, limit, err := c.GetAllSeachParams()
	if err != nil {
		c.apiResult.Code = constant.InvalidKeyValue
		c.apiResult.Error = err
		c.apiResult.Data = nil
		c.ServeJSON()
	}
	//先查出所有的member信息
	memberList, _ := models.GetAllSunMemberExt(query, fields, sortby, order, offset, limit)
	if len(memberList) > 0 {
		for _, v := range memberList {
			member := v.(models.SunMemberExt)
			//找出顶层组
			groupInfo, _ := models.GetRootGroup(uint(member.MemberGroupId))
			models.ImportMemberGroup(uint64(groupInfo.Id), uint(member.MemberUser.Id), c.GetCurrentTime(), 2, c.GetCurrentTime())
		}
	}
}
