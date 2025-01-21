package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type PointProductsDao interface {
	SelectListByQo(query model.PointProduct) []model.PointProduct
	QueryPageByQoPointProducts(dto dto.QueryPointProductsDTO) vo.Page[model.PointProduct]
	DeleteByQo(product model.PointProduct)
	SelectByQo(product model.PointProduct) model.PointProduct
	SaveData(product model.PointProduct)
	UpdateDataByQO(queryGoods model.PointProduct, updateGoods model.PointProduct)
	SelectListByGreater(val int64) []model.PointProduct
	SelectListByLesser(val int64) []model.PointProduct
}

type PointProductsImpl struct {
}

func (pointProductsDao PointProductsImpl) SelectListByQo(query model.PointProduct) []model.PointProduct {
	db := common.GetDB()
	var pointProducts = []model.PointProduct{}
	db.Where(query).Find(&pointProducts)
	return pointProducts
}

func (pointProductsDao PointProductsImpl) QueryPageByQoPointProducts(dto dto.QueryPointProductsDTO) vo.Page[model.PointProduct] {
	var page vo.Page[model.PointProduct]
	db := common.GetDB()
	//分页查询出所有的商品积分记录
	db = db.Table("point_products").Select("*")

	name := dto.Name
	if name != "" && len(name) > 0 {
		db = db.Where("goods_name like ?", "%"+name+"%")
	}

	// 查询商品积分列表
	var records []model.PointProduct
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (pointProductsDao PointProductsImpl) DeleteByQo(product model.PointProduct) {
	db := common.GetDB()
	db.Where(&product).Delete(&product)
}

func (pointProductsDao PointProductsImpl) SelectByQo(product model.PointProduct) model.PointProduct {
	db := common.GetDB()
	var res model.PointProduct
	db.Where(&product).First(&res)
	return res
}

func (pointProductsDao PointProductsImpl) SaveData(product model.PointProduct) {
	db := common.GetDB()
	db.Create(&product)
}

func (pointProductsDao PointProductsImpl) UpdateDataByQO(queryGoods model.PointProduct, updateGoods model.PointProduct) {
	db := common.GetDB()
	db.Where("goods_id = ?", queryGoods.GoodsID).Update("integral", updateGoods.Integral)
}

func (pointProductsDao PointProductsImpl) SelectListByGreater(val int64) []model.PointProduct {
	db := common.GetDB()
	var pointProducts = []model.PointProduct{}
	db.Where("integral >= ?", val).Find(&pointProducts)
	return pointProducts
}

func (pointProductsDao PointProductsImpl) SelectListByLesser(val int64) []model.PointProduct {
	db := common.GetDB()
	var pointProducts = []model.PointProduct{}
	db.Where("integral <= ?", val).Find(&pointProducts)
	return pointProducts
}
