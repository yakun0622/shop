package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SunStoreJoinin struct {
	Id                                   int       `orm:"column(member_id);pk"`
	MemberName                           string    `orm:"column(member_name);size(50);null"`
	CompanyName                          string    `orm:"column(company_name);size(50);null"`
	CompanyProvinceId                    uint32    `orm:"column(company_province_id)"`
	CompanyAddress                       string    `orm:"column(company_address);size(50);null"`
	CompanyAddressDetail                 string    `orm:"column(company_address_detail);size(50);null"`
	CompanyPhone                         string    `orm:"column(company_phone);size(20);null"`
	CompanyEmployeeCount                 uint      `orm:"column(company_employee_count);null"`
	CompanyRegisteredCapital             uint      `orm:"column(company_registered_capital);null"`
	ContactsName                         string    `orm:"column(contacts_name);size(50);null"`
	ContactsPhone                        string    `orm:"column(contacts_phone);size(20);null"`
	ContactsEmail                        string    `orm:"column(contacts_email);size(50);null"`
	BusinessLicenceNumber                string    `orm:"column(business_licence_number);size(50);null"`
	BusinessLicenceAddress               string    `orm:"column(business_licence_address);size(50);null"`
	BusinessLicenceStart                 time.Time `orm:"column(business_licence_start);type(date);null"`
	BusinessLicenceEnd                   time.Time `orm:"column(business_licence_end);type(date);null"`
	BusinessSphere                       string    `orm:"column(business_sphere);size(1000);null"`
	BusinessLicenceNumberElectronic      string    `orm:"column(business_licence_number_electronic);size(50);null"`
	OrganizationCode                     string    `orm:"column(organization_code);size(50);null"`
	OrganizationCodeElectronic           string    `orm:"column(organization_code_electronic);size(50);null"`
	GeneralTaxpayer                      string    `orm:"column(general_taxpayer);size(50);null"`
	BankAccountName                      string    `orm:"column(bank_account_name);size(50);null"`
	BankAccountNumber                    string    `orm:"column(bank_account_number);size(50);null"`
	BankName                             string    `orm:"column(bank_name);size(50);null"`
	BankCode                             string    `orm:"column(bank_code);size(50);null"`
	BankAddress                          string    `orm:"column(bank_address);size(50);null"`
	BankLicenceElectronic                string    `orm:"column(bank_licence_electronic);size(50);null"`
	IsSettlementAccount                  int8      `orm:"column(is_settlement_account);null"`
	SettlementBankAccountName            string    `orm:"column(settlement_bank_account_name);size(50);null"`
	SettlementBankAccountNumber          string    `orm:"column(settlement_bank_account_number);size(50);null"`
	SettlementBankName                   string    `orm:"column(settlement_bank_name);size(50);null"`
	SettlementBankCode                   string    `orm:"column(settlement_bank_code);size(50);null"`
	SettlementBankAddress                string    `orm:"column(settlement_bank_address);size(50);null"`
	TaxRegistrationCertificate           string    `orm:"column(tax_registration_certificate);size(50);null"`
	TaxpayerId                           string    `orm:"column(taxpayer_id);size(50);null"`
	TaxRegistrationCertificateElectronic string    `orm:"column(tax_registration_certificate_electronic);size(50);null"`
	SellerName                           string    `orm:"column(seller_name);size(50);null"`
	StoreName                            string    `orm:"column(store_name);size(50);null"`
	StoreClassIds                        string    `orm:"column(store_class_ids);size(1000);null"`
	StoreClassNames                      string    `orm:"column(store_class_names);size(1000);null"`
	JoininState                          string    `orm:"column(joinin_state);size(50);null"`
	JoininMessage                        string    `orm:"column(joinin_message);size(200);null"`
	JoininYear                           uint8     `orm:"column(joinin_year)"`
	SgName                               string    `orm:"column(sg_name);size(50);null"`
	SgId                                 uint      `orm:"column(sg_id);null"`
	SgInfo                               string    `orm:"column(sg_info);size(200);null"`
	ScName                               string    `orm:"column(sc_name);size(50);null"`
	ScId                                 uint      `orm:"column(sc_id);null"`
	ScBail                               uint32    `orm:"column(sc_bail)"`
	StoreClassCommisRates                string    `orm:"column(store_class_commis_rates);size(200);null"`
	PayingMoneyCertificate               string    `orm:"column(paying_money_certificate);size(50);null"`
	PayingMoneyCertificateExplain        string    `orm:"column(paying_money_certificate_explain);size(200);null"`
	PayingAmount                         float64   `orm:"column(paying_amount);digits(10);decimals(2)"`
}

func (t *SunStoreJoinin) TableName() string {
	return "sun_store_joinin"
}

func init() {
	orm.RegisterModel(new(SunStoreJoinin))
}

// AddSunStoreJoinin insert a new SunStoreJoinin into database and returns
// last inserted Id on success.
func AddSunStoreJoinin(m *SunStoreJoinin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunStoreJoininById retrieves SunStoreJoinin by Id. Returns error if
// Id doesn't exist
func GetSunStoreJoininById(id int) (v *SunStoreJoinin, err error) {
	o := orm.NewOrm()
	v = &SunStoreJoinin{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunStoreJoinin retrieves all SunStoreJoinin matches certain condition. Returns empty list if
// no records exist
func GetAllSunStoreJoinin(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SunStoreJoinin))
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

	var l []SunStoreJoinin
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

// UpdateSunStoreJoinin updates SunStoreJoinin by Id and returns error if
// the record to be updated doesn't exist
func UpdateSunStoreJoininById(m *SunStoreJoinin) (err error) {
	o := orm.NewOrm()
	v := SunStoreJoinin{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunStoreJoinin deletes SunStoreJoinin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunStoreJoinin(id int) (err error) {
	o := orm.NewOrm()
	v := SunStoreJoinin{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SunStoreJoinin{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
