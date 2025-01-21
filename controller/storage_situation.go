package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/service"
)

func QueryPageStorageSituationByQo(c *gin.Context) {
	var dto dto.QueryStorageSituationDTO
	c.ShouldBind(&dto)
	var goodsStoreService service.IGoodsStoreService = &service.GoodsStoreServiceImpl{}
	goodsStoreService.QueryPageStorageSituationByQo(c, dto)
}

func QueryStoreGoodsByStoreId(c *gin.Context) {
	var dto dto.QueryDetailStorageSituationDTO
	c.ShouldBind(&dto)
	var goodsStoreService service.IGoodsStoreService = &service.GoodsStoreServiceImpl{}
	goodsStoreService.QueryStoreGoodsByStoreId(c, dto)
}
