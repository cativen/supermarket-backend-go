package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type ExchangePointProductsDao interface {
	QueryPageByQoExchangePointProducts(dto dto.QueryExchangePointProductsRecordsDTO) vo.Page[model.ExchangePointProductsRecord]
	QueryOptionsMemberPhone() []int64
	DelByCn(cn string)
	SaveData(products model.ExchangePointProductsRecord)
}

type ExchangePointProductsDaoImpl struct {
}

func (exchangePointProductsDao ExchangePointProductsDaoImpl) QueryPageByQoExchangePointProducts(dto dto.QueryExchangePointProductsRecordsDTO) vo.Page[model.ExchangePointProductsRecord] {
	var page vo.Page[model.ExchangePointProductsRecord]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("exchange_point_products_records").Select("*").Where("state = ?", common.STATE_NORMAL)

	memberId := dto.MemberId
	if memberId != 0 {
		db = db.Where("member_id = ?", memberId)
	}

	cnString := dto.Cn
	if cnString != "" && len(cnString) > 0 {
		db = db.Where("cn LIKE ?", "%"+cnString+"%")
	}

	startTime := dto.StartTime
	if startTime != "" && len(startTime) > 0 {
		db = db.Where("update_time >= ?", startTime)
	}

	endTime := dto.EndTime
	if endTime != "" && len(endTime) > 0 {
		db = db.Where("update_time <= ?", endTime)
	}
	var records []model.ExchangePointProductsRecord
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (exchangePointProductsDao ExchangePointProductsDaoImpl) QueryOptionsMemberPhone() []int64 {
	db := common.GetDB()
	var list []model.ExchangePointProductsRecord
	db.Select("member_id").Where("state = ?", common.STATE_NORMAL).Group("member_id").Find(&list)
	var memberIds = make([]int64, len(list))
	for index, v := range list {
		memberIds[index] = v.MemberID
	}
	return memberIds
}

func (exchangePointProductsDao ExchangePointProductsDaoImpl) DelByCn(cn string) {
	db := common.GetDB()
	update := &model.ExchangePointProductsRecord{
		State: common.STATE_DEL,
	}
	db.Where("cn = ?", cn).Updates(update)
}

func (exchangePointProductsDao ExchangePointProductsDaoImpl) SaveData(products model.ExchangePointProductsRecord) {
	db := common.GetDB()
	db.Create(&products)
}
