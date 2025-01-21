package service

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
)

var storeDao dao.StoreDao = &dao.StoreDaoImpl{}

type IStoreService interface {
	StoreList(c *gin.Context, dto dto.QueryStoreDTO)
	SaveStore(c *gin.Context, store model.Store)
	UpdateStore(c *gin.Context, store model.Store)
	DeactivateStore(c *gin.Context, str string)
}

type StoreServiceImpl struct {
}

func (StoreServiceImpl) StoreList(c *gin.Context, dto dto.QueryStoreDTO) {
	list := storeDao.SelectListByQo(dto)
	response.Success(c, list, "操作成功")
}

func (StoreServiceImpl) SaveStore(c *gin.Context, store model.Store) {
	existStore := storeDao.SelectStoreByQo(store)
	if existStore.ID != 0 || existStore != (model.Store{}) {
		response.Error(c, "创建失败，已有相同的仓库")
		return
	}
	storeDao.SaveStore(store)
	response.Success(c, nil, "操作成功")
}

func (StoreServiceImpl) UpdateStore(c *gin.Context, store model.Store) {
	existStore := storeDao.SelectStoreByQo(store)
	if common.STATE_BAN == existStore.State {
		var goodsStoreDao dao.GoodsStoreDao = &dao.GoodsStoreDaoImpl{}
		count := goodsStoreDao.StoreUsed(store.ID)
		//要修改为停用状态
		if count != 0 {
			response.Error(c, "仓库中存在商品，不能停用仓库")
			return
		}
		if existStore != (model.Store{}) {
			storeDao.DeleteByQo(store)
		}
	}
	storeDao.UpdateByQo(store)
	response.Success(c, nil, "操作成功")
}

func (StoreServiceImpl) DeactivateStore(c *gin.Context, str string) {
	toInt64Id := utils.ConvertStringToInt64(str)
	var goodsStoreDao dao.GoodsStoreDao = &dao.GoodsStoreDaoImpl{}
	count := goodsStoreDao.StoreUsed(uint(toInt64Id))
	if count != 0 {
		response.Error(c, "仓库中存在商品，不能停用仓库")
		return
	} else {
		var store = model.Store{
			ID:    uint(toInt64Id),
			State: common.STATE_BAN,
		}
		storeDao.UpdateByQo(store)
	}
	response.Success(c, nil, "操作成功")
}
