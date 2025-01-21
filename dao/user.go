package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type UserDao interface {
	SelectOneByUsername(username string) model.Employee
	SelectById(uid int64) model.Employee
	DeleteById(uid int64)
	DeleteByUserName(username string)
	SelectPageByQo(dto dto.EmpQueryDTO) vo.Page[model.Employee]
	SelectCountByPhoneOrIdCard(id int64, phone string, card string) int64
	UpdateEmployee(employee model.Employee)
	UpdateEmployeePwd(pwd string, phone string)
}

type UserDaoImpl struct {
}

func (useDao *UserDaoImpl) SelectOneByUsername(username string) model.Employee {
	db := common.GetDB()
	var emp model.Employee
	db.Where("phone = ?", username).First(&emp)
	return emp
}

func (useDao *UserDaoImpl) SelectById(uid int64) model.Employee {
	db := common.GetDB()
	var emp model.Employee
	db.Where("id = ?", uid).First(&emp)
	return emp
}

func (useDao *UserDaoImpl) DeleteById(uid int64) {
	db := common.GetDB()
	db.Where("id = ?", uid).Delete(&model.Employee{})
}

func (useDao *UserDaoImpl) DeleteByUserName(username string) {
	db := common.GetDB()
	db.Where("phone = ?", username).Delete(&model.Employee{})
}

func (useDao *UserDaoImpl) SelectPageByQo(dto dto.EmpQueryDTO) vo.Page[model.Employee] {
	db := common.GetDB()
	db = db.Table("employee").Select("*")
	var page vo.Page[model.Employee]
	age := dto.Age
	if age != "" && len(age) > 0 {
		db = db.Where("age = ?", age)
	}

	sex := dto.Sex
	if sex != "" && len(sex) > 0 {
		db = db.Where("sex = ?", sex)
	}

	departId := dto.DeptId
	if departId > 0 {
		db = db.Where("deptId = ?", departId)
	}

	username := dto.Username
	if username != "" && len(username) > 0 {
		db = db.Where("phone LIKE ?", "%"+username+"%")
	}

	nickName := dto.NickName
	if username != "" && len(username) > 0 {
		db = db.Where("nick_name LIKE ?", "%"+nickName+"%")
	}

	address := dto.Address
	if address != "" && len(address) > 0 {
		db = db.Where("address LIKE ?", "%"+address+"%")
	}
	db = db.Where("id <> ?", 1)
	var models []model.Employee
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, models)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (useDao *UserDaoImpl) SelectCountByPhoneOrIdCard(id int64, phone string, card string) int64 {
	var count int64
	db := common.GetDB()
	db = db.Table("employee").Select("*").Where("id != ?", id)
	if phone != "" && len(phone) > 0 {
		db = db.Where("(phone = ?", phone)
	}

	if card != "" && len(card) > 0 {
		db = db.Or("id_card = ? )", card)
	}

	db.Count(&count)
	return count
}

func (useDao *UserDaoImpl) UpdateEmployee(employee model.Employee) {
	db := common.GetDB()
	db.Updates(&employee)
}

func (useDao *UserDaoImpl) UpdateEmployeePwd(pwd string, phone string) {
	db := common.GetDB()
	var employee = model.Employee{
		Password: pwd,
	}
	db.Where("phone = ?", phone).Updates(employee)
}
