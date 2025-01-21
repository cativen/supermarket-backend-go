package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type GoodsCategoryDao interface {
	SelectByQo(goods model.GoodsCategory) model.GoodsCategory
	SelectPageByQo(dto dto.QueryGoodsCategoryDTO) vo.Page[model.GoodsCategory]
	SaveCategory(category model.GoodsCategory)
	SelectOneByQo(goods model.GoodsCategory) model.GoodsCategory
	UpdateCategory(category model.GoodsCategory)
	DeleteById(id uint)
	SelectListByQo(goods model.GoodsCategory) []model.GoodsCategory
	DeleteByQo(category model.GoodsCategory)
}

type GoodsCategoryImpl struct {
}

func (GoodsCategoryImpl) SelectByQo(goods model.GoodsCategory) model.GoodsCategory {
	db := common.GetDB()
	var goodsCategory model.GoodsCategory
	db.Where(goods).First(&goodsCategory)
	return goodsCategory
}

func (GoodsCategoryImpl) SelectOneByQo(goods model.GoodsCategory) model.GoodsCategory {
	db := common.GetDB()
	var goodsCategory model.GoodsCategory
	db.Where("name = ? and state = ? and id != ?", goods.Name, goods.State, goods.ID).First(&goodsCategory)
	return goodsCategory
}

func (GoodsCategoryImpl) SelectPageByQo(dto dto.QueryGoodsCategoryDTO) vo.Page[model.GoodsCategory] {
	var page vo.Page[model.GoodsCategory]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("goods_category").Select("*")

	name := dto.Name
	if name != "" && len(name) > 0 {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	state := dto.State
	if state != "" && len(state) > 0 {
		db = db.Where("state = ?", state)
	}

	// 查询分类列表
	var records []model.GoodsCategory
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (GoodsCategoryImpl) SaveCategory(category model.GoodsCategory) {
	db := common.GetDB()
	db.Create(&category)
}

func (GoodsCategoryImpl) UpdateCategory(category model.GoodsCategory) {
	db := common.GetDB()
	db.Updates(&category)
}

func (GoodsCategoryImpl) DeleteById(id uint) {
	db := common.GetDB()
	deleteCategory := model.GoodsCategory{ID: id}
	db.Delete(&deleteCategory)
}

func (GoodsCategoryImpl) SelectListByQo(goods model.GoodsCategory) []model.GoodsCategory {
	db := common.GetDB()
	var categorys []model.GoodsCategory
	db.Where(&goods).Find(&categorys)
	return categorys
}

func (GoodsCategoryImpl) DeleteByQo(category model.GoodsCategory) {
	db := common.GetDB()
	db.Delete(&category)
}
