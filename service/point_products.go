package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
	"time"
)

var pointProductsDao dao.PointProductsDao = &dao.PointProductsImpl{}

type IPointProductsService interface {
	QueryOptionGoods(c *gin.Context)
	QueryPointPageByQo(c *gin.Context, dto dto.QueryPointProductsDTO)
	DelProductPoint(c *gin.Context, id string)
	SavePointGoods(c *gin.Context, product model.PointProduct, token string)
	QueryPointGoodsById(c *gin.Context, id string)
	UpdatePointGoods(c *gin.Context, product model.PointProduct, token string)
}

type PointProductsServiceImpl struct {
}

func (pointProductsService *PointProductsServiceImpl) QueryOptionGoods(c *gin.Context) {
	list := pointProductsDao.SelectListByQo(model.PointProduct{})
	//商品IDSet
	var goodsIdSet = make(map[uint]struct{})
	for _, val := range list {
		goodsIdSet[val.GoodsID] = struct{}{}
	}
	goodsIds := make([]uint, 0)
	for goodId, _ := range goodsIdSet {
		goodsIds = append(goodsIds, goodId)
	}
	goods := goodDao.SelectByGoodsIds(goodsIds)
	var options = make([]map[string]interface{}, 0)
	for _, val := range goods {
		var goodsMap = make(map[string]interface{})
		goodsMap["id"] = val.ID
		goodsMap["name"] = val.Name
		options = append(options, goodsMap)
	}
	response.Success(c, options, "操作成功")
}

func (pointProductsService *PointProductsServiceImpl) QueryPointPageByQo(c *gin.Context, dto dto.QueryPointProductsDTO) {
	page := pointProductsDao.QueryPageByQoPointProducts(dto)
	response.Success(c, page, "操作成功")
}

func (pointProductsService *PointProductsServiceImpl) DelProductPoint(c *gin.Context, id string) {
	pointProductsDao.DeleteByQo(model.PointProduct{GoodsID: uint(utils.ConvertStringToInt64(id))})
	response.Success(c, nil, "操作成功")
}

func (pointProductsService *PointProductsServiceImpl) SavePointGoods(c *gin.Context, product model.PointProduct, token string) {
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
	one := pointProductsDao.SelectByQo(model.PointProduct{GoodsID: product.GoodsID})
	if one != (model.PointProduct{}) {
		response.Error(c, "该商品已经是积分商品")
	}
	product.UpdateBy = existEmployee.NickName
	product.UpdateTime = time.Now()
	product.UpdateID = uint(existEmployee.ID)
	product.State = common.STATE_NORMAL
	pointProductsDao.SaveData(product)
	response.Success(c, nil, "操作成功")
}

func (pointProductsService *PointProductsServiceImpl) QueryPointGoodsById(c *gin.Context, id string) {
	one := pointProductsDao.SelectByQo(model.PointProduct{GoodsID: uint(utils.ConvertStringToInt64(id))})
	response.Success(c, one, "操作成功")
}

func (pointProductsService *PointProductsServiceImpl) UpdatePointGoods(c *gin.Context, product model.PointProduct, token string) {
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
	product.UpdateBy = existEmployee.NickName
	product.UpdateTime = time.Now()
	product.UpdateID = uint(existEmployee.ID)
	pointProductsDao.UpdateDataByQO(model.PointProduct{GoodsID: product.GoodsID}, product)
	response.Success(c, nil, "操作成功")
}
