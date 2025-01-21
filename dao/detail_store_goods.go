package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type DetailStoreGoodsDao interface {
	SelectPageByQoDetailStoreGoods(dto dto.QueryDetailStoreGoodsOutDTO, storeTypes string) vo.Page[model.DetailStoreGoods]
	UpdateStoreGoods(goods model.DetailStoreGoods)
	SaveData(detail model.DetailStoreGoods)
	DeleteByQo(goods model.DetailStoreGoods)
	SelectListByQo(queryModel model.DetailStoreGoods) []model.DetailStoreGoods
}

type DetailStoreGoodsDaoImpl struct {
}

func (i DetailStoreGoodsDaoImpl) DeleteByQo(goods model.DetailStoreGoods) {
	db := common.GetDB()
	db.Delete(&goods)
}

func (DetailStoreGoodsDaoImpl) SelectPageByQoDetailStoreGoods(dto dto.QueryDetailStoreGoodsOutDTO, storeTypes string) vo.Page[model.DetailStoreGoods] {
	var page vo.Page[model.DetailStoreGoods]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("t_detail_store_goods").Select("*").Where("type = ?", storeTypes)
	cn := dto.Cn
	if cn != "" && len(cn) > 0 {
		db = db.Where("cn like ?", cn+"%")
	}

	goodName := dto.GoodsName
	if goodName != "" && len(goodName) > 0 {
		db = db.Where("goods_name LIKE ?", "%"+goodName+"%")
	}

	StartCreateTime := dto.StartCreateTime
	if StartCreateTime != "" && len(StartCreateTime) > 0 {
		db = db.Where("create_time >= ?", StartCreateTime)
	}

	endCreateTime := dto.EndCreateTime
	if endCreateTime != "" && len(endCreateTime) > 0 {
		db = db.Where("create_time <= ?", endCreateTime)
	}

	state := dto.State
	if state != "" && len(state) > 0 {
		db = db.Where("state = ?", state)
	}

	state1 := dto.State1
	if state1 != "" && len(state1) > 0 {
		db = db.Where("state1 = ?", state1)
	}

	// 查询出库记录
	var records []model.DetailStoreGoods
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (DetailStoreGoodsDaoImpl) UpdateStoreGoods(goods model.DetailStoreGoods) {
	db := common.GetDB()
	db.Updates(&goods)
}

func (DetailStoreGoodsDaoImpl) SaveData(detail model.DetailStoreGoods) {
	db := common.GetDB()
	db.Create(&detail)
}

func (DetailStoreGoodsDaoImpl) SelectListByQo(queryModel model.DetailStoreGoods) []model.DetailStoreGoods {
	db := common.GetDB()
	var records []model.DetailStoreGoods
	db.Where(&queryModel).Find(&records)
	return records
}
