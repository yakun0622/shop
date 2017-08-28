// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"github.com/yakun0622/shop/controllers"
	"github.com/yakun0622/shop/controllers/store"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/address",
			beego.NSInclude(
				&controllers.SunAddressController{},
			),
		),

		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.SunAdminController{},
			),
		),

		beego.NSNamespace("/admin_log",
			beego.NSInclude(
				&controllers.SunAdminLogController{},
			),
		),

		beego.NSNamespace("/album_class",
			beego.NSInclude(
				&controllers.SunAlbumClassController{},
			),
		),

		beego.NSNamespace("/album_pic",
			beego.NSInclude(
				&controllers.SunAlbumPicController{},
			),
		),

		beego.NSNamespace("/approve_order",
			beego.NSInclude(
				&controllers.SunApproveOrderController{},
			),
		),

		beego.NSNamespace("/area",
			beego.NSInclude(
				&controllers.SunAreaController{},
			),
		),

		beego.NSNamespace("/arrival_notice",
			beego.NSInclude(
				&controllers.SunArrivalNoticeController{},
			),
		),

		beego.NSNamespace("/attribute",
			beego.NSInclude(
				&controllers.SunAttributeController{},
			),
		),

		beego.NSNamespace("/attribute_value",
			beego.NSInclude(
				&controllers.SunAttributeValueController{},
			),
		),

		beego.NSNamespace("/brand",
			beego.NSInclude(
				&controllers.BrandController{},
			),
		),

		beego.NSNamespace("/cart",
			beego.NSInclude(
				&controllers.SunCartController{},
			),
		),

		beego.NSNamespace("/daddress",
			beego.NSInclude(
				&controllers.SunDaddressController{},
			),
		),

		beego.NSNamespace("/evaluate_goods",
			beego.NSInclude(
				&controllers.SunEvaluateGoodsController{},
			),
		),

		beego.NSNamespace("/evaluate_store",
			beego.NSInclude(
				&controllers.SunEvaluateStoreController{},
			),
		),

		beego.NSNamespace("/express",
			beego.NSInclude(
				&controllers.SunExpressController{},
			),
		),

		beego.NSNamespace("/favorites",
			beego.NSInclude(
				&controllers.SunFavoritesController{},
			),
		),

		beego.NSNamespace("/favorites_folder",
			beego.NSInclude(
				&controllers.SunFavoritesFolderController{},
			),
		),

		beego.NSNamespace("/gadmin",
			beego.NSInclude(
				&controllers.SunGadminController{},
			),
		),

		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.SunGoodsController{},
			),
		),

		beego.NSNamespace("/goods_class",
			beego.NSInclude(
				&controllers.SunGoodsClassController{},
			),
		),

		beego.NSNamespace("/goods_class_staple",
			beego.NSInclude(
				&controllers.SunGoodsClassStapleController{},
			),
		),

		beego.NSNamespace("/goods_class_tag",
			beego.NSInclude(
				&controllers.SunGoodsClassTagController{},
			),
		),

		beego.NSNamespace("/goods_common",
			beego.NSInclude(
				&controllers.SunGoodsCommonController{},
			),
		),

		beego.NSNamespace("/goods_images",
			beego.NSInclude(
				&controllers.SunGoodsImagesController{},
			),
		),

		beego.NSNamespace("/group",
			beego.NSInclude(
				&controllers.SunGroupController{},
			),
		),

		beego.NSNamespace("/invoice",
			beego.NSInclude(
				&controllers.SunInvoiceController{},
			),
		),

		beego.NSNamespace("/member",
			beego.NSInclude(
				&controllers.SunMemberController{},
			),
		),

		beego.NSNamespace("/member_common",
			beego.NSInclude(
				&controllers.SunMemberCommonController{},
			),
		),
		beego.NSNamespace("/member_ext",
			beego.NSInclude(
				&controllers.SunMemberExtController{},
			),
		),

		beego.NSNamespace("/member_msg_tpl",
			beego.NSInclude(
				&controllers.SunMemberMsgTplController{},
			),
		),

		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.SunMessageController{},
			),
		),

		beego.NSNamespace("/offpay_area",
			beego.NSInclude(
				&controllers.SunOffpayAreaController{},
			),
		),

		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.SunOrderController{},
			),
		),

		beego.NSNamespace("/order_bill",
			beego.NSInclude(
				&controllers.SunOrderBillController{},
			),
		),

		beego.NSNamespace("/order_bill_group",
			beego.NSInclude(
				&controllers.SunOrderBillGroupController{},
			),
		),

		beego.NSNamespace("/order_common",
			beego.NSInclude(
				&controllers.SunOrderCommonController{},
			),
		),

		beego.NSNamespace("/order_goods",
			beego.NSInclude(
				&controllers.SunOrderGoodsController{},
			),
		),

		beego.NSNamespace("/order_log",
			beego.NSInclude(
				&controllers.SunOrderLogController{},
			),
		),

		beego.NSNamespace("/order_pay",
			beego.NSInclude(
				&controllers.SunOrderPayController{},
			),
		),

		beego.NSNamespace("/order_statis",
			beego.NSInclude(
				&controllers.SunOrderStatisController{},
			),
		),

		beego.NSNamespace("/order_statis_group",
			beego.NSInclude(
				&controllers.SunOrderStatisGroupController{},
			),
		),

		beego.NSNamespace("/payment",
			beego.NSInclude(
				&controllers.SunPaymentController{},
			),
		),

		beego.NSNamespace("/refund_reason",
			beego.NSInclude(
				&controllers.SunRefundReasonController{},
			),
		),

		//beego.NSNamespace("/refund_return",
		//	beego.NSInclude(
		//		&controllers.SunRefundReturnController{},
		//	),
		//),

		beego.NSNamespace("/seller",
			beego.NSInclude(
				&controllers.SunSellerController{},
			),
		),

		beego.NSNamespace("/seller_group",
			beego.NSInclude(
				&controllers.SunSellerGroupController{},
			),
		),

		beego.NSNamespace("/seller_log",
			beego.NSInclude(
				&controllers.SunSellerLogController{},
			),
		),

		beego.NSNamespace("/setting",
			beego.NSInclude(
				&controllers.SunSettingController{},
			),
		),

		beego.NSNamespace("/spec",
			beego.NSInclude(
				&controllers.SunSpecController{},
			),
		),

		beego.NSNamespace("/spec_value",
			beego.NSInclude(
				&controllers.SunSpecValueController{},
			),
		),

		beego.NSNamespace("/store",
			beego.NSInclude(
				&controllers.SunStoreController{},
			),
		),

		beego.NSNamespace("/store_bind_class",
			beego.NSInclude(
				&controllers.SunStoreBindClassController{},
			),
		),

		beego.NSNamespace("/store_class",
			beego.NSInclude(
				&controllers.SunStoreClassController{},
			),
		),

		beego.NSNamespace("/store_extend",
			beego.NSInclude(
				&controllers.SunStoreExtendController{},
			),
		),

		beego.NSNamespace("/store_goods_class",
			beego.NSInclude(
				&controllers.SunStoreGoodsClassController{},
			),
		),

		beego.NSNamespace("/store_grade",
			beego.NSInclude(
				&controllers.SunStoreGradeController{},
			),
		),

		beego.NSNamespace("/store_joinin",
			beego.NSInclude(
				&controllers.SunStoreJoininController{},
			),
		),

		beego.NSNamespace("/store_msg",
			beego.NSInclude(
				&controllers.SunStoreMsgController{},
			),
		),

		beego.NSNamespace("/store_msg_tpl",
			beego.NSInclude(
				&controllers.SunStoreMsgTplController{},
			),
		),

		beego.NSNamespace("/store_sns_comment",
			beego.NSInclude(
				&controllers.SunStoreSnsCommentController{},
			),
		),

		beego.NSNamespace("/store_watermark",
			beego.NSInclude(
				&controllers.SunStoreWatermarkController{},
			),
		),

		beego.NSNamespace("/store_waybill",
			beego.NSInclude(
				&controllers.SunStoreWaybillController{},
			),
		),

		beego.NSNamespace("/transport",
			beego.NSInclude(
				&controllers.TransportController{},
			),
		),

		beego.NSNamespace("/transport_extend",
			beego.NSInclude(
				&controllers.TransportExtendController{},
			),
		),

		beego.NSNamespace("/type",
			beego.NSInclude(
				&controllers.SunTypeController{},
			),
		),

		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.SunUploadController{},
			),
		),

		beego.NSNamespace("/waybill",
			beego.NSInclude(
				&controllers.SunWaybillController{},
			),
		),

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.MemberLoginController{},
			),
		),

		beego.NSNamespace("/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),

		beego.NSNamespace("/register",
			beego.NSInclude(
				&controllers.RegisterControllerController{},
			),
		),

		beego.NSNamespace("/sms/code",
			beego.NSInclude(
				&controllers.SMSController{},
			),
		),

		beego.NSNamespace("/adminlogin",
			beego.NSInclude(
				&controllers.AdminLoginController{},
			),
		),

		beego.NSNamespace("/search",
			beego.NSInclude(
				&controllers.SearchControllerController{},
			),
		),

		beego.NSNamespace("/store/orders",
			beego.NSInclude(
				&controllers.StoreOrdersController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.Role{},
			),
		),

		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.Tag{},
			),
		),

		beego.NSNamespace("/member_group",
			beego.NSInclude(
				&controllers.MemberGroup{},
			),
		),

		beego.NSNamespace("/goods_tag",
			beego.NSInclude(
				&controllers.GoodsTag{},
			),
		),

		beego.NSNamespace("/store_deliver",
			beego.NSInclude(
				&controllers.StoreDeliverController{},
			),
		),
		beego.NSNamespace("/store/refund_return",
			beego.NSInclude(
				&store.RefundReturnController{},
			),
		),
	)
	beego.AddNamespace(ns)

	//自定义router
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/captcha", &controllers.CaptchaController{})
	beego.Handler("/captcha/*.png", captcha.Server(captcha.StdWidth, captcha.StdHeight))

	//错误处理
	beego.ErrorController(&controllers.ErrorController{})
}
