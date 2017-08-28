package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/redis"
	"github.com/yakun0622/shop/utils"
)

//SMSController oprations for member Logout
type SMSController struct {
	BaseController
}

//templateSMS
type templateSMS struct {
	CreateDate string `json:"createDate"`
	SmsID      string `json:"smsId"`
}

//resp
type resp struct {
	RespCode    string      `json:"respCode"`
	Failure     int         `json:"failure"`
	TemplateSMS templateSMS `json:"templateSMS"`
}

//UcPaasResult ...
type UcPaasResult struct {
	Resp resp `json:"resp"`
}

func (c *SMSController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// @Title 获取短信验证码
// @Description login user
// @Param	body		body 	models.SunMember	true		"body for SunMember content"
// @Success 201 {int} models.SunMember
// @Failure 403 body is empty
// @router /:id [get]
func (c *SMSController) GetOne() {
	mobile := c.Ctx.Input.Param(":id")
	//查询缓存中是否存在验证码，存在提示需等待，一分钟限制
	if smsCode := redis.Instance().Get("sms." + mobile); smsCode != nil {
		c.apiResult.Code = constant.SmsCodeAlreadySend
		c.ServeJSON()
		return
	}
	//随机四位数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomCode := ""
	for index := 0; index < 6; index++ {
		randomCode += strconv.Itoa(r.Intn(10))
	}

	if resultStr, err := utils.SendSms(mobile, "21235", randomCode+",10"); err != nil {
		c.apiResult.Code = constant.ActionFaild
		c.apiResult.Error = err
	} else {
		var result UcPaasResult

		if json.Unmarshal([]byte(resultStr), &result); result.Resp.RespCode != "000000" {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = errors.New("短信发送出错，请稍后再试")
			fmt.Println(result)
		} else {
			redis.Instance().Put("sms."+mobile, randomCode, 10*time.Minute)
			c.apiResult.Code = 0
		}
	}

	c.ServeJSON()
}
