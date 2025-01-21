package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
)

type RoleDao interface {
	GetAllRole() []model.Role
	QueryByEid(eid int64) []model.Role
	ClearRelationByEid(eid int64)
	ListRoleByQo(dto dto.RoleQueryDTO) []model.Role
	UpdateStateById(id int64)
	SelectRoleById(id int64) model.Role
	UpdateRole(role model.Role)
	SaveRole(role model.Role)
	SelectRoleByName(name string) model.Role
	SelectAllRole() []model.Role
	SelectAllRoleIds() []uint
	SelectRoleIdsByEid(eid int64) []uint
	ReSaveRelation(roleIds []string, eid int64)
	SaveEmpRole(id string, eid int64)
}

type RoleDaoImpl struct {
}

func (i RoleDaoImpl) SaveEmpRole(id string, eid int64) {
	db := common.GetDB()
	var empRole = model.EmpRole{
		RID: utils.ConvertStringToInt64(id),
		EID: eid,
	}
	db.Create(&empRole)
}

func (RoleDaoImpl) GetAllRole() []model.Role {
	db := common.GetDB()
	var roles []model.Role
	db.Find(&roles)
	return roles
}

func (RoleDaoImpl) QueryByEid(eid int64) []model.Role {
	db := common.GetDB()
	var roles []model.Role
	db.Table("t_role as r").Select("r.name,r.info,r.state").Joins("inner join t_emp_role as ter on r.id=ter.rid").Where("ter.eid = ?", eid).Find(&roles)
	return roles
}

func (RoleDaoImpl) ClearRelationByEid(eid int64) {
	db := common.GetDB()
	db.Table("t_emp_role").Where("eid = ?", eid).Delete(&model.Role{})
}

func (RoleDaoImpl) ListRoleByQo(dto dto.RoleQueryDTO) []model.Role {
	db := common.GetDB()
	name := dto.Name
	var roles []model.Role
	if name != "" && len(name) > 0 {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	state := dto.State
	if state != "" && len(state) > 0 {
		db = db.Where("state = ?", state)
	}
	db.Find(&roles)
	return roles
}

func (RoleDaoImpl) UpdateStateById(id int64) {
	db := common.GetDB()
	db.Where("id=?", id).Updates(&model.Role{
		State: common.STATE_BAN,
	})
}

func (RoleDaoImpl) SelectRoleById(id int64) model.Role {
	db := common.GetDB()
	var role = model.Role{}
	db.Where("id=?", id).First(&role)
	return role
}

func (RoleDaoImpl) UpdateRole(role model.Role) {
	db := common.GetDB()
	db.Updates(&role)
}

func (RoleDaoImpl) SaveRole(role model.Role) {
	db := common.GetDB()
	db.Create(&role)
}

func (RoleDaoImpl) SelectRoleByName(name string) model.Role {
	db := common.GetDB()
	var role = model.Role{}
	db.Where("name=?", name).First(&role)
	return role
}

func (RoleDaoImpl) SelectAllRole() []model.Role {
	db := common.GetDB()
	var roles = []model.Role{}
	intIds := []int{1, 2}
	db.Where("state = ? and id not in (?)", common.STATE_NORMAL, intIds).Find(&roles)
	return roles
}

func (RoleDaoImpl) SelectAllRoleIds() []uint {
	db := common.GetDB()
	var ids []uint
	intIds := []int{1, 2}
	db.Select("distinct id").Table("t_role").Where("state = ? and id not in (?)", common.STATE_NORMAL, intIds).Find(&ids)
	return ids
}

func (RoleDaoImpl) SelectRoleIdsByEid(eid int64) []uint {
	db := common.GetDB()
	var ids []uint
	intIds := []int{1, 2}
	db.Select("distinct rid").Table("t_emp_role").Where("eid= ?  and rid not in (?)", eid, intIds).Find(&ids)
	return ids
}

func (RoleDaoImpl) ReSaveRelation(roleIds []string, eid int64) {
	for _, id := range roleIds {
		SaveEmpRole(id, eid)
	}
}

func SaveEmpRole(id string, eid int64) {
	db := common.GetDB()
	var empRole = model.EmpRole{
		RID: utils.ConvertStringToInt64(id),
		EID: eid,
	}
	db.Create(&empRole)
}
