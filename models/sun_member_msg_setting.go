package models

type SunMemberMsgSetting struct {
	MmtCode   string `orm:"column(mmt_code);size(50)"`
	MemberId  uint   `orm:"column(member_id)"`
	IsReceive uint8  `orm:"column(is_receive)"`
}
