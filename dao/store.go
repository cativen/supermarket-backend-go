package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
)

type StoreDao interface {
	SelectListByQo(dto dto.QueryStoreDTO) []model.Store
	SelectStoreByQo(store model.Store) model.Store
	SaveStore(store model.Store)
	DeleteByQo(store model.Store)
	UpdateByQo(store model.Store)
	SelectById(id uint) model.Store
	ListByIds(ids []uint) []model.Store
	SelectResidueNumByGoodsId(id uint64) int64
	SelectListByQueryQo(query model.Store) []model.Store
}

type StoreDaoImpl struct {
}

func (StoreDaoImpl) SelectListByQo(dto dto.QueryStoreDTO) []model.Store {
	db := common.GetDB()
	db = db.Table("store").Select("*")
	name := dto.Name
	var stores []model.Store
	if name != "" && len(name) > 0 {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	state := dto.State
	if state != "" && len(state) > 0 {
		db = db.Where("state = ?", state)
	}
	db.Find(&stores)
	return stores
}

func (StoreDaoImpl) SelectStoreByQo(store model.Store) model.Store {
	db := common.GetDB()
	var existStore model.Store
	var queryStore = model.Store{
		Name:    store.Name,
		Address: store.Address,
		State:   store.State,
	}
	db.Where(&queryStore).First(&existStore)
	return existStore
}

func (StoreDaoImpl) SaveStore(store model.Store) {
	db := common.GetDB()
	db.Create(&store)
}

func (StoreDaoImpl) DeleteByQo(store model.Store) {
	db := common.GetDB()
	db.Delete(&store)
}

func (StoreDaoImpl) UpdateByQo(store model.Store) {
	db := common.GetDB()
	db.Updates(&store)
}

func (StoreDaoImpl) SelectById(id uint) model.Store {
	db := common.GetDB()
	var store model.Store
	db.Table("store").Where("id = ?", id).First(&store)
	return store
}

func (StoreDaoImpl) ListByIds(ids []uint) []model.Store {
	db := common.GetDB()
	var stores []model.Store
	db.Table("store").Where("id in (?)", ids).Find(&stores)
	return stores
}

func (StoreDaoImpl) SelectResidueNumByGoodsId(id uint64) int64 {
	db := common.GetDB()
	var count int64
	db.Table("t_goods_store").Select("sum(residue_num)").Where("goods_id = ?", id).Find(&count)
	return count
}

func (StoreDaoImpl) SelectListByQueryQo(query model.Store) []model.Store {
	db := common.GetDB()
	var stores []model.Store
	db.Where(&query).Find(&stores)
	return stores
}
