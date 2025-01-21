package service

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/common"
	"superMarket-backend/model"
	"superMarket-backend/response"
)

type IDeptService interface {
	ListByQo(c *gin.Context, name string, state string)
	SaveDept(c *gin.Context, dept model.Department)
	UpdateDept(c *gin.Context, dept model.Department)
	DeactivateDept(c *gin.Context, id uint)
}

type DeptServiceImpl struct {
}

func (deptService *DeptServiceImpl) ListByQo(c *gin.Context, name string, state string) {
	db := common.GetDB()
	var deptList []model.Department
	if len(name) > 0 {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	if len(state) > 0 {
		db = db.Where("state = ?", state)
	}
	db.Find(&deptList)
	response.Success(c, deptList, "操作成功")
}

func (deptService *DeptServiceImpl) SaveDept(c *gin.Context, dept model.Department) {
	name := dept.Name
	if len(name) > 0 && name != "" {
		db := common.GetDB()
		var existDept model.Department
		db.Where("name = ?", name).Find(&existDept)
		if existDept.ID != 0 {
			response.Error(c, "操作失败，该部门已存在")
			return
		}
	}
	dept.State = common.STATE_NORMAL
	db := common.GetDB()
	db.Create(&dept)
	response.Success(c, "success", "操作成功")
}

func (deptService *DeptServiceImpl) UpdateDept(c *gin.Context, dept model.Department) {
	db := common.GetDB()
	if common.STATE_BAN == dept.State {
		if dept.ID != 0 {
			var list []model.Employee
			db.Where("deptId = ?", dept.ID).Find(&list)
			if len(list) > 0 {
				response.Error(c, "操作失败，该部门正在使用")
				return
			}
		}
	}
	db.Updates(dept)
	response.Success(c, "success", "操作成功")
}

func (deptService *DeptServiceImpl) DeactivateDept(c *gin.Context, id uint) {
	db := common.GetDB()
	if id != 0 {
		var list []model.Employee
		db.Where("deptId = ?", id).Find(&list)
		if len(list) > 0 {
			response.Error(c, "操作失败，该部门正在使用")
			return
		}
	}
	var dept = &model.Department{
		State: common.STATE_BAN,
	}
	db.Where("id =?", id).Updates(dept)
	response.Success(c, "success", "操作成功")
}
