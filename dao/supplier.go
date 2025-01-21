package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type SupplierDao interface {
	SelectByCn(id int64) model.Supplier
	SelectListByQo(supplier model.Supplier) []model.Supplier
	SelectPageByQo(dto dto.QuerySupplierDTO) vo.Page[model.Supplier]
	SelectOneByQo(supplier model.Supplier) model.Supplier
	SaveData(supplier model.Supplier)
	UpdateSupplier(supplier model.Supplier)
}

type SupplierDaoImpl struct {
}

func (SupplierDaoImpl) SelectByCn(id int64) model.Supplier {
	db := common.GetDB()
	var supplier model.Supplier
	db.Where("cn = ?", id).First(&supplier)
	return supplier
}

func (SupplierDaoImpl) SaveData(supplier model.Supplier) {
	db := common.GetDB()
	db.Create(supplier)
}

func (SupplierDaoImpl) UpdateSupplier(supplier model.Supplier) {
	db := common.GetDB()
	db.Updates(supplier)
}

func (SupplierDaoImpl) SelectListByQo(supplier model.Supplier) []model.Supplier {
	db := common.GetDB()
	var suppliers []model.Supplier
	db.Where(supplier).Find(&suppliers)
	return suppliers
}

func (SupplierDaoImpl) SelectOneByQo(supplier model.Supplier) model.Supplier {
	db := common.GetDB()
	var suppliers model.Supplier
	db.Where(supplier).First(&suppliers)
	return suppliers
}

func (SupplierDaoImpl) SelectPageByQo(dto dto.QuerySupplierDTO) vo.Page[model.Supplier] {
	var page vo.Page[model.Supplier]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("supplier").Select("*").Where("state = ?", common.STATE_NORMAL)

	name := dto.Name
	if name != "" && len(name) > 0 {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	address := dto.Address
	if address != "" && len(address) > 0 {
		db = db.Where("address LIKE ?", "%"+address+"%")
	}

	info := dto.Info
	if info != "" && len(info) > 0 {
		db = db.Where("info LIKE ?", "%"+info+"%")
	}

	// 查询供应商列表
	var records []model.Supplier
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}
