package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yakun0622/shop/config"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/constant/userType"
	"github.com/yakun0622/shop/models"
	"github.com/yakun0622/shop/redis"
)

//AdminLoginWithCaptcha admin login struct
type AdminLoginWithCaptcha struct {
	models.SunAdmin
	CaptchaStruct
	MemberPassword string `orm:"column(member_password);size(32)" json:"member_password"`
}

//AdminLoginController oprations for admin Login
type AdminLoginController struct {
	BaseController
}

func (c *AdminLoginController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title Post
// @Description login user
// @Param	body		body 	models.SunAdmin	true		"body for SunAdmin content"
// @Success 201 {int} models.SunAdmin
// @Failure 403 body is empty
// @router / [post]
func (hc *AdminLoginController) Post() {
	var v AdminLoginWithCaptcha
	if err := json.Unmarshal([]byte(hc.GetDataString()), &v); err == nil {
		if v.AdminName == "" || v.AdminPassword == "" {
			hc.apiResult.Code = constant.LoginWrongData
		} else {
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(v.AdminPassword))
			v.AdminPassword = hex.EncodeToString(md5Ctx.Sum(nil))

			// if captchaOK := captcha.VerifyString(v.CaptchaID, v.CaptchaValue); !captchaOK {
			// 	result["code"] = "CAP001"
			// 	result["msg"] = CONST.APICode["CAP001"]
			// } else {
			if admin, err := models.CheckAdminLogin(v.AdminName, v.AdminPassword); err == nil {
				hc.Ctx.Output.SetStatus(201)
				//颁发jwt令牌
				token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
					"id":   admin.Id,
					"name": admin.AdminName,
					"type": userType.ADMIN,
					"nbf":  time.Now().Unix(),
					"exp":  time.Now().Add(time.Minute * 30).Unix(),
				})
				if tokenString, err := token.SignedString(constant.JWTPrivateKey); err != nil {
					hc.apiResult.Code = constant.ActionFaild
					hc.apiResult.Error = err
				} else {
					hc.apiResult.Code = constant.OK
					hc.apiResult.Data = tokenString
					redis.Instance().Put("token."+tokenString, tokenString, time.Duration(config.JWTTokenExpireTime)*time.Hour)
				}
			} else {
				hc.apiResult.Code = constant.LoginWrongNameOrPassword
			}
			// }
		}
	} else {
		hc.apiResult.Code = constant.InvalidRequestData
	}
	hc.ServeJSON()
}
