package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/dchest/captcha"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yakun0622/shop/redis"
	_ "github.com/yakun0622/shop/routers"
)

func init() {
	//数据库设置
	orm.RegisterDataBase("default", "mysql", "hinew:hinew,2017@tcp(139.196.201.45:3306)/hinew")
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/sunlinked")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.SetLogFuncCall(true)
		beego.BConfig.Log.FileLineNum = true
		beego.SetLogger("file", `{"filename":"logs/log.log"}`)
	}
	//orm.Debug = true
	captcha.SetCustomStore(&redis.CaptchaStroe{})
	beego.Run()
}
