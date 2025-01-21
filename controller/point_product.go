package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/service"
)

var pointProductsService service.IPointProductsService = &service.PointProductsServiceImpl{}

func QueryOptionGoods(c *gin.Context) {
	pointProductsService.QueryOptionGoods(c)
}

func QueryPointPageByQo(c *gin.Context) {
	var dto dto.QueryPointProductsDTO
	c.ShouldBind(&dto)
	pointProductsService.QueryPointPageByQo(c, dto)
}

func DelProductPoint(c *gin.Context) {
	id := c.Query("id")
	pointProductsService.DelProductPoint(c, id)
}

func SavePointGoods(c *gin.Context) {
	var pointProduct model.PointProduct
	c.ShouldBind(&pointProduct)
	token := c.GetHeader("token")
	pointProductsService.SavePointGoods(c, pointProduct, token)
}

func QueryPointGoodsById(c *gin.Context) {
	goodsId := c.Query("goodsId")
	pointProductsService.QueryPointGoodsById(c, goodsId)
}

func UpdatePointGoods(c *gin.Context) {
	var pointProduct model.PointProduct
	c.ShouldBind(&pointProduct)
	token := c.GetHeader("token")
	pointProductsService.UpdatePointGoods(c, pointProduct, token)
}
