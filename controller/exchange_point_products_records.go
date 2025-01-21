package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/service"
)

func QueryPageByQoExchangePointProducts(c *gin.Context) {
	var exchangeDTO dto.QueryExchangePointProductsRecordsDTO
	// 绑定请求体中的 JSON 数据到结构体
	if err := c.ShouldBindJSON(&exchangeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 打印出 saleDTO 的值来调试
	log.Printf("Received saleDTO: %+v", exchangeDTO)
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	// 确保服务方法正确处理错误
	ExchangePointProductsService.QueryPageByQoExchangePointProducts(c, exchangeDTO)
}

func QueryOptionsMemberPhone(c *gin.Context) {
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	// 确保服务方法正确处理错误
	ExchangePointProductsService.QueryOptionsMemberPhone(c)
}

func DelExchangePointProducts(c *gin.Context) {
	cn := c.Query("cn")
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	// 确保服务方法正确处理错误
	ExchangePointProductsService.DelExchangePointProducts(c, cn)
}

func QueryOptionsMember(c *gin.Context) {
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	ExchangePointProductsService.QueryOptionsMember(c)
}

func QueryOptionsPointProducts(c *gin.Context) {
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	ExchangePointProductsService.QueryOptionsPointProducts(c)
}

func QueryPointProductByMemberId(c *gin.Context) {
	memberId := c.Query("memberId")
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	ExchangePointProductsService.QueryPointProductByMemberId(c, memberId)
}

func QueryMemberByGoodsId(c *gin.Context) {
	goodsId := c.Query("goodsId")
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	ExchangePointProductsService.QueryMemberByGoodsId(c, goodsId)
}

func QueryPointProductByGoodsId(c *gin.Context) {
	goodsId := c.Query("goodsId")
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	ExchangePointProductsService.QueryPointProductByGoodsId(c, goodsId)
}

func SaveExchangePointProductRecords(c *gin.Context) {
	var exchangePointProducts model.ExchangePointProductsRecord
	token := c.GetHeader("token")
	var ExchangePointProductsService service.IExchangePointProductsService = &service.ExchangePointProductsServiceImpl{}
	ExchangePointProductsService.SaveExchangePointProductRecords(c, exchangePointProducts, token)
}
