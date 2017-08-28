package constant

/**
 *  订单状态
 */
const (
	//已取消
	ORDER_STATE_CANCEL = 0
	//已产生但未审核
	ORDER_STATE_NEW = 10
	//审核未通过
	ORDER_STATE_APPROVE_NOT_PASS = 20

	//已支付，审批通过，待发货
	ORDER_STATE_APPROVE_PASS = 30
	//已发货
	ORDER_STATE_SEND = 40
	//已收货，交易成功
	ORDER_STATE_SUCCESS = 50
	//未付款订单，自动取消的天数
	ORDER_AUTO_CANCEL_DAY = 3
	//已发货订单，自动确认收货的天数
	ORDER_AUTO_RECEIVE_DAY = 7
	//兑换码支持过期退款，可退款的期限，默认为7天
	CODE_INVALID_REFUND = 7
	//默认未删除
	ORDER_DEL_STATE_DEFAULT = 0
	//已删除
	ORDER_DEL_STATE_DELETE = 1
	//彻底删除
	ORDER_DEL_STATE_DROP = 2
	//订单结束后可评论时间，15天，60*60*24*15
	ORDER_EVALUATE_TIME = 1296000
	//抢购订单状态
	OFFLINE_ORDER_CANCEL_TIME = 3//单位为天
	//订单待审批状态
	ORDER_WAITING_APPROVE = 0
	//订单审批通过状态
	ORDER_PASS_APPROVE = 0
	//订单审批驳回状态
	ORDER_REFUSE_APPROVE = 0
)