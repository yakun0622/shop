package controllers

import (
	"encoding/json"
	"github.com/yakun0622/shop/models"
	"strconv"
	"strings"

	"github.com/yakun0622/shop/tools"
	"github.com/yakun0622/shop/constant"
)

// oprations for SunAlbumPic
type SunAlbumPicController struct {
	BaseController
}

func (c *SunAlbumPicController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create SunAlbumPic
// @Param	body		body 	models.SunAlbumPic	true		"body for SunAlbumPic content"
// @Success 201 {int} models.SunAlbumPic
// @Failure 403 body is empty
// @router / [post]
func (c *SunAlbumPicController) Post() {
	stroreId := c.GetCurrentStoreID()
	width := c.GetString("Width")
	height := c.GetString("Height")
	f, h, err := c.GetFile("Image")
	size, _ := c.GetInt("Size")
	if err == nil {
		albumPic := &models.SunAlbumPic{}
		fileNames := strings.Split(h.Filename, ".")
		albumPic.ApicName = fileNames[0]
		albumPic.ApicCover = tools.GetStoreImageCover(int(stroreId), fileNames[1])
		albumPic.UploadTime = tools.GetTime()
		albumPic.ApicSize = uint(size)
		albumPic.ApicSpec = width + "×" + height
		albumPic.StoreId = stroreId
		_, err := models.Save(albumPic)
		if err == nil {
			c.SaveToFile("Image", tools.GetStoreImageSavePath(int(stroreId), albumPic.ApicCover))
			c.apiResult.Code = constant.OK
			c.apiResult.Data = albumPic
		} else {
			c.Display("save", err)
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.Display("getfile", err)
		c.apiResult.Code = constant.FAIL
	}
	f.Close()
	c.ServeJSON()
}

// @Title Get
// @Description get SunAlbumPic by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunAlbumPic
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunAlbumPicController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunAlbumPicById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get SunAlbumPic
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunAlbumPic
// @Failure 403
// @router / [get]
func (c *SunAlbumPicController) GetAll() {
	page, _ := c.GetInt("Page")
	pics, count, err := models.GetAllSunAlbumPic(c.GetCurrentStoreID(), page)
	if err == nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = pics
		c.apiResult.Count = count
	} else {
		c.Display("GetAll", err)
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunAlbumPic
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunAlbumPic	true		"body for SunAlbumPic content"
// @Success 200 {object} models.SunAlbumPic
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunAlbumPicController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunAlbumPic{Id: uint(id)}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunAlbumPicById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the SunAlbumPic
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunAlbumPicController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunAlbumPic(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
