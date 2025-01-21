package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
	"time"
)

var detailStoreGoodsDao dao.DetailStoreGoodsDao = &dao.DetailStoreGoodsDaoImpl{}

type IDetailStoreGoodsService interface {
	QueryStorePageByQoOut(c *gin.Context, dto dto.QueryDetailStoreGoodsOutDTO)
	InitOutOptions(c *gin.Context)
	DelOut(c *gin.Context, cn string)
	SaveQoIn(c *gin.Context, detail model.DetailStoreGoods, token string)
	QueryPageByQoIn(c *gin.Context, inDTO dto.QueryDetailStoreGoodsOutDTO)
}

type DetailStoreGoodsServiceImpl struct {
}

func (detailStoreGoodsService DetailStoreGoodsServiceImpl) QueryStorePageByQoOut(c *gin.Context, dto dto.QueryDetailStoreGoodsOutDTO) {
	detailStoreGoodsService.QueryStorePageByQo(c, dto, common.TYPE_OUT)
}

func (detailStoreGoodsService DetailStoreGoodsServiceImpl) QueryPageByQoIn(c *gin.Context, dto dto.QueryDetailStoreGoodsOutDTO) {
	detailStoreGoodsService.QueryStorePageByQo(c, dto, common.TYPE_IN)
}

func (DetailStoreGoodsServiceImpl) QueryStorePageByQo(c *gin.Context, dto dto.QueryDetailStoreGoodsOutDTO, storeTypes string) {
	var storeDao dao.StoreDao = &dao.StoreDaoImpl{}
	detailStoreGoodPage := detailStoreGoodsDao.SelectPageByQoDetailStoreGoods(dto, storeTypes)
	var page vo.Page[vo.DetailStoreGoodsOutVo]
	var records = make([]vo.DetailStoreGoodsOutVo, 0)
	goods := detailStoreGoodPage.Records
	for _, good := range goods {
		var detail vo.DetailStoreGoodsOutVo
		copier.Copy(&detail, &good)
		//通过仓库ID查询仓库信息
		store := storeDao.SelectById(good.StoreID)
		if store != (model.Store{}) {
			detail.StoreName = store.Name
		}
		detail.CreateTime = good.CreateTime.Format("2006-01-02 15:04:05")
		records = append(records, detail)
	}
	page.Records = records
	page.Current = dto.CurrentPage
	page.Total = detailStoreGoodPage.Total
	page.Size = dto.PageSize
	response.Success(c, page, "操作成功")
}
func (DetailStoreGoodsServiceImpl) InitOutOptions(c *gin.Context) {
	//查询库存大于0的商品
	var goodsStoreDao dao.GoodsStoreDao = &dao.GoodsStoreDaoImpl{}
	list := goodsStoreDao.SelectGoodListWithResidueNum()
	if len(list) == 0 {
		response.Error(c, "库存中没有存放商品")
	}
	//商品IDSet
	var goodsIdSet = make(map[uint]struct{})
	var storeIdSet = make(map[uint]struct{})
	for _, value := range list {
		goodsIdSet[value.GoodsID] = struct{}{}
		storeIdSet[value.StoreID] = struct{}{}
	}
	goodsIds := make([]uint, len(goodsIdSet))
	storeIds := make([]uint, len(storeIdSet))
	for goodId, _ := range goodsIdSet {
		goodsIds = append(goodsIds, goodId)
	}

	for storeId, _ := range storeIdSet {
		storeIds = append(storeIds, storeId)
	}

	var resMap = make(map[string]interface{})

	goods := goodsDao.ListByIds(goodsIds)
	var goodsList = make([]map[string]interface{}, 0)
	for _, good := range goods {
		var goodsMap = make(map[string]interface{})
		goodsMap["id"] = good.ID
		goodsMap["name"] = good.Name
		goodsList = append(goodsList, goodsMap)
	}

	stores := storeDao.ListByIds(storeIds)
	var storeList = make([]map[string]interface{}, 0)
	for _, store := range stores {
		var storeMap = make(map[string]interface{})
		storeMap["id"] = store.ID
		storeMap["name"] = store.Name
		storeList = append(storeList, storeMap)
	}
	resMap["goods"] = goodsList
	resMap["stores"] = storeList

	response.Success(c, resMap, "操作成功")

}

func (DetailStoreGoodsServiceImpl) DelOut(c *gin.Context, cn string) {
	var model = model.DetailStoreGoods{
		Cn:     cn,
		State1: common.STATE_DEL,
	}
	detailStoreGoodsDao.UpdateStoreGoods(model)
	response.Success(c, nil, "操作成功")
}

func (DetailStoreGoodsServiceImpl) SaveQoIn(c *gin.Context, detail model.DetailStoreGoods, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println(err)
	}
	var existEmployee model.Employee
	if err := json.Unmarshal([]byte(val), &existEmployee); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}
	detail.Type = common.TYPE_IN
	detail.State = common.STATE_NORMAL
	detail.CreateID = existEmployee.ID
	detail.CreateBy = existEmployee.NickName
	//雪花ID
	sf, err := utils.NewSnowflake(1)
	if err != nil {
		response.Error(c, "生成ID失败")
		return
	}
	generateId, _ := sf.Generate()
	detail.Cn = utils.ConvertInt64ToString(generateId)
	detail.CreateTime = time.Now()
	detail.State1 = common.STATE_NORMAL

	//更新库存
	var goodsStoreService IGoodsStoreService = &GoodsStoreServiceImpl{}
	goodsStoreService.GoodsInStore(detail.GoodsID, detail.GoodsNum, detail.StoreID)
	//查询供应商信息
	var supplierDao dao.SupplierDao = &dao.SupplierDaoImpl{}
	supplier := supplierDao.SelectByCn(detail.SupplierID)
	detail.SupplierName = &supplier.Name
	detailStoreGoodsDao.SaveData(detail)
}
