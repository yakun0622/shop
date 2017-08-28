package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
	"github.com/yakun0622/shop/tools"
)

// oprations for SunGoodsCommon
type SunGoodsCommonController struct {
	BaseController
}

func (c *SunGoodsCommonController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create SunGoodsCommon
// @Param	body		body 	models.SunGoodsCommon	true		"body for SunGoodsCommon content"
// @Success 201 {int} models.SunGoodsCommon
// @Failure 403 body is empty
// @router / [post]
func (c *SunGoodsCommonController) Post() {
	var goodsCommon models.SunGoodsCommon
	var goodses []models.SunGoods

	err := json.Unmarshal([]byte(c.GetString("GoodsCommon")), &goodsCommon)
	goodsCommon.StoreId = c.GetCurrentStoreID()

	c.Display("goodscommon",goodsCommon)
	c.Display("goodsstr", c.GetString("Goodses"))
	err = json.Unmarshal([]byte(c.GetString("Goodses")), &goodses)
	c.Display("goodses",goodses)

	if err== nil {
		err = models.SaveGoodsesAndCommon(&goodsCommon, goodses)
		if err == nil {
			c.apiResult.Code = constant.OK
		} else {
			c.apiResult.Code = constant.FAIL
		}
	} else {
		c.Display("gddddd", err)
		c.apiResult.Code = constant.FAIL
	}
	c.ServeJSON()
}

// @Title Get
// @Description get SunGoodsCommon by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SunGoodsCommon
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SunGoodsCommonController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSunGoodsCommonById(id)
	if err != nil {
		c.apiResult.Code = constant.InvalidAction
		c.apiResult.Error = err
	} else {
		c.apiResult.Code = constant.OK
		c.apiResult.Data = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description 根据种类查询商品，无其他筛选条件
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SunGoodsCommon
// @Failure 403
// @router / [get]
func (c *SunGoodsCommonController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0
	var classID = 0
	needSpec := false

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.apiResult.Code = constant.InvalidKeyValue
				c.ServeJSON()
				return
			}
			if kv[0] == "GcId" || kv[0] == "GcId1" || kv[0] == "GcId2" || kv[0] == "GcId3" {
				classID, _ = strconv.Atoi(kv[1])
				// 三级分类才回传规格信息，，，
				if kv[0] == "GcId2" || kv[0] == "GcId3" {
					needSpec = true
				}
			}
			query[kv[0]] = kv[1]
		}
	}
	//TODO:子分类带入
	l, totalCount, err := models.GetAllSunGoodsCommon(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.apiResult.Code = constant.FindNoData
		c.apiResult.Error = err

	} else {
		data := make(map[string]interface{}, 0)
		data["goodsCommon"] = l
		//TODO:可依据fields严格显示数据输出
		var fields []string
		ids := make([]int, len(l))

		//遍历找出所有goodsCommonId
		for index, item := range l {
			common, ok := item.(models.SunGoodsCommon)
			if ok {
				ids[index] = common.Id
			}
		}

		//获取goods
		goods, _ := models.GetAllSunGoodsByCommonID(ids, fields)
		data["goods"] = goods

		//带goods参数的查询即为筛选查询，不反馈brand、speces、specValues
		if classID != 0 {

			//根据goodsClassID获取所有子节点的ID
			childIDs, _ := models.GetGoodsClassChildList(classID)

			brands, _ := models.GetAllSunBrandByClassID(childIDs)
			data["brands"] = brands

			if needSpec {
				//获取specs
				specValues, _ := models.GetAllSpecValueByGoodsClassID(childIDs, fields)
				data["specValues"] = specValues

				var specIDS []int
				for _, item := range specValues {
					specValue := item.(models.SunSpecValue)
					specIDS = append(specIDS, int(specValue.SpId))
				}
				specIDS = tools.RemoveDuplicates(specIDS)
				specs, _ := models.GetAllSpecByIDS(specIDS, fields)
				data["speces"] = specs
			}

		}

		c.apiResult.Code = constant.OK
		c.apiResult.Data = data
		c.apiResult.Count = int(totalCount)
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the SunGoodsCommon
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SunGoodsCommon	true		"body for SunGoodsCommon content"
// @Success 200 {object} models.SunGoodsCommon
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SunGoodsCommonController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SunGoodsCommon{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSunGoodsCommonById(&v); err == nil {
			c.apiResult.Code = constant.OK
			c.apiResult.Data = v
		} else {
			c.apiResult.Code = constant.ActionFaild
			c.apiResult.Error = err
		}
	} else {
		c.apiResult.Code = constant.InvalidRequestData
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the SunGoodsCommon
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SunGoodsCommonController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSunGoodsCommon(id); err == nil {
		c.apiResult.Code = constant.OK
	} else {
		c.apiResult.Code = constant.ActionFaild
		c.apiResult.Error = err
	}
	c.ServeJSON()
}

// @Description 商品删除
// @router /delete_goods [post]
func (c *SunGoodsCommonController) DelGoods() {
	commonIdsStr := c.GetString("commonids")
	commonIdsStrSlice := strings.Split(commonIdsStr, ",")
	if len(commonIdsStr) <= 0 {
		c.HandleResult(nil, "删除商品id为空", constant.FAIL, nil, 0)
		return
	}
	var commonIds []int
	for _, commonId := range commonIdsStrSlice {
		temp, _ := strconv.Atoi(commonId)
		commonIds = append(commonIds, temp)
	}

	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}
	beego.Info("goodsCommonId", commonIds)
	result, err := models.DelGoodsNoLock(storeID, commonIds)

	if result {
		c.HandleResult(err, "删除商品成功", constant.OK, result, 0)
	} else {
		c.HandleResult(err, "删除商品失败", constant.FAIL, nil, 0)
	}

}

