package service

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
)

type IGoodsCategoryService interface {
	QueryCategoryPageByQo(c *gin.Context, dto dto.QueryGoodsCategoryDTO)
	SaveCategory(c *gin.Context, supplier model.GoodsCategory)
	UpdateCategory(c *gin.Context, supplier model.GoodsCategory)
	NormalCategoryAll(c *gin.Context)
	DeactivateCategory(c *gin.Context, cid string)
}

type GoodsCategoryServiceImpl struct {
}

func (GoodsCategoryServiceImpl) QueryCategoryPageByQo(c *gin.Context, dto dto.QueryGoodsCategoryDTO) {
	page := goodsCategoryDao.SelectPageByQo(dto)
	response.Success(c, page, "操作成功")
}

func (GoodsCategoryServiceImpl) SaveCategory(c *gin.Context, goodsCategory model.GoodsCategory) {
	var queryModel = model.GoodsCategory{
		Name:  goodsCategory.Name,
		State: common.STATE_NORMAL,
	}
	one := goodsCategoryDao.SelectByQo(queryModel)
	if one != (model.GoodsCategory{}) {
		response.Error(c, "该分类已被创建")
		return
	}
	goodsCategory.State = common.STATE_NORMAL
	goodsCategoryDao.SaveCategory(goodsCategory)
	response.Success(c, nil, "操作成功")
}

func (GoodsCategoryServiceImpl) UpdateCategory(c *gin.Context, goodsCategory model.GoodsCategory) {
	category := goodsCategoryDao.SelectOneByQo(goodsCategory)
	if common.STATE_BAN == goodsCategory.State {
		//查看是否有上架商品正在使用
		var queryWrapper = model.Goods{
			CategoryID: goodsCategory.ID,
			State:      common.STATE_UP,
		}
		list := goodDao.SelectListByQo(queryWrapper)
		if len(list) > 0 {
			response.Error(c, "该分类正在被某个上架商品使用，请解除关系后，再操作")
			return
		}

		//移除分类
		if category != (model.GoodsCategory{}) {
			goodsCategoryDao.DeleteById(category.ID)
		}
	} else {
		if category != (model.GoodsCategory{}) {
			response.Error(c, "该分类已经存在")
			return
		}
	}
	goodsCategoryDao.UpdateCategory(goodsCategory)
	response.Success(c, nil, "操作成功")
}

func (GoodsCategoryServiceImpl) NormalCategoryAll(c *gin.Context) {
	var list = make([]map[string]interface{}, 0)
	for _, val := range goodsCategoryDao.SelectListByQo(model.GoodsCategory{State: common.STATE_NORMAL}) {
		var categoryMap = make(map[string]interface{})
		categoryMap["id"] = val.ID
		categoryMap["label"] = val.Name
		list = append(list, categoryMap)
	}
	response.Success(c, list, "操作成功")
}

func (GoodsCategoryServiceImpl) DeactivateCategory(c *gin.Context, cid string) {
	var queryWrapper = model.Goods{
		CategoryID: uint(utils.ConvertStringToInt64(cid)),
		State:      common.STATE_UP,
	}
	list := goodDao.SelectListByQo(queryWrapper)
	if len(list) > 0 {
		response.Error(c, "该分类正在被某个上架商品使用，请解除关系后，再操作")
		return
	}
	goodsCategory := goodsCategoryDao.SelectByQo(model.GoodsCategory{ID: uint(utils.ConvertStringToInt64(cid))})
	var goodsCategoryQueryModel = model.GoodsCategory{
		State: common.STATE_BAN,
		Name:  goodsCategory.Name,
		ID:    uint(utils.ConvertStringToInt64(cid)),
	}
	one := goodsCategoryDao.SelectOneByQo(goodsCategoryQueryModel)
	if one != (model.GoodsCategory{}) {
		var deleteQueryModel = model.GoodsCategory{
			State: common.STATE_BAN,
			Name:  goodsCategory.Name,
		}
		goodsCategoryDao.DeleteByQo(deleteQueryModel)
	}
	var updateQueryModel = model.GoodsCategory{
		State: common.STATE_BAN,
		ID:    uint(utils.ConvertStringToInt64(cid)),
	}
	goodsCategoryDao.UpdateCategory(updateQueryModel)
	response.Success(c, nil, "操作成功")
}
