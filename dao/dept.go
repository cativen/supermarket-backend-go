package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/model"
)

type DeptDao interface {
	SelectDeptByQo(dept model.Department) []model.Department
}

type DeptDaoImpl struct {
}

func (deptDao DeptDaoImpl) SelectDeptByQo(dept model.Department) []model.Department {
	db := common.GetDB()
	var depts []model.Department
	db.Where(&dept).Find(&depts)
	return depts
}
