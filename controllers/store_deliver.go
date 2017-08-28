package controllers

import (
	"fmt"
	"strconv"

	"github.com/yakun0622/shop/constant"
	"github.com/yakun0622/shop/models"
)

type StoreDeliverController struct {
	BaseController
}

func (c *StoreDeliverController) URLMapping() {
	c.Mapping("Send", c.Send)
}

// @Title getSendInfo
// @Description 获取商户发货时所需的订单信息
// @Failure 403
// @router /send_info [post]
func (c *StoreDeliverController) Send() {
	var fields []string
	orderId, err := c.GetInt("order_id")
	storeID := c.GetCurrentStoreID()
	fmt.Println("当前商户ID：", storeID)
	if storeID == 0 {
		c.HandleResult(err, "未找到当前用户信息", constant.FAIL, nil, 0)
	}
	//获取订单信息
	orderInfo, err := models.GetSunOrderById(orderId)
	//获取商户发货地址
	var daddressInfo *models.SunDaddress
	if orderInfo.DaddressId > 0 {
		daddressInfo, _ = models.GetSunDaddressById(int(orderInfo.DaddressId))
	} else {
		var query map[string]string = make(map[string]string)
		sortby := []string{"IsDefault"}
		order := []string{"desc"}
		query["StoreId"] = strconv.Itoa(int(storeID))
		daddressList, err := models.GetAllSunDaddress(query, fields, sortby, order, 0, 1)
		if err != nil {
			c.HandleResult(err, "获取发货地址错误", constant.FAIL, nil, 0)
		}
		if len(daddressList) > 0 {
			daddressInfo = daddressList[0]
		}

	}
	//获取商户物流
	var my_express_list []models.SunStoreExtend
	my_express_list, err = models.GetStoreExtendByStoreId(storeID, fields)
	if err != nil {
		c.HandleResult(err, "获取商户物流数据错误", constant.FAIL, nil, 0)
	}
	//从缓存获取物流列表
	//var expressList []models.SunExpress
	expressList, err := models.GetAllSunExpress()

	var result map[string]interface{} = make(map[string]interface{})
	result["orderInfo"] = orderInfo
	result["daddressInfo"] = daddressInfo
	result["myExpressList"] = my_express_list
	result["expressList"] = expressList

	if err != nil {
		c.HandleResult(err, "获取商户物流数据错误", constant.FAIL, nil, 0)
	} else {
		c.HandleResult(err, "成功获取商户订单发货信息", constant.OK, result, 0)
	}
}

// @Title send add
// @Description 商户提交发货
// @Failure 403
// @router /send_add [post]
func (c *StoreDeliverController) SendAdd() {
	orderId, err := c.GetInt("order_id")
	storeID := c.GetCurrentStoreID()
	fmt.Println("当前商户ID：", storeID)
	if storeID == 0 {
		c.HandleResult(err, "未找到当前用户信息", constant.FAIL, nil, 0)
	}
	var sendInfo map[string]interface{} = make(map[string]interface{})
	//receiverInfo, err := getReceiverInfo(c)
	//sendInfo["receiverInfo"] = receiverInfo
	//sendInfo["reciverName"] = c.GetString("reciverName")
	sendInfo["deliverExplain"] = c.GetString("reciverName")
	sendInfo["daddressId"], err = c.GetInt32("daddressId")
	sendInfo["shippingExpressId"], err = c.GetInt8("shippingExpressId")
	sendInfo["shippingCode"] = c.GetString("shippingCode")
	sendInfo["invoiceNumber"] = c.GetString("invoiceNumber")
	sendInfo["invoiceCode"] = c.GetString("invoiceCode")
	if err != nil {
		c.HandleResult(err, "发货参数获取失败，请排查问题", constant.FAIL, nil, 0)
	}
	result := models.ChangeOrderSend(orderId, sendInfo)

	if result {
		c.HandleResult(err, "发货提交成功", constant.OK, result, 0)
	} else {
		c.HandleResult(err, "发货提交失败", constant.FAIL, nil, 0)
	}
}

//func getReceiverInfo(c *StoreDeliverController)  (receiverInfo string, err error){
//	var params map[string]interface{} = make(map[string]interface{})
//	params["address"] = c.GetString("reciverArea") + " " + c.GetString("reciverStreet")
//	params["phone"] = c.GetString("reciverMobPhone") + "," + c.GetString("reciverTelPhone")
//	params["area"] = c.GetString("reciverArea")
//	params["street"] = c.GetString("reciverStreet")
//	params["mob_phone"] = c.GetString("reciverMobPhone")
//	params["tel_phone"] = c.GetString("reciverTelPhone")
//	receiverInfoJson, err := json.Marshal(params)
//	receiverInfo = string(receiverInfoJson)
//	return receiverInfo, err
//}
