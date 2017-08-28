package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunMember
type SunMemberController struct {
	BaseController
}

func (c *SunMemberController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Password", c.Password)
	c.Mapping("NameNoExist", c.NameNoExist)
	c.Mapping("EmailNoExist", c.EmailNoExist)
}

// @Title Post
// @Description create SunMember
// @Param	body		body 	models.SunMember	true		"body for SunMember content"
// @Success 201 {int} models.SunMember
// @Failure 403 body is empty
// @router / [post]
func (c *SunMemberController) Post() {
	var v models.SunMember
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSunMember(&v); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v
		} else {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.InvalidRequestData
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Get
// @Description get SunMember by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunMember
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunMemberController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunMemberById(id)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunMember
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunMember
// @Failure 403
// @router / [get]
func (c *SunMemberController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.apiResult.Code = constant.InvalidKeyValue
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSunMember(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = l
	}
	c.ServeJSON()
}

//用户密码验证
// @router /password [post]
func (c *SunMemberController) Password() {
	password := c.GetString("oldPassword")
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	password = hex.EncodeToString(md5Ctx.Sum(nil))
	m := models.SunMember{Id: int(c.GetUserID()), MemberPasswd: password}
	if m.CheckMemberPassword() {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

//用户名验证
// @router /name [post]
func (c *SunMemberController) NameNoExist() {
	m := models.SunMember{Id: int(c.GetUserID()), MemberName: c.GetString("MemberName")}
	if m.CheckedName() {
		c.apiResult.Code = constant.FAIL
	} else {
		c.apiResult.Code = constant.OK
	}
	c.ServeJSON()
}

//用户邮箱验证
// @router /email [post]
func (c *SunMemberController) EmailNoExist() {
	m := models.SunMember{Id: int(c.GetUserID()), MemberEmail: c.GetString("MemberEmail")}
	if m.CheckedEmail() {
		c.apiResult.Code = constant.FAIL
	} else {
		c.apiResult.Code = constant.OK
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunMember
// @Param	body		body 	models.SunMember	true		"body for SunMember content"
// @Success 200 {object} models.SunMember
// @Failure 403 :id is empty
// @router / [put]
func (c *SunMemberController) Put() {
	// m := models.SunMember{}
	c.Display("memberdata", c.GetString("fields"))
	fields := strings.Split(c.GetString("fields"), ",")

	v := &models.SunMember{}
	if err := json.Unmarshal([]byte(c.GetString("data")), v); err == nil {
		c.Display("member", v)
		if v.MemberPasswd != "" {
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(v.MemberPasswd))
			v.MemberPasswd = hex.EncodeToString(md5Ctx.Sum(nil))
		}

		if err := models.SaveByID(v, fields...); err == nil {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the SunMember
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunMemberController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunMember(id); err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.ActionFaild
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunMemberExt
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunMemberExt
// @Failure 403
// @router /membertree [get]
func (c *SunMemberController) GetAllMemberTree() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.apiResult.Code = constant.InvalidKeyValue
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllMemberTree(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = l
	}
	c.ServeJSON()
}
