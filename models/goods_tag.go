package models

type GoodsTag struct {
	GoodsId uint   `orm:"column(goods_id)"`
	GroupId uint64 `orm:"column(group_id)"`
	TagId   uint64 `orm:"column(tag_id)"`
	TagName string `orm:"column(tag_name)"`
}

const GoodsTagTableName = "sun_goods_tag"

func GetTagsByGoodsId(goodsIds string, groupIds string) (l []GoodsTag, err error) {
	o, q := GetQueryBuilder()
	q = q.Select("*").From("sun_goods_tag as gt").InnerJoin("sun_tag as t").
		On("gt.tag_id = t.tag_id").Where("gt.goods_id in (" + goodsIds + ")")
	if groupIds != "" {
		q = q.And("t.group_id in (" + groupIds + ")")
	}
	_, err = o.Raw(q.String()).QueryRows(&l)

	if err != nil {
		Display("GetTagsByGoodsId", goodsIds)
		Display("GetTagsByGoodsId", groupIds)
	}
	return
}

func SaveGoodsTags(goodsAndTags []string) error {
	o := Orm()
	p, prepareErr := o.Raw("INSERT INTO sun_goods_tag (goods_id, tag_id) value (?, ?)").Prepare()
	if prepareErr == nil {
		var goodsId string
		for i, value := range goodsAndTags {
			if i%2 == 1 {
				_, err := p.Exec(goodsId, value)
				if err != nil {
					return err
				}
			} else {
				goodsId = value
			}
		}
	} else {
		return prepareErr
	}
	p.Close()
	return nil
}

func RemoveGoodsTag(goodsId int, tagId int) error {
	o:= Orm()

	num, err := o.Raw("DELETE FROM sun_goods_tag WHERE goods_id=? AND tag_id=?", goodsId, tagId).Exec()
	Display("RemoveGoodsTag", num, "goodsId", goodsId, "tagId", tagId)
	return err
}
