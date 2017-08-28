package controllers

import (
	"encoding/json"

	"github.com/dchest/captcha"
	"github.com/yakun0622/shop/constant"
)

type CaptchaController struct {
	BaseController
}

type captchaModel struct {
	CaptchaID    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
}

func (c *CaptchaController) Get() {
	c.apiResult.Code = constant.OK
	c.apiResult.Data = captcha.New()
	c.ServeJSON()
}

func (c *CaptchaController) Post() {
	var cm captchaModel
	json.Unmarshal([]byte(c.GetDataString()), &cm)
	b := captcha.VerifyString(cm.CaptchaID, cm.CaptchaValue)
	if b {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.WrongCaptcha
	}
	c.ServeJSON()
}
