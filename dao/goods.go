package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type GoodsDao interface {
	SelectById(id int64) model.Goods
	ListByIds(ids []uint) []model.Goods
	SelectListByQo(goods model.Goods) []model.Goods
	SelectPageListByQo(dto dto.QueryGoodsDTO) vo.Page[vo.GoodsListVo]
	SaveGoods(good model.Goods)
	UpdateDataByQO(queryGoods model.Goods, UpdateGoods model.Goods)
	SelectByGoodsIds(ids []uint) []model.Goods
	QueryPageStatisticSaleByQo(name string) int64
	SelectHaveResidueGood() []model.Goods
}

type GoodsDaoImpl struct {
}

func (GoodsDao *GoodsDaoImpl) SelectById(id int64) model.Goods {
	var goods model.Goods
	db := common.GetDB()
	db.First(&goods, id)
	return goods
}

func (GoodsDao *GoodsDaoImpl) ListByIds(ids []uint) []model.Goods {
	var goods []model.Goods
	db := common.GetDB()
	db.Where("id in (?)", ids).Find(&goods)
	return goods
}

func (GoodsDao *GoodsDaoImpl) SelectListByQo(goods model.Goods) []model.Goods {
	db := common.GetDB()
	var goodList []model.Goods
	db.Where(&goods).Find(&goodList)
	return goodList
}

func (GoodsDao *GoodsDaoImpl) SelectPageListByQo(dto dto.QueryGoodsDTO) vo.Page[vo.GoodsListVo] {
	var page vo.Page[vo.GoodsListVo]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("goods").Select("*")

	id := dto.ID
	if id != 0 {
		db = db.Where("id = ?", id)
	}

	sellPrice := dto.SellPrice
	if sellPrice != 0 {
		db = db.Where("sell_price = ?", sellPrice)
	}

	name := dto.Name
	if name != "" && len(name) > 0 {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	categoryId := dto.CategoryID
	if categoryId != 0 {
		db = db.Where("category_id = ?", categoryId)
	}

	state := dto.State
	if state != "" && len(state) > 0 {
		db = db.Where("state = ?", state)
	}

	startTime := dto.OperateStartTime
	if startTime != "" && len(startTime) > 0 {
		db = db.Where("update_time >= ?", startTime)
	}

	endTime := dto.OperateEndTime
	if endTime != "" && len(endTime) > 0 {
		db = db.Where("update_time <= ?", endTime)
	}

	var records []vo.GoodsListVo
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (GoodsDao *GoodsDaoImpl) SaveGoods(good model.Goods) {
	db := common.GetDB()
	db.Create(&good)
}

func (GoodsDao *GoodsDaoImpl) UpdateDataByQO(queryGoods model.Goods, UpdateGoods model.Goods) {
	db := common.GetDB()
	db.Where(&queryGoods).Updates(&UpdateGoods)
}

func (GoodsDao *GoodsDaoImpl) SelectByGoodsIds(ids []uint) []model.Goods {
	db := common.GetDB()
	var goodList []model.Goods
	db.Where("state =  ? and id not in (?)", common.STATE_UP, ids).Find(&goodList)
	return goodList
}

func (GoodsDao *GoodsDaoImpl) QueryPageStatisticSaleByQo(name string) int64 {
	db := common.GetDB()
	var sum int64
	db.Table("goods").Select("sum(sales_volume) as sum").Where("state='0'").Find(&sum)
	return sum
}

func (GoodsDao *GoodsDaoImpl) SelectHaveResidueGood() []model.Goods {
	db := common.GetDB()
	var goodList []model.Goods
	db.Where("residue_num>=0").Find(&goodList)
	return goodList
}
