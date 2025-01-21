package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/model"
	"superMarket-backend/service"
)

var deptService service.IDeptService = &service.DeptServiceImpl{}

func List(c *gin.Context) {
	name := c.Query("name")
	state := c.Query("state")
	deptService.ListByQo(c, name, state)
}

func SaveDept(c *gin.Context) {
	var dept model.Department
	c.ShouldBind(&dept)
	deptService.SaveDept(c, dept)
}

func UpdateDept(c *gin.Context) {
	var dept model.Department
	c.ShouldBind(&dept)
	deptService.UpdateDept(c, dept)
}

func DeactivateDept(c *gin.Context) {
	var dept model.Department
	c.ShouldBind(&dept)
	deptService.DeactivateDept(c, dept.ID)
}
