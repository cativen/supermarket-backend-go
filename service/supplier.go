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

var supplierDao dao.SupplierDao = &dao.SupplierDaoImpl{}

type ISupplierService interface {
	QuerySupplierPageByQo(c *gin.Context, dto dto.QuerySupplierDTO)
	SaveSupplier(c *gin.Context, supplier model.Supplier)
	UpdateSupplier(c *gin.Context, supplier model.Supplier)
	QueryByCnSupplier(c *gin.Context, cn string)
	DeactivateSupplier(c *gin.Context, cn string)
}

type SupplierServiceImpl struct {
}

func (SupplierServiceImpl) QuerySupplierPageByQo(c *gin.Context, dto dto.QuerySupplierDTO) {
	page := supplierDao.SelectPageByQo(dto)
	response.Success(c, page, "操作成功")
}

func (SupplierServiceImpl) SaveSupplier(c *gin.Context, supplier model.Supplier) {
	supplier.State = common.STATE_NORMAL
	var querySupplier = model.Supplier{
		Name:  supplier.Name,
		State: common.STATE_NORMAL,
	}
	one := supplierDao.SelectOneByQo(querySupplier)
	if one != (model.Supplier{}) {
		response.Error(c, "已存在供货商的联系方式")
		return
	}
	snowflake, _ := utils.NewSnowflake(1)
	generateId, _ := snowflake.Generate()
	supplier.Cn = generateId
	supplierDao.SaveData(supplier)
	response.Success(c, nil, "操作成功")
}

func (SupplierServiceImpl) UpdateSupplier(c *gin.Context, supplier model.Supplier) {
	if supplier.State == common.STATE_NORMAL {
		var querySupplier = model.Supplier{
			Name:  supplier.Name,
			State: common.STATE_NORMAL,
			Cn:    supplier.Cn,
		}
		one := supplierDao.SelectOneByQo(querySupplier)
		if one != (model.Supplier{}) && one.Cn != supplier.Cn {
			response.Error(c, "该供货商已存在")
			return
		}
	}
	supplierDao.UpdateSupplier(supplier)
	response.Success(c, nil, "操作成功")
}

func (SupplierServiceImpl) QueryByCnSupplier(c *gin.Context, cn string) {
	byCn := supplierDao.SelectByCn(utils.ConvertStringToInt64(cn))
	response.Success(c, byCn, "操作成功")
}

func (SupplierServiceImpl) DeactivateSupplier(c *gin.Context, cn string) {
	var queryModel = model.DetailStoreGoods{
		State1:     common.STATE_NORMAL,
		Type:       common.TYPE_IN,
		State:      common.STATE_NORMAL,
		SupplierID: utils.ConvertStringToInt64(cn),
	}
	list := detailStoreGoodsDao.SelectListByQo(queryModel)
	if len(list) > 0 {
		response.Error(c, "该供货商正在被入库订单使用，请解除关系之后在停用")
		return
	}
	var updateModel = model.Supplier{
		State: common.STATE_BAN,
		Cn:    utils.ConvertStringToInt64(cn),
	}
	supplierDao.UpdateSupplier(updateModel)
	response.Success(c, nil, "操作成功")
}
