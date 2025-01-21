package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/service"
)

var roleService service.IRoleService = &service.RoleServiceImpl{}

func RoleList(c *gin.Context) {
	var dto dto.RoleQueryDTO
	c.ShouldBind(&dto)
	roleService.RoleList(c, dto)
}

func ForbiddenRole(c *gin.Context) {
	var dto dto.RoleQueryDTO
	c.ShouldBind(&dto)
	roleService.ForbiddenRole(c, dto.RoleId)
}

func EditRole(c *gin.Context) {
	var role model.Role
	c.ShouldBind(&role)
	roleService.EditRole(c, role)
}

func SaveRole(c *gin.Context) {
	var role model.Role
	c.ShouldBind(&role)
	roleService.SaveRole(c, role)
}

func CheckRolePermissions(c *gin.Context) {
	cn := c.Query("cn")
	var saleRecordService service.ISaleRecordsService = &service.SaleRecordsServiceImpl{}
	// 确保服务方法正确处理错误
	saleRecordService.DelSaleRecords(c, cn)
}

func SaveRolePermissions(c *gin.Context) {
	cn := c.Query("cn")
	var saleRecordService service.ISaleRecordsService = &service.SaleRecordsServiceImpl{}
	// 确保服务方法正确处理错误
	saleRecordService.DelSaleRecords(c, cn)
}

func AllRole(c *gin.Context) {
	roleService.AllRole(c)
}

func QueryRoleIdsByEid(c *gin.Context) {
	eid := c.Query("eid")
	num, err := strconv.Atoi(eid)
	if err != nil {
		log.Println("Error:", err)
	} else {
		// 手动转为 int64
		num64 := int64(num)
		roleService.QueryRoleIdsByEid(c, num64)
	}
}

func SaveRoleEmp(c *gin.Context) {
	eid, _ := c.GetPostForm("eid")
	empRoleIdsStr, _ := c.GetPostForm("empRoleIds")
	empRoleIds := strings.Split(empRoleIdsStr, ",")
	token := c.GetHeader("token")
	var roleService service.IRoleService = &service.RoleServiceImpl{}
	roleService.SaveRoleEmp(c, eid, empRoleIds, token)
	response.Success(c, nil, "操作成功")
}

func ExportRoleExcel(c *gin.Context) {
	roleService.ExportRoleExcel(c)
}
