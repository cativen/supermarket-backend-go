package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/service"
)

var menuService service.IMenuService = &service.MenuServiceImpl{}

func QueryMenuPageByQo(c *gin.Context) {
	var dto dto.MenuQueryDTO
	c.ShouldBind(&dto)
	menuService.QueryMenuPageByQo(c, dto)
}
