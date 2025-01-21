package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/service"
)

var storeService service.IStoreService = &service.StoreServiceImpl{}

func StoreList(c *gin.Context) {
	var dto dto.QueryStoreDTO
	c.ShouldBind(&dto)
	// 确保服务方法正确处理错误
	storeService.StoreList(c, dto)
}

func SaveStore(c *gin.Context) {
	var store model.Store
	c.ShouldBind(&store)
	storeService.SaveStore(c, store)
}

func UpdateStore(c *gin.Context) {
	var store model.Store
	c.ShouldBind(&store)
	storeService.UpdateStore(c, store)
}

func DeactivateStore(c *gin.Context) {
	sidStr, _ := c.GetPostForm("sid")
	storeService.DeactivateStore(c, sidStr)
}
