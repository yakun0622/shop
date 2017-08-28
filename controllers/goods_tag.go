package controllers

import (
	"strings"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

type GoodsTag struct {
	BaseController
}

func (c *GoodsTag) URLMapping() {
	c.Mapping("GetTagsByGoodsId", c.GetTagsByGoodsId)
	// c.Mapping("GetOne", c.GetOne)
	c.Mapping("Save", c.Save)
	c.Mapping("Remove", c.Remove)
}

// 根据goodsId 获取相应tags
// @router / [get]
func (c *GoodsTag) GetTagsByGoodsId() {
	if l, err := models.GetTagsByGoodsId(c.GetString("GoodsIds"), c.GetString("GroupIds")); err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = l
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}

// @router / [post]
func (c *GoodsTag) Save() {
	goodsAndTags := strings.Split(c.GetString("GoodsAndTags"), ",")
	if err := models.SaveGoodsTags(goodsAndTags); err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}

// @router /remove [post]
func (c *GoodsTag) Remove() {
	goodsId, _ := c.GetInt("GoodsId")
	tagId, _ := c.GetInt("TagId")
	err := models.RemoveGoodsTag(goodsId, tagId)
	if err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}

	c.ServeJSON()
}