// @Description 商品下线
// @router /unshow_goods [post]
func (c *SunGoodsCommonController) UnshowGoods() {
	commonIdsStr := c.GetString("commonids")
	commonIdsStrSlice := strings.Split(commonIdsStr, ",")
	if len(commonIdsStr) <= 0 {
		c.HandleResult(nil, "下架商品id为空", constant.FAIL, nil, 0)
		return
	}
	var commonIds []int
	for _, commonId := range commonIdsStrSlice {
		temp, _ := strconv.Atoi(commonId)
		commonIds = append(commonIds, temp)
	}

	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}

	beego.Info("goodsCommonId", commonIds)
	result, err := models.EditProducesOffline(storeID, commonIds)

	if result {
		c.HandleResult(err, "商品下架成功", constant.OK, result, 0)
	} else {
		c.HandleResult(err, "商品下架失败", constant.FAIL, nil, 0)
	}
}

// @Description 商品上线
// @router /show_goods [post]
func (c *SunGoodsCommonController) ShowGoods() {
	commonIdsStr := c.GetString("commonids")
	commonIdsStrSlice := strings.Split(commonIdsStr, ",")
	if len(commonIdsStr) <= 0 {
		c.HandleResult(nil, "上架商品id为空", constant.FAIL, nil, 0)
		return
	}
	var commonIds []int
	for _, commonId := range commonIdsStrSlice {
		temp, _ := strconv.Atoi(commonId)
		commonIds = append(commonIds, temp)
	}

	storeID := c.GetCurrentStoreID()
	if storeID == 0 {
		c.HandleResult(nil, "未找到当前用户信息", constant.FAIL, nil, 0)
		return
	}

	beego.Info("goodsCommonId", commonIds)
	result, err := models.EditProducesOnline(storeID, commonIds)

	if result {
		c.HandleResult(err, "商品上架成功", constant.OK, result, 0)
	} else {
		c.HandleResult(err, "商品上架失败", constant.FAIL, nil, 0)
	}
}