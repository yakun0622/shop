package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (hc *HomeController) Get() {
	hc.Data["json"] = "cb api server"
	hc.ServeJSON()
}
