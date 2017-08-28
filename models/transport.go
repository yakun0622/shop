package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Transport struct {
	Id         int    `orm:"column(id);auto"`
	Title      string `orm:"column(title);size(30)"`
	SendTplId  uint32 `orm:"column(send_tpl_id);null"`
	StoreId    uint32 `orm:"column(store_id)"`
	UpdateTime uint   `orm:"column(update_time);null"`

	Extends []TransportExtend `orm:"-"`
}

const TransportTableName = "shop_transport"

func (t *Transport) TableName() string {
	return "shop_transport"
}

func init() {
	orm.RegisterModel(new(Transport))
}

// AddSunTransport insert a new Transport into database and returns
// last inserted Id on success.
func AddTransport(m *Transport) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSunTransportById retrieves Transport by Id. Returns error if
// Id doesn't exist
func GetSunTransportById(id int) (v *Transport, err error) {
	o := orm.NewOrm()
	v = &Transport{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSunTransport retrieves all Transport matches certain condition. Returns empty list if
// no records exist
func GetTransportListByStoreId(store_id uint) (list []Transport, count int64, err error) {
	o, q := GetQueryBuilder()
	q = q.Select("*").From(TransportTableName).Where("store_id = " + strconv.Itoa(int(store_id))).OrderBy("id desc")
	sql := q.String()
	count, err = o.Raw(sql).QueryRows(&list)
	return list, count, err

}

// UpdateSunTransport updates Transport by Id and returns error if
// the record to be updated doesn't exist
func UpdateTransportById(m *Transport) (err error) {
	o := orm.NewOrm()
	v := Transport{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSunTransport deletes Transport by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSunTransport(id int) (err error) {
	o := orm.NewOrm()
	v := Transport{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Transport{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetTransportfee(goodses string, goodsNums []int, amount float64, areaId int) (transportfee float64) {
	o, goodsQuery := GetQueryBuilder()
	var transportIds orm.ParamsList
	goodsQuery.Select("transport_id").From("shop_goods").
		Where("goods_id in (" + goodses + ")").OrderBy("transport_id")

	_, err := o.Raw(goodsQuery.String()).ValuesFlat(&transportIds)
	if err != nil {
		return
	}

	areaIdStr := "%" + strconv.Itoa(areaId) + "%"


	q := QueryBuilder()
	q.Select("te.free_price").From("shop_transport as t").
		InnerJoin("shop_transport_extend as te").On("t.id = te.transport_id").
		Where("t.id=?").And("(te.area_id like '"+ areaIdStr + "' OR te.area_name='全国')").And("te.free_line > ?")
	sql := q.String()

	var feelLine float64
	var transportIdTemp string
	var feelLineTemp float64
	for i, transportId := range transportIds {
		transportId := transportId.(string)
		if transportIdTemp != transportId {
			err := o.Raw(sql, transportId, amount).QueryRow(&feelLine)
			if err != nil {
				beego.Error("GetTransportfee---", transportIds)
			}
			if feelLine != 0 {
				transportfee += feelLine * float64(goodsNums[i])
			}
			feelLineTemp = feelLine
			transportIdTemp = transportId
		} else {
			transportfee += feelLineTemp * float64(goodsNums[i])
		}
	}
	return
}
