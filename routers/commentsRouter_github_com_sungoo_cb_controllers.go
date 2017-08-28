package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:AdminLoginController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:AdminLoginController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:BrandController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:BrandController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:GoodsTag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:GoodsTag"],
		beego.ControllerComments{
			"GetTagsByGoodsId",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:GoodsTag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:GoodsTag"],
		beego.ControllerComments{
			"Save",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:GoodsTag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:GoodsTag"],
		beego.ControllerComments{
			"Remove",
			`/remove`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:LogoutController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"],
		beego.ControllerComments{
			"Apple",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"],
		beego.ControllerComments{
			"Pass",
			`/`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"],
		beego.ControllerComments{
			"AddRole",
			`/addrole`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"],
		beego.ControllerComments{
			"ChangeRole",
			`/changerole`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"],
		beego.ControllerComments{
			"Delete",
			`/`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberGroup"],
		beego.ControllerComments{
			"Import",
			`/import`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberLoginController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberLoginController"],
		beego.ControllerComments{
			"Exist",
			`/exist`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberLoginController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:MemberLoginController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:RegisterControllerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:RegisterControllerController"],
		beego.ControllerComments{
			"GetSMS",
			`/sms/:mobile`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:RegisterControllerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:RegisterControllerController"],
		beego.ControllerComments{
			"NoExist",
			`/noexist`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:RegisterControllerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:RegisterControllerController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"],
		beego.ControllerComments{
			"Permission",
			`/permission/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"],
		beego.ControllerComments{
			"Put",
			`/`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"],
		beego.ControllerComments{
			"Lock",
			`/lock/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Role"],
		beego.ControllerComments{
			"UnLock",
			`/unlock/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SMSController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SMSController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SearchControllerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SearchControllerController"],
		beego.ControllerComments{
			"GetAll",
			`/goods_common/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SearchControllerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SearchControllerController"],
		beego.ControllerComments{
			"GetFilter",
			`/goods/filter`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreDeliverController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreDeliverController"],
		beego.ControllerComments{
			"Send",
			`/send_info`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreDeliverController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreDeliverController"],
		beego.ControllerComments{
			"SendAdd",
			`/send_add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"],
		beego.ControllerComments{
			"StoreOrders",
			`/list`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"],
		beego.ControllerComments{
			"GetOrderInfo",
			`/orderInfo`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"],
		beego.ControllerComments{
			"EditShippingFee",
			`/edit_shippingFee`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"],
		beego.ControllerComments{
			"ChangeState",
			`/change_state/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:StoreOrdersController"],
		beego.ControllerComments{
			"GetOrderGoodes",
			`/goodses`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAddressController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAdminLogController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumClassController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAlbumPicController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunApproveOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunApproveOrderController"],
		beego.ControllerComments{
			"Approve",
			`/approve`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunApproveOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunApproveOrderController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAreaController"],
		beego.ControllerComments{
			"GetAreaList",
			`/get_area`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunArrivalNoticeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunAttributeValueController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunBrandController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"],
		beego.ControllerComments{
			"Put",
			`/change_num`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunCartController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunDaddressController"],
		beego.ControllerComments{
			"SetDefault",
			`/default/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"],
		beego.ControllerComments{
			"GetEvaluate",
			`/getevaluate`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateGoodsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunEvaluateStoreController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunExpressController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"],
		beego.ControllerComments{
			"Put",
			`/change_num`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesController"],
		beego.ControllerComments{
			"Delete",
			`/remove`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"],
		beego.ControllerComments{
			"Put",
			`/`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunFavoritesFolderController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGadminController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassStapleController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsClassTagController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"],
		beego.ControllerComments{
			"DelGoods",
			`/delete_goods`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"],
		beego.ControllerComments{
			"UnshowGoods",
			`/unshow_goods`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsCommonController"],
		beego.ControllerComments{
			"ShowGoods",
			`/show_goods`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"],
		beego.ControllerComments{
			"Details",
			`/details`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsImagesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsImagesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsImagesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsImagesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsImagesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsImagesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsImagesController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGoodsImagesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"],
		beego.ControllerComments{
			"GetOwner",
			`/owner`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"],
		beego.ControllerComments{
			"Search",
			`/search`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunGroupController"],
		beego.ControllerComments{
			"ChangeBelong",
			`/changeBelong`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunInvoiceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberCommonController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"Password",
			`/password`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"NameNoExist",
			`/name`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"EmailNoExist",
			`/email`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"Put",
			`/`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberController"],
		beego.ControllerComments{
			"GetAllMemberTree",
			`/membertree`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberExtController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMemberMsgTplController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunMessageController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOffpayAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOffpayAreaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOffpayAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOffpayAreaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOffpayAreaController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOffpayAreaController"],
		beego.ControllerComments{
			"Save",
			`/save`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderBillGroupController"],
		beego.ControllerComments{
			"GetList",
			`/list`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderCommonController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"User",
			`/user`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"GetGoodesAndApprovers",
			`/goodses`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"GroupOrders",
			`/get`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"ChangeState",
			`/change_state`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"GetSettlement",
			`/get_settlement`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderController"],
		beego.ControllerComments{
			"GetTransportfee",
			`/transportfee`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderGoodsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderLogController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderPayController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunOrderStatisGroupController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunPaymentController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReasonController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReturnController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReturnController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReturnController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReturnController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReturnController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReturnController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReturnController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunRefundReturnController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerGroupController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSellerLogController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSettingController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunSpecValueController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreBindClassController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreClassController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"GetSimpleStatic",
			`/static/simple`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"GetGoodsCommonList",
			`/goods_common/online`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"GetGoodsCommonOfflineList",
			`/goods_common/offline`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreController"],
		beego.ControllerComments{
			"SetFreeFreight",
			`/free_freight`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreExtendController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGoodsClassController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreGradeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreJoininController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreMsgTplController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreSnsCommentController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWatermarkController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunStoreWaybillController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunTypeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunUploadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:SunWaybillController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"GetUserGroupTags",
			`/user_group_tags`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"GetShare",
			`/share`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"Share",
			`/share`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"UnShare",
			`/unshare`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"GetRoleTag",
			`/role`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"RoleTags",
			`/role`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"RemoveRoleTags",
			`/role`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"Put",
			`/`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"Lock",
			`/lock/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:Tag"],
		beego.ControllerComments{
			"UnLock",
			`/unlock/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"],
		beego.ControllerComments{
			"AddTransport",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"],
		beego.ControllerComments{
			"GetList",
			`/list`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportExtendController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportExtendController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportExtendController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportExtendController"],
		beego.ControllerComments{
			"Save",
			`/save`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportExtendController"] = append(beego.GlobalControllerRouter["github.com/yakun0622/shop/controllers:TransportExtendController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

}
