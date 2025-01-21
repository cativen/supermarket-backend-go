package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
)

func QueryPageStatisticSaleByQo(c *gin.Context) {
	var dto dto.QueryStatisticSaleDTO
	c.ShouldBind(&dto)
	goodService.QueryPageStatisticSaleByQo(c, dto)
}
