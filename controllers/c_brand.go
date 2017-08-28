package controllers

import (
	"strconv"

	"github.com/yakun0622/shop/constant"

	"github.com/yakun0622/shop/models"
)

//BrandController oprations for member Login
type BrandController struct {
	BaseController
}

func (c *BrandController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// @Title Post
// @Description login user
// @Param	class_id    query   string  true	"goods_class_id"
// @Success 201 {object} models.SunBrand
// @Failure 403
// @router / [get]
func (c *BrandController) GetAll() {
	id, err := strconv.Atoi(c.GetString("class_id"))
	c.Display("brand", id)
	if err != nil {
		c.apiResult.Code = constant.InvalidRequestData
	} else {
		childIDs, err := models.GetGoodsClassChildList(id)
		if err != nil {
			c.apiResult.Code = constant.FindNoData
		} else {
			l, err := models.GetAllSunBrandByClassID(childIDs)
			if err != nil {
				c.apiResult.Code = constant.FindNoData
			} else {
				c.apiResult.Code = constant.OK
				c.apiResult.Data = l
			}
		}
	}
	c.ServeJSON()
}
