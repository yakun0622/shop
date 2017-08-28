package permission

const (
	ORGANIZATIONAL = 1 << iota
	MEMBER
	APPROVE
	EMERGENCY
	ORDER
	WELFARE
	SETTLEMENT

)

var Permission = map[uint64]string{
	ORGANIZATIONAL: "组织管理",
	MEMBER:         "成员管理",
	APPROVE:        "审批",
	EMERGENCY:      "应急审批",
	ORDER:          "订单管理",
	WELFARE:        "福利审批",
	SETTLEMENT:	"结算审批",
}
