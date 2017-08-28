package models

type SunStoreMsgSetting struct {
	SmtCode          string `orm:"column(smt_code);size(100)"`
	StoreId          uint   `orm:"column(store_id)"`
	SmsMessageSwitch uint8  `orm:"column(sms_message_switch)"`
	SmsShortSwitch   uint8  `orm:"column(sms_short_switch)"`
	SmsMailSwitch    uint8  `orm:"column(sms_mail_switch)"`
	SmsShortNumber   string `orm:"column(sms_short_number);size(11)"`
	SmsMailNumber    string `orm:"column(sms_mail_number);size(100)"`
}
