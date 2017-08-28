package controllers

import (
	"encoding/json"
	"strings"
	"time"

	"errors"

	"fmt"
	"github.com/astaxie/beego"
	beegoUtils "github.com/astaxie/beego/utils"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
	"github.com/yakun0622/shop/utils"
)

//BaseController controller 基类
type BaseController struct {
	beego.Controller
	apiResult APIResult
}

// APIResult API返回数据结构体,Msg:友好提示错误，Error：详细错误描述
type APIResult struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Count     int         `json:"count"`
	Msg       string      `json:"msg"`
	MsgValues []string    `json:"-"`
	//Error作为传入使用，并不输出，，，最终由ErrorString输出
	Error     error  `json:"-"`
	ErrString string `json:"error"`
}

// GetDataString 从请求中户获取数据
func (c *BaseController) GetDataString() string {
	return c.GetString("data")
}

func (c *BaseController) GetDataBytes() []byte {
	return []byte(c.GetString("data"))
}

//Display 打印
func (c *BaseController) Display(value ...interface{}) {
	beegoUtils.Display(value...)
}

func (c *BaseController) Exist(model interface{}, key string, value interface{}) {
	if models.IsExist(model, key, value) {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}

func (c *BaseController) NotExist(model interface{}, key string, value interface{}) {
	if models.IsExist(model, key, value) {
		c.apiResult.Code = constant.FAIL
	} else {
		c.apiResult.Code = constant.OK
	}

	c.ServeJSON()
}

//GetAllSeachParams 组装get all的参数
func (c *BaseController) GetAllSeachParams() (map[string]string, []string, []string, []string, int64, int64, error) {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 20
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
				return query, fields, sortby, order, offset, limit, errors.New("wrong key value pair")
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	return query, fields, sortby, order, offset, limit, nil

}

// BindModelWithPut PUT请求，通用处理函数
func (c *BaseController) BindModelWithPut(model interface{}, fields ...string) {
	if f := c.GetString("fields"); f != "" {
		fields = strings.Split(f, ",")
	}

	//如果data为空，前端使用formdata直接设置的键值对数据，由对应controller手动赋值
	//如果有则解析json
	if data := c.GetString("data"); data == "" {
		if err := models.SaveByID(model, fields...); err == nil {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		if err := json.Unmarshal([]byte(data), model); err == nil {
			c.Display("putmodel", model)
			if err := models.SaveByID(model, fields...); err == nil {
				c.apiResult.Code = constant.OK
			} else {
				c.apiResult.Code = constant.FAIL
				c.apiResult.Error = err
			}
		} else {
			c.apiResult.Code = constant.FAIL
			c.apiResult.Error = err
		}
	}
	c.ServeJSON()
}

//BindModelWithPost POST请求，通用处理函数
func (c *BaseController) BindModelWithPost(model interface{}) {
	if err := json.Unmarshal([]byte(c.GetString("data")), model); err == nil {
		c.Display("bindmodel", model)
		if id, err := models.Save(model); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = id
		} else {
			c.apiResult.Code = constant.FAIL
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.FAIL
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

//Remove DELETE请求，通用处理函数
func (c *BaseController) Remove(model interface{}) {
	if err := models.Remove(model); err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

//GetUserID 获取当前登陆用户的ID
func (c *BaseController) GetUserID() uint {
	return uint(c.GetCurrentUser().Id)
}

//GetCurrentStoreID 获取当前商户
func (c *BaseController) GetCurrentStoreID() uint {
	return uint(c.GetCurrentUser().StoreID)
}

// GetCurrentUser 获取当前登录用户信息
func (c *BaseController) GetCurrentUser() models.PayLoad {
	user, _ := utils.GetUserFromToken(c.Controller.Ctx.Request.Header.Get("Authorization"))
	fmt.Println("当前登录用户信息\n", user)
	return user
}

//GetCurrentTime 获取当前时间
func (c *BaseController) GetCurrentTime() uint {
	return uint(time.Now().UnixNano() / 1e9)
}

// ServeJSON 处理HTTP回传数据(apiResult)
func (c *BaseController) ServeJSON(encoding ...bool) {
	var (
		hasIndent   = true
		hasEncoding = false
	)
	if beego.BConfig.RunMode == beego.PROD {
		hasIndent = false
	}
	if len(encoding) > 0 && encoding[0] == true {
		hasEncoding = true
	}
	if c.apiResult.Code != 0 {
		c.apiResult.Msg = constant.StatusText(c.apiResult.Code, c.apiResult.MsgValues...)
	}
	if c.apiResult.Error != nil {
		c.apiResult.ErrString = c.apiResult.Error.Error()
	}
	c.Ctx.Output.JSON(c.apiResult, hasIndent, hasEncoding)
}

// IsCurrentMemberRight 校验当前token用户是否跟实际操作用户相同,不会直接输出结果
func (c *BaseController) IsCurrentMemberRight(id int) bool {
	user, err := utils.GetUserFromToken(c.Controller.Ctx.Request.Header.Get("Authorization"))
	if err != nil {
		return false
	}
	return user.Id == id
}

//ValidUser 校验用户封装，用户不对，直接输出结果，建议调用结束（用户异常的情况）直接：return
//如：
//		if !c.ValidUser(int(v.BuyerId)) {
//			return
//		}
func (c *BaseController) ValidUser(id int) bool {
	user, err := utils.GetUserFromToken(c.Controller.Ctx.Request.Header.Get("Authorization"))
	if err != nil || user.Id != id {
		c.apiResult.Code = constant.InvalidAction
		c.ServeJSON()
		return false
	}
	return true
}

//GetCurrentToken 获取当前Token
func (c *BaseController) GetCurrentToken() string {
	tokenString := c.Controller.Ctx.Request.Header.Get("Authorization")
	//无token
	if strings.Contains(strings.ToLower(tokenString), "bearer") {
		return tokenString[7:]
	}
	return ""
}

/**
接口返回值统一处理，包含了错误处理
*/
func (c *BaseController) HandleResult(err error, msg string, code int, data interface{}, count int) {
	c.apiResult.Code = code
	c.apiResult.Data = data
	c.apiResult.Count = count
	if code == constant.FAIL {
		beego.Error("接口返回信息：", msg)
	} else {
		beego.Info("接口返回信息：", msg)
	}
	if err != nil {
		beego.Warning(err)
	}
	c.ServeJSON()
}
