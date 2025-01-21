package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/service"
)

func Information(c *gin.Context) {
	token := c.GetHeader("token")
	var employeeService service.IEmployeeService = &service.EmployeeServiceImpl{}
	employeeService.Information(c, token)
}

func EditUserPwd(c *gin.Context) {
	var dto dto.QueryEditPwdDTO
	c.ShouldBind(&dto)
	token := c.GetHeader("token")
	var employeeService service.IEmployeeService = &service.EmployeeServiceImpl{}
	employeeService.EditUserPwd(c, dto, token)
}
