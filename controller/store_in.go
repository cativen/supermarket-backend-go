package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/service"
)

func SaveQoIn(c *gin.Context) {
	token := c.GetHeader("token")
	var detail model.DetailStoreGoods
	detailStoreGoodsService.SaveQoIn(c, detail, token)
}

func QueryPageByQoIn(c *gin.Context) {
	var dto dto.QueryDetailStoreGoodsOutDTO
	c.ShouldBind(&dto)
	detailStoreGoodsService.QueryPageByQoIn(c, dto)
}

func DelIn(c *gin.Context) {
	cn := c.Query("cn")
	var goodServices service.IGoodsService = &service.GoodsServiceImpl{}
	goodServices.DelIn(c, cn)
}

func QueryOptionsSuppliers(c *gin.Context) {
	var goodServices service.IGoodsService = &service.GoodsServiceImpl{}
	goodServices.QueryOptionsSuppliers(c)
}
