package controllers

import (
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

// oprations for SunFavoritesFolder
type SunFavoritesFolderController struct {
	BaseController
}

//favoritesFolderData 收藏夹数据封装
type favoritesFolderData struct {
	Id         int
	FolderType int
	FolderName string
	FolderData []models.SunFavorites
}

func (c *SunFavoritesFolderController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description 为当前登录用户创建新的收藏文件夹
// @Param	body		body 	models.SunFavoritesFolder	true		"body for SunFavoritesFolder content"
// @Success 201 {int} models.SunFavoritesFolder
// @Failure 403 body is empty
// @router / [post]
func (c *SunFavoritesFolderController) Post() {
	c.Display("post", c.GetString("data"))
	v := models.SunFavoritesFolder{MemberId: c.GetUserID(), CreatedAt: c.GetCurrentTime()}
	c.BindModelWithPost(&v)
}

// @Title Get
// @Description 根据收藏文件夹ID获取当前登录用户的收藏文件夹
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunFavoritesFolder
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunFavoritesFolderController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunFavoritesFolderById(id)
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

// @Title GetAll
// @Description 获取当前登录用户的所有收藏文件夹信息
// @Success 200 {object} models.SunFavoritesFolder
// @Failure 403
// @router / [get]
func (c *SunFavoritesFolderController) GetAll() {
	userId := c.GetUserID()
	l, err := models.GetAllSunFavoritesFolderByUserID(userId)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err
	} else {
		// c.Display("id", userId)

		// favsDatas := make([]favoritesFolderData, len(l))

		// for index, folder := range l {
		// 	f := folder.(models.SunFavoritesFolder)

		// 	if favs, err := models.FindFavoritesByFolderId(userId, f.Id); err == nil {

		// 		favsDatas[index] = favoritesFolderData{
		// 			FolderName: f.FolderName,
		// 			FolderType: f.FolderType,
		// 			Id:         f.Id,
		// 			FolderData: favs,
		// 		}
		// 	} else {
		// 		c.apiResult.Code = constant.FAIL
		// 	}

		// }
		c.apiResult.Code = constant.OK
		c.apiResult.Data = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description 更新当前登录用户的特定收藏文件夹名
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunFavoritesFolder	true		"body for SunFavoritesFolder content"
// @Success 200 {object} models.SunFavoritesFolder
// @Failure 403 :id is not int
// @router / [put]
func (c *SunFavoritesFolderController) Put() {
	c.Display("put", c.GetString("data"))
	v := models.SunFavoritesFolder{}
	c.BindModelWithPut(&v, "FolderType", "FolderName")
}

// @Title Delete
// @Description delete the SunFavoritesFolder
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunFavoritesFolderController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.RemoveFavoritesFolderWithAllFavorites(id); err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}
