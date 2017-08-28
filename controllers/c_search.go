package controllers

import (
	"strconv"
	"strings"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
	"github.com/yakun0622/shop/tools"
)

//SearchControllerController search
type SearchControllerController struct {
	BaseController
}

func (c *SearchControllerController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetFilter", c.GetFilter)
}

// @Title Get All
// @Description 搜索商品
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	keyword	query	string	false	"Filter. e.g. 电脑 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Failure 403
// @router /goods_common/ [post]
func (c *SearchControllerController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 20
	var offset int64
	var classID = 0
	keywords := c.GetString("keyword")
	needSpec := false

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
			k, v := kv[0], kv[1]
			if kv[0] == "GcId" || kv[0] == "GcId1" || kv[0] == "GcId2" || kv[0] == "GcId3" {
				classID, _ = strconv.Atoi(kv[1])
				if kv[0] == "GcId3" {
					needSpec = true
				}
			}
			query[k] = v
		}
	}

	if keywords != "" {
		query["keywords"] = keywords
	}
	query["GoodsState"] = "1"
	query["GoodsVerify"] = "1"
	query["GoodsLock"] = "0"

	l, totalCount, err := models.GetAllSunGoodsCommon(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.apiResult.Code = constant.OK
		c.apiResult.Error = err
	} else {
		data := make(map[string]interface{}, 0)
		data["goodsCommon"] = l
		//可依据fields严格显示数据输出
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

// @Title Get Filter
// @Description 根据规格、品牌等进行商品筛选
// @Param	keyword	query	string	false	"Filter. e.g. 电脑 ..."
// @Param	price	query	string	false	"Filter. e.g. ~100,50~100,100~ ..."
// @Param	spec	query	string	false	"Filter. e.g. key:value,key:value"
// @Param	brands	query	string	false	"Filter. e.g. val,val,val,val"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Failure 403
// @router /goods/filter [post]
func (c *SearchControllerController) GetFilter() {
	var order []string
	var sortby []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 20
	var offset int64

	keyword := c.GetString("keyword")
	price := c.GetString("price")
	specValues := make(map[string]string)
	var brands []int

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	//规格值筛选: k:v,k:v
	if v := c.GetString("spec"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.apiResult.Code = constant.InvalidKeyValue
				c.ServeJSON()
				return
			}
			specValues[kv[0]] = kv[1]
		}
	}
	//brands: val,val,val
	if v := c.GetString("brands"); v != "" {
		brandArray := strings.Split(v, ",")
		for _, val := range brandArray {
			brandID, _ := strconv.Atoi(val)
			brands = append(brands, brandID)
		}
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
			query[kv[0]] = kv[1]
		}
	}
	l, err := models.FilterSunGoodsCommon(keyword, query, specValues, brands, price, sortby, order, offset, limit)
	totalCount, err2 := models.FilterSunGoodsCommonAllCount(keyword, query, specValues, brands, price)
	if err != nil || err2 != nil || len(l) <= 0 {
		c.apiResult.Code = constant.OK
		if err != nil {
			c.apiResult.Error = err
		}
		if err2 != nil {
			c.apiResult.Error = err2
		}
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

		c.apiResult.Code = constant.OK
		c.apiResult.Data = data
		c.apiResult.Count = int(totalCount)
	}
	c.ServeJSON()
}
