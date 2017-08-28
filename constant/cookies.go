package constant

import "github.com/astaxie/beego"

//CrossCookiesDomain 子域名支持
var CrossCookiesDomain string

func init() {
	CrossCookiesDomain = beego.AppConfig.DefaultString("SmsSid", "")
}
