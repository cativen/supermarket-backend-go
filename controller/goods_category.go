package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/service"
)

var goodsCategoryService service.IGoodsCategoryService = &service.GoodsCategoryServiceImpl{}

func QueryCategoryPageByQo(c *gin.Context) {
	var dto dto.QueryGoodsCategoryDTO
	c.ShouldBind(&dto)
	goodsCategoryService.QueryCategoryPageByQo(c, dto)
}

func SaveCategory(c *gin.Context) {
	var supplier model.GoodsCategory
	c.ShouldBind(&supplier)
	goodsCategoryService.SaveCategory(c, supplier)
}

func UpdateCategory(c *gin.Context) {
	var supplier model.GoodsCategory
	c.ShouldBind(&supplier)
	goodsCategoryService.UpdateCategory(c, supplier)
}

func NormalCategoryAll(c *gin.Context) {
	goodsCategoryService.NormalCategoryAll(c)
}

func DeactivateCategory(c *gin.Context) {
	cid, _ := c.GetPostForm("cid")
	goodsCategoryService.DeactivateCategory(c, cid)
}
