package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"strconv"

	"github.com/astaxie/beego/orm"
)

type SunMember struct {
	Id                  int       `orm:"column(member_id);auto"`
	MemberName          string    `orm:"column(member_name);size(50)"`
	MemberTruename      string    `orm:"column(member_truename);size(20);null"`
	MemberAvatar        string    `orm:"column(member_avatar);size(50);null"`
	MemberSex           int8      `orm:"column(member_sex);null"`
	MemberBirthday      time.Time `orm:"column(member_birthday);type(date);null"`
	MemberPasswd        string    `orm:"column(member_passwd);size(32)"`
	MemberPaypwd        string    `orm:"column(member_paypwd);size(32);null"`
	MemberEmail         string    `orm:"column(member_email);size(100)"`
	MemberEmailBind     int8      `orm:"column(member_email_bind)"`
	MemberMobile        string    `orm:"column(member_mobile);size(11);null"`
	MemberMobileBind    int8      `orm:"column(member_mobile_bind)"`
	MemberQq            string    `orm:"column(member_qq);size(100);null"`
	MemberWw            string    `orm:"column(member_ww);size(100);null"`
	MemberLoginNum      int       `orm:"column(member_login_num)"`
	MemberTime          string    `orm:"column(member_time);size(10)"`
	MemberLoginTime     string    `orm:"column(member_login_time);size(10)"`
	MemberOldLoginTime  string    `orm:"column(member_old_login_time);size(10)"`
	MemberLoginIp       string    `orm:"column(member_login_ip);size(20);null"`
	MemberOldLoginIp    string    `orm:"column(member_old_login_ip);size(20);null"`
	MemberQqopenid      string    `orm:"column(member_qqopenid);size(100);null"`
	MemberQqinfo        string    `orm:"column(member_qqinfo);null"`
	MemberSinaopenid    string    `orm:"column(member_sinaopenid);size(100);null"`
	MemberSinainfo      string    `orm:"column(member_sinainfo);null"`
	MemberPoints        int       `orm:"column(member_points)"`
	AvailablePredeposit float64   `orm:"column(available_predeposit);digits(10);decimals(2)"`
	FreezePredeposit    float64   `orm:"column(freeze_predeposit);digits(10);decimals(2)"`
	AvailableRcBalance  float64   `orm:"column(available_rc_balance);digits(10);decimals(2)"`
	FreezeRcBalance     float64   `orm:"column(freeze_rc_balance);digits(10);decimals(2)"`
	InformAllow         int8      `orm:"column(inform_allow)"`
	IsBuy               int8      `orm:"column(is_buy)"`
	IsAllowtalk         int8      `orm:"column(is_allowtalk)"`
	MemberState         int8      `orm:"column(member_state)"`
	MemberSnsvisitnum   int       `orm:"column(member_snsvisitnum)"`
	MemberAreaid        int       `orm:"column(member_areaid);null"`
	MemberCityid        int       `orm:"column(member_cityid);null"`
	MemberProvinceid    int       `orm:"column(member_provinceid);null"`
	MemberAreainfo      string    `orm:"column(member_areainfo);size(255);null"`
	MemberPrivacy       string    `orm:"column(member_privacy);null"`
	MemberQuicklink     string    `orm:"column(member_quicklink);size(255);null"`
	MemberExppoints     int       `orm:"column(member_exppoints)"`
	InviterId           int       `orm:"column(inviter_id);null"`
	MemberAdminAdd      int8      `orm:"column(member_admin_add)"`
}

func (t *SunMember) TableName() string {
	return "sun_member"
}

func init() {
	orm.RegisterModel(new(SunMember))
}

// AddSunMember insert a new SunMember into database and returns
// last inserted Id on success.
func AddSunMember(m *SunMember) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunMemberById retrieves SunMember by Id. Returns error if
// Id doesn't exist
func GetSunMemberById(id int) (v *SunMember, err error) {
	o := orm.NewOrm()
	v = &SunMember{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunMember retrieves all SunMember matches certain condition. Returns empty list if
// no records exist
func GetAllSunMember(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunMember))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SunMember
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSunMember updates SunMember by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunMemberById(m *SunMember) (err error) {
	o := orm.NewOrm()
	v := SunMember{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunMember deletes SunMember by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunMember(id int) (err error) {
	o := orm.NewOrm()
	v := SunMember{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunMember{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//检查密码是否存在
func (t *SunMember) CheckMemberPassword() bool {
	o := Orm()
	return o.QueryTable(t.TableName()).Filter("member_id", t.Id).Filter("member_passwd", t.MemberPasswd).Exist()
}

//检查用户名是否存在
func (m *SunMember) CheckedName() bool {
	o := Orm()
	return o.QueryTable(m.TableName()).Exclude("member_id", m.Id).Filter("member_name", m.MemberName).Exist()
}

//检查邮箱是否存在
func (m *SunMember) CheckedEmail() bool {
	o := Orm()
	return o.QueryTable(m.TableName()).Exclude("member_id", m.Id).Filter("member_name", m.MemberName).Exist()
}

//CheckMemberLogin 会员登录操作，结构体带入数据：member_name,member_password,member_login_ip
func CheckMemberLogin(member SunMember) (SunMember, error) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	v := SunMember{MemberName: member.MemberName, MemberPasswd: member.MemberPasswd}
	// err := o.QueryTable(v).Filter("member_name", member.MemberName).Filter("member_passwd", member.MemberPasswd).One(&v)
	err := o.QueryTable(v).SetCond(cond.AndCond(cond.Or("member_name", member.MemberName).Or("member_mobile", member.MemberMobile)).AndCond(cond.And("member_passwd", member.MemberPasswd))).One(&v)
	if err == orm.ErrNoRows {
		return v, err
	}

	//TODO:时间和IP
	v.MemberLoginNum++
	v.MemberOldLoginIp = v.MemberLoginIp
	v.MemberOldLoginTime = v.MemberLoginTime
	v.MemberLoginTime = strconv.Itoa(int(time.Now().Unix()))
	v.MemberLoginIp = member.MemberLoginIp

	o = orm.NewOrm()
	o.Update(&v, "member_login_num", "member_login_time", "member_old_login_time",
		"member_login_ip", "member_old_login_ip")

	return v, nil
}

func (m *SunMember) IsAccountExist(account string) (exist bool) {

	o := orm.NewOrm()
	cond := orm.NewCondition()

	exist = o.QueryTable(m).SetCond(cond.And("MemberName", account).Or("MemberMobile", account)).Exist()
	return
}
