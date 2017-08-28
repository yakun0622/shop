package controllers

import "github.com/yakun0622/shop/redis"

//LogoutController oprations for member Logout
type LogoutController struct {
	BaseController
}

func (c *LogoutController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @Title Get
// @Description login user
// @Success 201 {int} models.SunMember
// @Failure 403 body is empty
// @router / [get]
func (c *LogoutController) Get() {
	//将携带的token存入缓存，标记为已过期token
	token := c.GetCurrentToken()
	redis.Instance().Delete("token." + token)
	c.apiResult.Code = 1
	c.ServeJSON()
}
