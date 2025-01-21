package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/service"
)

var detailStoreGoodsService service.IDetailStoreGoodsService = &service.DetailStoreGoodsServiceImpl{}

func QueryStorePageByQoOut(c *gin.Context) {
	var dto dto.QueryDetailStoreGoodsOutDTO
	c.ShouldBind(&dto)
	detailStoreGoodsService.QueryStorePageByQoOut(c, dto)
}

func InitOutOptions(c *gin.Context) {
	detailStoreGoodsService.InitOutOptions(c)
}

func DelOut(c *gin.Context) {
	cn, _ := c.GetPostForm("cn")
	detailStoreGoodsService.DelOut(c, cn)
}
