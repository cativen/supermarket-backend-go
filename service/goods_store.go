package service

import (
	"github.com/gin-gonic/gin"
	"math"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/vo"
)

var goodsStoreDao dao.GoodsStoreDao = &dao.GoodsStoreDaoImpl{}

type IGoodsStoreService interface {
	GoodsInStore(goodId int64, goodNum int64, storeId uint)
	QueryPageStorageSituationByQo(c *gin.Context, dto dto.QueryStorageSituationDTO)
	QueryStoreGoodsByStoreId(c *gin.Context, situationDTO dto.QueryDetailStorageSituationDTO)
}

type GoodsStoreServiceImpl struct {
}

func (GoodsStoreServiceImpl) GoodsInStore(goodId int64, goodNum int64, storeId uint) {
	var goodsStore model.GoodsStore
	records := goodsStoreDao.SelectByQo(goodsStore)
	if records != (model.GoodsStore{}) {
		residueNum := records.ResidueNum + goodNum
		inNum := records.InNum + goodNum
		var updateModel = model.GoodsStore{
			ResidueNum: residueNum,
			InNum:      inNum,
		}
		var whereModel = model.GoodsStore{
			GoodsID: uint(goodId),
			StoreID: storeId,
		}
		goodsStoreDao.UpdateStoreGoodsByQo(updateModel, whereModel)
	} else {
		var goodsStores = new(model.GoodsStore)
		goodsStores.GoodsID = uint(goodId)
		goodsStores.InNum = goodNum
		goodsStores.StoreID = storeId
		goodsStores.ResidueNum = goodNum
		one := storeDao.SelectById(storeId)
		goodsStore.StoreName = one.Name
		goodsStoreDao.Save(goodsStore)
	}
}

func (GoodsStoreServiceImpl) QueryPageStorageSituationByQo(c *gin.Context, dto dto.QueryStorageSituationDTO) {
	var storeMap = make(map[string]interface{})
	totalNum := goodsStoreDao.TotalStoreNum()
	storeMap["totalStoreNum"] = totalNum
	page := goodsStoreDao.SelectPageByQo(dto)
	records := page.Records
	var storageSituationPage vo.Page[vo.StorageSituationVo]
	var storageSituationList = make([]vo.StorageSituationVo, 0)
	for _, val := range records {
		var vo = vo.StorageSituationVo{}
		vo.StoreId = val.StoreID
		vo.StoreName = val.StoreName
		vo.ResidueNum = val.ResidueNum
		storageSituationList = append(storageSituationList, vo)
	}
	storageSituationPage.Records = storageSituationList
	storageSituationPage.Total = page.Total
	storeMap["page"] = storageSituationPage
	response.Success(c, storeMap, "操作成功")
}

func (GoodsStoreServiceImpl) QueryStoreGoodsByStoreId(c *gin.Context, dto dto.QueryDetailStorageSituationDTO) {
	var storeMap = make(map[string]interface{})
	totalStoreNum := goodsStoreDao.SelectTotalStoreNumById(dto.StoreId)
	storeMap["totalStoreNum1"] = totalStoreNum //该仓库的存储量

	var queryGoodsStore = model.GoodsStore{
		ResidueNum: 0,
		StoreID:    dto.StoreId,
	}
	list := goodsStoreDao.SelectListByQo(queryGoodsStore)
	var goodsIdSet = make(map[uint]struct{})
	for _, val := range list {
		goodsIdSet[val.GoodsID] = struct{}{}
	}
	if len(goodsIdSet) <= 0 {
		response.Error(c, "该仓库没有存放任何的商品")
		return
	}
	var goodsIdList []uint
	for uid, _ := range goodsIdSet {
		goodsIdList = append(goodsIdList, uid)
	}
	goodsList := goodDao.ListByIds(goodsIdList)
	var optionsStoreGoods = make([]map[string]interface{}, 0)
	for _, val := range goodsList {
		var goodsMap = make(map[string]interface{})
		goodsMap["id"] = val.ID
		goodsMap["name"] = val.Name
		optionsStoreGoods = append(optionsStoreGoods, goodsMap)
	}
	storeMap["optionsStoreGoods"] = optionsStoreGoods

	goodsStorePage := goodsStoreDao.SelectGoodsStorePageByQo(dto)

	var voPage = new(vo.Page[vo.DetailStorageSituationVo])
	voPage.Current = goodsStorePage.Current
	voPage.Size = goodsStorePage.Size
	voPage.Total = goodsStorePage.Total
	records := goodsStorePage.Records
	var detailStorageSituationVoList = make([]vo.DetailStorageSituationVo, 0)
	for _, val := range records {
		var vo = vo.DetailStorageSituationVo{}
		vo.GoodsId = val.GoodsID
		vo.ResidueNum = val.ResidueNum
		// 将 ResidueNum 和 totalStoreNum 转换为 float64 类型
		percentage := float64(vo.ResidueNum) / float64(totalStoreNum)
		// 保留两位小数
		vo.Percentage = math.Round(percentage * 100)
		good := goodDao.SelectById(int64(val.GoodsID))
		vo.GoodsName = good.Name
		detailStorageSituationVoList = append(detailStorageSituationVoList, vo)
	}
	voPage.Records = detailStorageSituationVoList
	storeMap["vos"] = voPage
	response.Success(c, storeMap, "操作成功")
}
