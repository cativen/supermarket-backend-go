package route

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/controller"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有源
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	r.POST("/login", controller.Login)                 //登录
	r.GET("/checkedToken", controller.CheckedToken)    //鉴权
	r.GET("/empMenu", controller.EmpMenu)              //查询菜单
	r.GET("/exit", controller.Exit)                    //退出功能
	r.POST("/logout", controller.Logout)               //注销功能
	r.GET("/static/img/:filename", controller.ViewImg) //文件读取

	//个人信息用户组
	personalGroup := r.Group("/personal")
	personalGroup.GET("/information", controller.Information)
	personalGroup.POST("/edit_pwd", controller.EditUserPwd)

	//部门接口
	deptGroup := r.Group("/personnel_management/dept")
	deptGroup.GET("/list", controller.List)
	deptGroup.POST("/save", controller.SaveDept)
	deptGroup.POST("/update", controller.UpdateDept)
	deptGroup.POST("/deactivate", controller.DeactivateDept)

	//销售记录接口
	saleRecordGroup := r.Group("/sale_management/sale_record")
	saleRecordGroup.POST("/queryPageByQoSaleRecords", controller.QueryPageByQoSaleRecords)
	saleRecordGroup.GET("/delSaleRecords", controller.DelSaleRecords)
	saleRecordGroup.GET("/getCn", controller.GetSaleRecordsCn)
	saleRecordGroup.GET("/getOptionSaleRecordsGoods", controller.GetOptionSaleRecordsGoods)
	saleRecordGroup.GET("/saveSaleRecords", controller.SaveSaleRecords)
	saleRecordGroup.POST("/pay", controller.PaySaleItems)

	//积分兑换接口
	exchangePointProductsGroup := r.Group("/sale_management/exchange_point_products_records")
	exchangePointProductsGroup.POST("/queryPageByQoExchangePointProducts", controller.QueryPageByQoExchangePointProducts)
	exchangePointProductsGroup.GET("/queryOptionsMemberPhone", controller.QueryOptionsMemberPhone)
	exchangePointProductsGroup.GET("/delExchangePointProducts", controller.DelExchangePointProducts)
	exchangePointProductsGroup.GET("/queryOptionsMember", controller.QueryOptionsMember)
	exchangePointProductsGroup.GET("/queryOptionsPointProducts", controller.QueryOptionsPointProducts)
	exchangePointProductsGroup.GET("/queryPointProductBymemberId", controller.QueryPointProductByMemberId)
	exchangePointProductsGroup.GET("/queryMemberByGoodsId", controller.QueryMemberByGoodsId)
	exchangePointProductsGroup.GET("/queryPointProductByGoodsId", controller.QueryPointProductByGoodsId)
	exchangePointProductsGroup.POST("/saveExchangePointProductRecords", controller.SaveExchangePointProductRecords)

	//角色管理接口
	roleGroup := r.Group("/system/role")
	roleGroup.POST("/list", controller.RoleList)
	roleGroup.POST("/forbiddenRole", controller.ForbiddenRole)
	roleGroup.POST("/edit_role", controller.EditRole)
	roleGroup.POST("/save", controller.SaveRole)
	roleGroup.POST("/checkPermissons", controller.CheckRolePermissions)
	roleGroup.POST("/saveRolePermissons", controller.SaveRolePermissions)
	roleGroup.GET("/all", controller.AllRole)
	roleGroup.GET("/queryRoleIdsByEid", controller.QueryRoleIdsByEid)
	roleGroup.POST("/saveRoleEmp", controller.SaveRoleEmp)
	roleGroup.POST("/exportExcel", controller.ExportRoleExcel)

	//员工管理接口
	employeeGroup := r.Group("/personnel_management/employee")
	employeeGroup.POST("/list", controller.EmployeeList)
	employeeGroup.GET("/detail", controller.EmployeeDetail)
	employeeGroup.POST("/uploadImg", controller.UploadImg)
	employeeGroup.POST("/update", controller.UpdateEmployee)
	employeeGroup.GET("/editbtn", controller.EditEmployeeBtn)
	employeeGroup.POST("/deactivate", controller.DeactivateEmp)
	employeeGroup.POST("/resetPwd", controller.ResetEmpPwd)

	//仓库管理接口
	storeGroup := r.Group("/inventory_management/store")
	storeGroup.POST("/list", controller.StoreList)
	storeGroup.POST("/save", controller.SaveStore)
	storeGroup.POST("/update", controller.UpdateStore)
	storeGroup.POST("/deactivate", controller.DeactivateStore)

	//出库管理接口
	outGoodGroup := r.Group("/inventory_management/detail_store_goods_out")
	outGoodGroup.POST("/queryPageByQoOut", controller.QueryStorePageByQoOut)
	outGoodGroup.GET("/initOutOptions", controller.InitOutOptions)
	outGoodGroup.POST("/delOut", controller.DelOut)

	//入库管理接口
	inGoodGroup := r.Group("/inventory_management/detail_store_goods_in")
	inGoodGroup.POST("/save", controller.SaveQoIn)
	inGoodGroup.POST("/queryPageByQo", controller.QueryPageByQoIn)
	inGoodGroup.POST("/delIn", controller.DelIn)
	inGoodGroup.GET("/queryOptionsSuppliers", controller.QueryOptionsSuppliers)

	//商品信息
	goodGroup := r.Group("/goods_management/goods")
	goodGroup.GET("/selected_goodsAll", controller.SelectedGoodsAll)
	goodGroup.POST("/queryPageByQo", controller.QueryGoodPageByQo)
	goodGroup.POST("/save", controller.SaveGoods)
	goodGroup.POST("/upOrdown", controller.UpOrDownGoods)
	goodGroup.GET("/queryGoodsById", controller.QueryGoodsById)
	goodGroup.GET("/selected_storeAll", controller.SelectedStoreAll)
	goodGroup.POST("/uploadImg", controller.UploadGoodImg)
	goodGroup.POST("/update", controller.UpdateGoods)
	goodGroup.POST("/gopayById", controller.GoPayById)
	goodGroup.GET("/return", controller.Return)

	//供应商信息
	supplierGroup := r.Group("/inventory_management/supplier")
	supplierGroup.POST("/queryPageByQo", controller.QuerySupplierPageByQo)
	supplierGroup.POST("/save", controller.SaveSupplier)
	supplierGroup.POST("/update", controller.UpdateSupplier)
	supplierGroup.GET("/queryByCn", controller.QueryByCnSupplier)
	supplierGroup.POST("/deactivate", controller.DeactivateSupplier)

	//库存统计
	storageGroup := r.Group("/inventory_management/store/storage_situation")
	storageGroup.POST("/queryPageByQo", controller.QueryPageStorageSituationByQo)
	storageGroup.POST("/queryStoreGoodsByStoreId", controller.QueryStoreGoodsByStoreId)

	//会员管理
	memberGroup := r.Group("/member_management/member")
	memberGroup.POST("/queryPageByQo", controller.QueryMemberPageByQo)
	memberGroup.POST("/delMember", controller.DelMember)
	memberGroup.POST("/save", controller.SaveMember)
	memberGroup.GET("/queryMemberById", controller.QueryMemberById)
	memberGroup.POST("/update", controller.UpdateMember)
	memberGroup.GET("/queryMemberByPhone", controller.QueryMemberByPhone)

	//菜单管理
	menuGroup := r.Group("/system/menu")
	menuGroup.POST("/queryPageByQo", controller.QueryMenuPageByQo)

	//分类管理
	categoryGroup := r.Group("/goods_management/goods_category")
	categoryGroup.POST("/queryPageByQo", controller.QueryCategoryPageByQo)
	categoryGroup.POST("/save", controller.SaveCategory)
	categoryGroup.POST("/update", controller.UpdateCategory)
	categoryGroup.GET("/normalCategoryAll", controller.NormalCategoryAll)
	categoryGroup.POST("/deactivate", controller.DeactivateCategory)

	//积分商品
	pointProductGroup := r.Group("/goods_management/point_products")
	pointProductGroup.GET("/queryOptionGoods", controller.QueryOptionGoods)
	pointProductGroup.POST("/queryPageByQo", controller.QueryPointPageByQo)
	pointProductGroup.GET("/del", controller.DelProductPoint)
	pointProductGroup.POST("/savePointGoods", controller.SavePointGoods)
	pointProductGroup.GET("/queryPointGoodsById", controller.QueryPointGoodsById)
	pointProductGroup.POST("/updatePointGoods", controller.UpdatePointGoods)

	//销售统计
	staticSaleGroup := r.Group("/goods_management/statistic_sale")
	staticSaleGroup.POST("/queryPageByQo", controller.QueryPageStatisticSaleByQo)
	return r
}
