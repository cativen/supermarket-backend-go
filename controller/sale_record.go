package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/service"
	"superMarket-backend/utils"
)

func QueryPageByQoSaleRecords(c *gin.Context) {
	var saleDTO dto.SaleRecordPageDTO
	// 绑定请求体中的 JSON 数据到结构体
	if err := c.ShouldBindJSON(&saleDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 打印出 saleDTO 的值来调试
	log.Printf("Received saleDTO: %+v", saleDTO)
	var saleRecordService service.ISaleRecordsService = &service.SaleRecordsServiceImpl{}
	// 确保服务方法正确处理错误
	saleRecordService.QueryPageByQoSaleRecords(c, saleDTO)
}

func DelSaleRecords(c *gin.Context) {
	cn := c.Query("cn")
	var saleRecordService service.ISaleRecordsService = &service.SaleRecordsServiceImpl{}
	// 确保服务方法正确处理错误
	saleRecordService.DelSaleRecords(c, cn)
}

func GetSaleRecordsCn(c *gin.Context) {
	snowflake, err := utils.NewSnowflake(1)
	if err != nil {
		response.Error(c, "生成订单号失败")
	}
	generate, err := snowflake.Generate()
	if err != nil {
		response.Error(c, "生成订单号失败")
	}
	response.Success(c, generate, "操作成功")
}

func GetOptionSaleRecordsGoods(c *gin.Context) {
	var saleRecordService service.ISaleRecordsService = &service.SaleRecordsServiceImpl{}
	saleRecordService.GetOptionSaleRecordsGoods(c)
}

func SaveSaleRecords(c *gin.Context) {
	var saleRecordService service.ISaleRecordsService = &service.SaleRecordsServiceImpl{}
	cn := c.Query("cn")
	token := c.Query("token")
	saleRecordService.SaveSaleRecords(c, cn, token)
}

func PaySaleItems(c *gin.Context) {
	var saleRecordService service.ISaleRecordsService = &service.SaleRecordsServiceImpl{}
	var saleRecord model.SaleRecord
	c.ShouldBind(&saleRecord)
	token := c.GetHeader("token")
	saleRecordService.PaySaleItems(saleRecord, token)
}
