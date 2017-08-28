package controllers

import (
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunFavorites
type SunFavoritesController struct {
	BaseController
}

func (c *SunFavoritesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create SunFavorites
// @Param	body		body 	models.SunFavorites	true		"body for SunFavorites content"
// @Success 201 {int} models.SunFavorites
// @Failure 403 body is empty
// @router / [post]
func (c *SunFavoritesController) Post() {
	v := models.SunFavorites{MemberId: c.GetUserID(), FavTime: c.GetCurrentTime()}
	c.BindModelWithPost(&v)
}

// @Title Get
// @Description get SunFavorites by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunFavorites
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunFavoritesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunFavoritesById(id)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		if !c.ValidUser(int(v.MemberId)) {
			return
		}
		c.apiResult.Code = constant.OK
		c.apiResult.Data = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunFavorites
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunFavorites
// @Failure 403
// @router / [get]
func (c *SunFavoritesController) GetAll() {
	folderId, _ := c.GetInt("FolderId")
	favs, err := models.GetAllSunFavoritesBuyFolderId(folderId, c.GetUserID())
	if err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = favs
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunFavorites
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunFavorites	true		"body for SunFavorites content"
// @Success 200 {object} models.SunFavorites
// @Failure 403 :id is not int
// @router /change_num [post]
func (c *SunFavoritesController) Put() {
	id, _ := c.GetInt("Id")
	goodsNum, _ := c.GetInt16("GoodsNum")
	v := models.SunFavorites{Id: id, GoodsNum: uint16(goodsNum)}
	c.BindModelWithPut(&v, "GoodsNum")
}

// @Title Delete
// @Description delete the SunFavorites
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /remove [post]
func (c *SunFavoritesController) Delete() {
	id, _ := c.GetInt("Id")
	if id == 0 {
		if goodsID, err := c.GetInt("GoodsId"); err != nil {
			c.apiResult.Code = constant.InvalidRequestData
		} else {
			if err := models.RemoveFavoritesByGoodsId(uint(goodsID), c.GetUserID()); err == nil {
				c.apiResult.Code = constant.OK
			} else {
				c.apiResult.Code = constant.FAIL
			}
		}
	} else {
		if err := models.DeleteSunFavorites(id); err == nil {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = err
		}
	}
	c.ServeJSON()
}
