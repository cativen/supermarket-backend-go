package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type GoodsStoreDao interface {
	StoreUsed(storeId uint) int64
	SelectGoodListWithResidueNum() []model.GoodsStore
	SelectByQo(store model.GoodsStore) model.GoodsStore
	Save(store model.GoodsStore)
	UpdateStoreGoodsByQo(updateModel model.GoodsStore, whereModel model.GoodsStore)
	SelectListByQo(store model.GoodsStore) []model.GoodsStore
	TotalStoreNum() int64
	SelectPageByQo(dto dto.QueryStorageSituationDTO) vo.Page[model.GoodsStore]
	SelectTotalStoreNumById(id uint) int64
	SelectGoodsStorePageByQo(dto dto.QueryDetailStorageSituationDTO) vo.Page[model.GoodsStore]
}

type GoodsStoreDaoImpl struct {
}

func (deptDao GoodsStoreDaoImpl) StoreUsed(storeId uint) int64 {
	var count int64
	db := common.GetDB()
	db.Table("t_goods_store").Where("store_id = ?", storeId).Count(&count)
	return count
}

func (deptDao GoodsStoreDaoImpl) SelectGoodListWithResidueNum() []model.GoodsStore {
	db := common.GetDB()
	var models []model.GoodsStore
	db.Where("residue_num > 0").Find(&models)
	return models
}

func (deptDao GoodsStoreDaoImpl) SelectByQo(store model.GoodsStore) model.GoodsStore {
	db := common.GetDB()
	var models model.GoodsStore
	db.Where(&store).Find(&models)
	return models
}

func (deptDao GoodsStoreDaoImpl) Save(store model.GoodsStore) {
	db := common.GetDB()
	db.Create(&store)
}

func (deptDao GoodsStoreDaoImpl) UpdateStoreGoodsByQo(updateModel model.GoodsStore, whereModel model.GoodsStore) {
	db := common.GetDB()
	db.Where(&whereModel).Updates(&updateModel)
}

func (deptDao GoodsStoreDaoImpl) SelectListByQo(store model.GoodsStore) []model.GoodsStore {
	db := common.GetDB()
	var models []model.GoodsStore
	db.Where(&store).Find(&models)
	return models
}

func (deptDao GoodsStoreDaoImpl) TotalStoreNum() int64 {
	var sum int64
	db := common.GetDB()
	db.Table("t_goods_store").Select("SUM(residue_num) as sum").Find(&sum)
	return sum
}

func (deptDao GoodsStoreDaoImpl) SelectTotalStoreNumById(id uint) int64 {
	var sum int64
	db := common.GetDB()
	db.Table("t_goods_store").Select("SUM(residue_num) as sum").Where("store_id = ?", id).Find(&sum)
	return sum
}

func (deptDao GoodsStoreDaoImpl) SelectPageByQo(dto dto.QueryStorageSituationDTO) vo.Page[model.GoodsStore] {
	var page vo.Page[model.GoodsStore]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("t_goods_store").Select("store_id,store_name,SUM(residue_num) residue_num")

	name := dto.Name
	if name != "" && len(name) > 0 {
		db = db.Where("store_name like ?", name+"%")
	}
	db = db.Group("store_id,store_name")

	// 查询记录
	var records []model.GoodsStore
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (deptDao GoodsStoreDaoImpl) SelectGoodsStorePageByQo(dto dto.QueryDetailStorageSituationDTO) vo.Page[model.GoodsStore] {
	var page vo.Page[model.GoodsStore]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("t_goods_store").Select("*")

	storeId := dto.StoreId
	if storeId > 0 {
		db = db.Where("store_id = ?", storeId)
	}
	db = db.Where("residue_num > ?", 0)
	id := dto.Id
	if id > 0 {
		db = db.Where("goods_id = ?", id)
	}

	// 查询记录
	var records []model.GoodsStore
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}
