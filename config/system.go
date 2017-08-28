package config

import "github.com/astaxie/beego"

//MaxTryLoginCount 最大登陆尝试次数
var MaxTryLoginCount int

//FailLockMinutes 超过尝试登陆次数封禁时间（分钟）
var FailLockMinutes int

//JWTTokenExpireTime JWTToken默认过期时间（小时）
var JWTTokenExpireTime int

func init() {
	FailLockMinutes = beego.AppConfig.DefaultInt("FailLockMinutes", 10)
	MaxTryLoginCount = beego.AppConfig.DefaultInt("MaxTryLoginCount", 5)
	JWTTokenExpireTime = beego.AppConfig.DefaultInt("JWTTokenExpireTime", 1440)
}
