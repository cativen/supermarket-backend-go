package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/service"
)

var supplierService service.ISupplierService = &service.SupplierServiceImpl{}

func QuerySupplierPageByQo(c *gin.Context) {
	var dto dto.QuerySupplierDTO
	c.ShouldBind(&dto)
	supplierService.QuerySupplierPageByQo(c, dto)
}

func SaveSupplier(c *gin.Context) {
	var supplier model.Supplier
	c.ShouldBind(&supplier)
	supplierService.SaveSupplier(c, supplier)
}

func UpdateSupplier(c *gin.Context) {
	var supplier model.Supplier
	c.ShouldBind(&supplier)
	supplierService.UpdateSupplier(c, supplier)
}

func QueryByCnSupplier(c *gin.Context) {
	cn := c.Query("cn")
	supplierService.QueryByCnSupplier(c, cn)
}

func DeactivateSupplier(c *gin.Context) {
	cn, _ := c.GetPostForm("cn")
	supplierService.DeactivateSupplier(c, cn)
}
