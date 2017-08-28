package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yakun0622/shop/config"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/constant/userType"
	"github.com/yakun0622/shop/models"
	"github.com/yakun0622/shop/redis"
	"github.com/yakun0622/shop/utils"
)

//MemberLoginWithCaptcha member login struct
type MemberLoginWithCaptcha struct {
	models.SunMember
	CaptchaStruct
	MemberPasswd string `orm:"column(member_passwd);size(32)"`
}

//CaptchaStruct captcha struct
type CaptchaStruct struct {
	CaptchaID    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
}

//MemberLoginController oprations for member Login
type MemberLoginController struct {
	BaseController
}

func (c *MemberLoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Exist", c.Exist)
}

// @Title Update
// @Description exist the SunMember
// @Param	body		body 	models.SunMember	true		"body for SunMember content"
// @Success 200 {object} models.SunMember
// @router /exist [post]
func (c *MemberLoginController) Exist() {
	m := models.SunMember{}
	if m.IsAccountExist(c.GetString("Account")) {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}

// @Title Post
// @Description login user
// @Param	body		body 	models.SunMember	true		"body for SunMember content"
// @Success 201 {int} models.SunMember
// @Failure 403 body is empty
// @router / [post]
func (c *MemberLoginController) Post() {
	account := c.GetString("Account")
	password := c.GetString("Password")
	if account == "" || password == "" {
		c.apiResult.Code = constant.LoginWrongNameOrPassword
	} else {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(password))
		password = hex.EncodeToString(md5Ctx.Sum(nil))

		// if captchaOK := captcha.VerifyString(v.CaptchaID, v.CaptchaValue); !captchaOK {
		// 	result["code"] = "CAP001"
		// 	result["msg"] = CONST.APICode["CAP001"]
		// } else {
		memberLogin := models.SunMember{
			MemberName:    account,
			MemberMobile:  account,
			MemberPasswd:  password,
			MemberLoginIp: c.Ctx.Input.IP(),
		}

		if utils.IsLoginCountOver(account) {
			c.apiResult.Code = constant.FailLoginCountTooMany
			c.apiResult.MsgValues = []string{strconv.Itoa(config.FailLockMinutes)}
		} else {
			if member, err := models.CheckMemberLogin(memberLogin); err == nil {

				var mapClaims jwt.MapClaims

				if store, err := models.GetStoreByMemberID(uint(member.Id)); err == nil {
					mapClaims = jwt.MapClaims{
						"id":      member.Id,
						"name":    member.MemberName,
						"storeid": store.Id,
						"type":    userType.STORE,
						"nbf":     time.Now().Unix(),
						"exp":     time.Now().Add(time.Hour * time.Duration(config.JWTTokenExpireTime)).Unix(),
					}
				} else {
					mapClaims = jwt.MapClaims{
						"id":   member.Id,
						"name": member.MemberName,
						"type": userType.USER,
						"nbf":  time.Now().Unix(),
						"exp":  time.Now().Add(time.Hour * time.Duration(config.JWTTokenExpireTime)).Unix(),
					}
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
			} else {
				utils.LoginFailBangUser(memberLogin.MemberName, memberLogin.MemberLoginIp)
				c.apiResult.Code = constant.LoginWrongNameOrPassword
			}
		}
	}
	c.ServeJSON()
}
