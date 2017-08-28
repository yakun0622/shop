package controllers

import "github.com/yakun0622/shop/constant"

//ErrorController http error controller
type ErrorController struct {
	BaseController
}

//Error404 404 http error
func (c *ErrorController) Error404() {
	c.apiResult.Code = constant.Status404
	c.ServeJSON()
}

//Error501 501 http error
func (c *ErrorController) Error501() {
	c.apiResult.Code = constant.Status501
	c.ServeJSON()
}
