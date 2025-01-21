package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type SaleRecordsDao interface {
	QueryPageByQoSaleRecords(dto dto.SaleRecordPageDTO) vo.Page[model.SaleRecord]
	DelSaleRecords(cn string)
	SaveRecords(saleRecords model.SaleRecord)
}

type SaleRecordsDaoImpl struct {
}

func (saleRecordsDao SaleRecordsDaoImpl) QueryPageByQoSaleRecords(dto dto.SaleRecordPageDTO) vo.Page[model.SaleRecord] {
	var page vo.Page[model.SaleRecord]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("t_sale_records").Select("*").Where("state = ?", common.STATE_NORMAL)
	typeString := dto.Type
	if typeString != "" && len(typeString) > 0 {
		db = db.Where("type = ?", typeString)
	}

	cnString := dto.Cn
	if cnString != "" && len(cnString) > 0 {
		db = db.Where("cn LIKE ?", "%"+cnString+"%")
	}

	startSellTime := dto.StartSellTime
	if startSellTime != "" && len(startSellTime) > 0 {
		db = db.Where("sell_time >= ?", startSellTime)
	}

	endSellTime := dto.EndSellTime
	if endSellTime != "" && len(endSellTime) > 0 {
		db = db.Where("sell_time <= ?", endSellTime)
	}

	sellway := dto.Sellway
	if sellway != "" && len(sellway) > 0 {
		db = db.Where("sellway = ?", sellway)
	}

	// 查询销售记录
	var saleRecords []model.SaleRecord
	count, saleRecords := utils.Paginate(db, dto.CurrentPage, dto.PageSize, saleRecords)

	for index, record := range saleRecords {
		var details []model.DetailSaleRecord
		db := common.GetDB()
		db.Where("sell_cn = ?", record.CN).Find(&details)
		saleRecords[index].DetailSaleRecords = details
		saleRecords[index].SellTimeResp = utils.FormatTimeToString(saleRecords[index].SellTime)
	}

	page.Records = saleRecords
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (saleRecordsDao SaleRecordsDaoImpl) DelSaleRecords(cn string) {
	db := common.GetDB()
	db.Where("cn = ?", cn).Update("state", common.STATE_DEL)
}

func (saleRecordsDao SaleRecordsDaoImpl) SaveRecords(saleRecords model.SaleRecord) {
	db := common.GetDB()
	db.Create(&saleRecords)
}
