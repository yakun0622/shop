package routers

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
)

func init() {
	//api_key Filter
	var FilterAPI = func(ctx *context.Context) {
		apiKey := ctx.Request.URL.Query().Get("api_key")
		if apiKey != "sungoo" {
			ctx.Output.JSON("invalid apiId", false, false)
		}
	}
	beego.InsertFilter("/v1/*", beego.BeforeRouter, FilterAPI)
}
