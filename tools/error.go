package tools

import "github.com/astaxie/beego"

/**
error 处理
*/
func HasError(err error, msg string) bool {
	if err != nil {
		beego.Error("接口返回信息：", msg)
		beego.Warning(err)
		return false
	}
	return true
}