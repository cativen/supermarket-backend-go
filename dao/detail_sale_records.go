package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/model"
)

type DetailSaleRecordsServiceDao interface {
	SaveDetailSaleRecords(detailSaleRecords []model.DetailSaleRecord)
}

type DetailSaleRecordsServiceImpl struct {
}

func (DetailSaleRecordsServiceImpl) SaveDetailSaleRecords(detailSaleRecords []model.DetailSaleRecord) {
	db := common.GetDB()
	db.Create(detailSaleRecords)
}
