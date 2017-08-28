package constant

import (
	"github.com/yakun0622/shop/tools"
)

//请求API状态
const (
	FAIL = iota //失败
	OK
	//http
	Status404
	Status501
	//value
	InvalidKeyValue
	InvalidRequestData
	FindNoData
	//action
	InvalidAction
	ActionFaild
	//login
	LoginWrongNameOrPassword
	LoginWrongData
	AccountLocked
	//jwt
	JWTTokenNeeded
	JWTTokenInvalidFormat
	JWTTokenExpired
	JWTTokenInvalid
	//captcha
	WrongCaptcha
	//goodscommon
	WrongGoodsCommonData
	//goodsclass
	WrongGoodsClassData
	//goodsbrand
	WrongGoodsBrandID
	WrongGoodsBrandData
	WrongGoodsBrandChild
	//cart
	//address
	SetDefaultAddress
	//favorites
	AlreadyInSunFavorites
	//sms
	SmsCodeAlreadySend
	InvalidSmsCode
	FailLoginCountTooMany
)

var statusText = map[int]string{
	//http
	Status404: "该URL地址不存在",
	Status501: "服务器错误",
	//value
	InvalidKeyValue:    "非法键值对",
	InvalidRequestData: "请求的数据有误",
	FindNoData:         "未查询到相关数据",
	//action
	InvalidAction: "非法操作",
	ActionFaild:   "操作失败",
	//login
	LoginWrongNameOrPassword: "账号密码有误",
	LoginWrongData:           "输入的内容有误",
	AccountLocked:            "账号已被锁定",
	//jwt
	JWTTokenNeeded:        "无token",
	JWTTokenInvalidFormat: "token格式有误",
	JWTTokenExpired:       "token已过期或未生效",
	JWTTokenInvalid:       "非法token",
	//captcha
	WrongCaptcha: "验证码有误",
	//goodscommon
	WrongGoodsCommonData: "查询商品发生错误",

	//address
	SetDefaultAddress: "设置默认地址失败",
	//favorites
	AlreadyInSunFavorites: "已在收藏夹",
	//sms
	SmsCodeAlreadySend:    "短信验证码一分钟内只能请求一次",
	InvalidSmsCode:        "无效手机验证码",
	FailLoginCountTooMany: "当前登录失败次数过多，请等待{0}分钟再试！",
}

//StatusText 状态码转换成字符
func StatusText(code int, moreValues ...string) string {
	return tools.GetFormatString(statusText[code], moreValues...)
}
