package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/yakun0622/shop/config"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/constant/userType"
	"github.com/yakun0622/shop/models"
	"github.com/yakun0622/shop/redis"
	"github.com/yakun0622/shop/utils"
	"math/rand"
	"strconv"
	"time"
)

// operations for RegisterController
type RegisterControllerController struct {
	BaseController
}

type RegisterBody struct {
	Name       string
	Password   string
	Email      string
	Mobile     string
	MobileCode string
}

func (c *RegisterControllerController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetSms", c.GetSMS)
	c.Mapping("NoExist", c.NoExist)
}

// @Title 获取注册短信验证码
// @Description login user
// @Param	mobile	path	string	true	"手机号"
// @router /sms/:mobile [get]
func (c *RegisterControllerController) GetSMS() {
	mobile := c.Ctx.Input.Param(":mobile")
	//查询缓存中是否存在验证码，存在提示需等待，一分钟限制
	if smsCode := redis.Instance().Get("sms.register." + mobile); smsCode != nil {
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

// @Title Post
// @Description 判断数据不存在
// @router /noexist [post]
func (c *RegisterControllerController) NoExist() {
	v := models.SunMember{}
	// c.Display("register", c.GetString("MemberMobile"))
	mobile := c.GetString("MemberMobile")
	c.NotExist(&v, "MemberMobile", mobile)
}

// @Title Post
// @Description 用户注册
// @Param	body	body	string	true	"注册信息"
// @Success 201 "注册成功" {object}
// @Failure 403 body is empty
// @router / [post]
func (c *RegisterControllerController) Post() {
	// requestData := c.GetString("data")
	// var register RegisterBody

	mobile := c.GetString("MemberMobile")
	password := c.GetString("MemberPasswd")

	c.Display("mobile", mobile)
	c.Display("password", password)

	if mobile == "" || password == "" {
		c.apiResult.Code = constant.FAIL
	} else {
		// phoneCode := c.GetString("PhoneCode")

		// code := redis.Instance().Get("sms.register." + mobile)
		// b := make([]byte, len(code.([]uint8)))
		// for i, v := range code.([]uint8) {
		// 	b[i] = byte(v)
		// }
		// if string(b) == phoneCode {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		member := models.SunMember{
			MemberName:   mobile,
			MemberPasswd: hex.EncodeToString(md5Ctx.Sum(nil)),
			// MemberEmail:  register.Email,
			MemberMobile: mobile,
		}
		if _, err := models.Save(&member); err != nil {
			c.apiResult.Code = constant.ActionFaild
		} else {
			var mapClaims jwt.MapClaims

			mapClaims = jwt.MapClaims{
				"id":   member.Id,
				"name": member.MemberName,
				"type": userType.USER,
				"nbf":  time.Now().Unix(),
				"exp":  time.Now().Add(time.Hour * time.Duration(config.JWTTokenExpireTime)).Unix(),
			}

			//颁发jwt令牌
			token := jwt.NewWithClaims(jwt.SigningMethodRS256, mapClaims)
			if tokenString, err := token.SignedString(constant.JWTPrivateKey); err != nil {
				c.apiResult.Code = constant.ActionFaild
				c.apiResult.Error = err
			} else {
				//token 存入缓存
				redis.Instance().Put("token."+tokenString, tokenString, time.Duration(config.JWTTokenExpireTime)*time.Hour)
				// c.Ctx.SetCookie("loginToken", tokenString, 240, "/", constant.CrossCookiesDomain)
				c.apiResult.Code = constant.OK
				c.apiResult.Data = tokenString
			}
		}

		// } else {
		// 	c.apiResult.Code = constant.InvalidSmsCode
		// }
	}

	c.ServeJSON()
}
