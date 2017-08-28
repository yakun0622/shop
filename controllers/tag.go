package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

type Tag struct {
	BaseController
}

func (c *Tag) URLMapping() {
	c.Mapping("Post", c.Post)
	// c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Lock", c.Lock)
	c.Mapping("UnLock", c.UnLock)
	c.Mapping("UnShare", c.UnShare)
	c.Mapping("Share", c.Share)
	c.Mapping("GetShare", c.GetShare)
	c.Mapping("RoleTags", c.RoleTags)
	c.Mapping("RemoveRoleTags", c.RemoveRoleTags)
	c.Mapping("GetRoleTag", c.GetRoleTag)
	c.Mapping("GetUserGroupTags", c.GetUserGroupTags)
}

// @Title Get All
// @Description get Tag
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Tag
// @Failure 403
// @router / [get]
func (c *Tag) GetAll() {
	query, fields, sortby, order, offset, limit, err := c.GetAllSeachParams()
	if err != nil {
		c.apiResult.Code = constant.InvalidKeyValue
		c.apiResult.Error = err
	} else {
		l, err := models.GetAllTag(query, fields, sortby, order, offset, limit)
		if err != nil {
			c.apiResult.Code = constant.FAIL
		} else {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = l
		}
	}

	c.ServeJSON()
}

// @Title Post
// @router / [post]
func (r *Tag) Post() {
	v := models.Tag{}
	v.TagCtime = r.GetCurrentTime()
	v.TagLock = 1
	r.BindModelWithPost(&v)
}

// @Title GetUserGroupTags
// @router /user_group_tags [get]
func (c *Tag) GetUserGroupTags() {
	ownerGroupIds := c.GetString("ownerGroupIds")
	joinOwnerGroupIds := c.GetString("joinOwnerGroupIds")
	joinGroupIds := c.GetString("joinGroupIds")

	m := map[string][]models.UserGroupTags{}

	if ownerGroupIds != "" {
		if l, err := models.GetGroupTags(ownerGroupIds); err == nil {
			m["ownerGroupTags"] = l
		} else {
			c.apiResult.Code = constant.FAIL
			c.ServeJSON()
			return
		}
	}

	if joinOwnerGroupIds != "" {
		if l, err := models.GetGroupTags(joinOwnerGroupIds); err == nil {
			m["joinOwnerGroupTags"] = l
		} else {
			c.apiResult.Code = constant.FAIL
			c.ServeJSON()
			return
		}
	}

	if joinGroupIds != "" {
		if l, err := models.GetShareGroupTags(joinGroupIds); err == nil {
			m["joinGroupTags"] = l
		} else {
			c.apiResult.Code = constant.FAIL
			c.ServeJSON()
			return
		}
	}

	c.apiResult.Code = constant.OK
	c.apiResult.Data = m

	c.ServeJSON()
}

// @Title GetShare
// @router /share [get]
func (c *Tag) GetShare() {
	t, err := models.GetTagShare(c.GetString("GroupId"), c.GetString("TagId"))
	if err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = t
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Post
// @router /share [post]
func (c *Tag) Share() {
	var shares []uint64
	tagId, _ := c.GetInt64("TagId")
	c.Display("groupId", c.GetString("GroupIds"))
	if err := json.Unmarshal([]byte(c.GetString("GroupIds")), &shares); err == nil {
		if models.ShareTags(shares, tagId) {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.apiResult.Code = constant.FAIL
		c.apiResult.Data = err
	}
	c.ServeJSON()
}

// @Title Post
// @router /unshare [post]
func (c *Tag) UnShare() {
	var shares []uint64
	tagId, _ := c.GetInt64("TagId")
	c.Display("groupId", c.GetString("GroupIds"))
	if err := json.Unmarshal([]byte(c.GetString("GroupIds")), &shares); err == nil {
		if models.UnShareTags(shares, tagId) {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.apiResult.Code = constant.FAIL
		c.apiResult.Data = err
	}
	c.ServeJSON()
}

// @Title GetRoleTag
// @router /role [get]
func (c *Tag) GetRoleTag() {
	t, err := models.GetRoleTag(c.GetString("RoleId"), c.GetString("TagId"))
	if err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = t
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Post
// @router /role [post]
func (c *Tag) RoleTags() {
	var shares []uint64
	roleId, _ := c.GetInt64("RoleId")
	// c.Display("groupId", c.GetString("GroupIds"))
	if err := json.Unmarshal([]byte(c.GetString("TagIds")), &shares); err == nil {
		if models.RoleTags(roleId, shares) {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.apiResult.Code = constant.FAIL
		c.apiResult.Data = err
	}
	c.ServeJSON()
}

// @Title Post
// @router /role [delete]
func (c *Tag) RemoveRoleTags() {
	var shares []uint64
	roleId, _ := c.GetInt64("RoleId")
	// c.Display("groupId", c.GetString("GroupIds"))
	if err := json.Unmarshal([]byte(c.GetString("TagIds")), &shares); err == nil {
		if models.RemoveRoleTags(roleId, shares) {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.apiResult.Code = constant.FAIL
		c.apiResult.Data = err
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Tag
// @Success 200 {object} models.Tag
// @Failure 403 :id is not int
// @router / [put]
func (c *Tag) Put() {
	v := models.Tag{}

	c.BindModelWithPut(&v, "TagName", "GroupId", "MemberId")
}

// @Title Put
// @Description lock the Tag
// @Success 200 {object} models.Tag
// @Failure 403 :id is not int
// @router /lock/:id [put]
func (c *Tag) Lock() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Tag{Id: uint64(id), TagLock: 0}
	c.BindModelWithPut(&v, "TagLock")
}

// @Title Put
// @Description unlock the Tag
// @Success 200 {object} models.Tag
// @Failure 403 :id is not int
// @router /unlock/:id [put]
func (c *Tag) UnLock() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Tag{Id: uint64(id), TagLock: 1}
	c.BindModelWithPut(&v, "TagLock")
}
