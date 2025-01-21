package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
	"strconv"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
	"time"
)

var userDao dao.UserDao = &dao.UserDaoImpl{}
var departDao dao.DeptDao = &dao.DeptDaoImpl{}

type IEmployeeService interface {
	Detail(uid int64) vo.DetailEmpVo
	Information(c *gin.Context, token string)
	PageEmployeeByQo(c *gin.Context, dto dto.EmpQueryDTO)
	UpdateEmployee(c *gin.Context, employee model.Employee, token string)
	GetEmpById(c *gin.Context, uid string)
	DeactivateEmp(c *gin.Context, id int64)
	ResetEmpPwd(c *gin.Context, eid string, code string)
	EditUserPwd(c *gin.Context, pwdDTO dto.QueryEditPwdDTO, token string)
}

type EmployeeServiceImpl struct {
}

func (employeeService *EmployeeServiceImpl) Detail(uid int64) vo.DetailEmpVo {
	var vo vo.DetailEmpVo
	//查询员工信息
	var userDao dao.UserDao = &dao.UserDaoImpl{}
	employee := userDao.SelectById(uid)
	copier.Copy(&vo, &employee)
	vo.CreateTime = utils.FormatTimeToString(employee.CreateTime)
	vo.LeaveTime = utils.FormatTimeToString(employee.LeaveTime)
	//补全角色信息
	roleNames := []string{}
	var roleDao dao.RoleDao = &dao.RoleDaoImpl{}
	if employee.IsAdmin == true {
		//查询所有角色
		role := roleDao.GetAllRole()
		for _, val := range role {
			roleNames = append(roleNames, val.Name)
		}
	} else {
		role := roleDao.QueryByEid(uid)
		for _, val := range role {
			roleNames = append(roleNames, val.Name)
		}
	}
	vo.RoleNames = roleNames
	return vo
}

func (employeeService *EmployeeServiceImpl) Information(c *gin.Context, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println(err)
	}
	var employee model.Employee
	if err := json.Unmarshal([]byte(val), &employee); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}
	impl := EmployeeServiceImpl{}
	detail := impl.Detail(employee.ID)
	var vo vo.InformationVo
	copier.Copy(&vo, &detail)
	vo.Username = vo.Phone
	response.Success(c, vo, "操作成功")
}

func (employeeService *EmployeeServiceImpl) PageEmployeeByQo(c *gin.Context, dto dto.EmpQueryDTO) {
	page := userDao.SelectPageByQo(dto)
	//补全部门信息
	depts := departDao.SelectDeptByQo(model.Department{})
	var deptMap = make(map[uint]model.Department)
	for _, dept := range depts {
		deptMap[dept.ID] = dept
	}

	records := page.Records
	for index, record := range records {
		department := deptMap[record.DeptID]
		if department != (model.Department{}) {
			records[index].DepartName = department.Name
		}
		records[index].UserName = records[index].Phone
		records[index].CreateTimeResp = utils.FormatTimeToString(records[index].CreateTime)
		records[index].LeaveTimeResp = utils.FormatTimeToString(records[index].LeaveTime)
	}
	page.Records = records
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	response.Success(c, page, "操作成功")
}

func (employeeService *EmployeeServiceImpl) UpdateEmployee(c *gin.Context, employee model.Employee, token string) {
	if employee.State == common.STATE_DEL {
		employee.LeaveTime = time.Now()
		if employee.IsAdmin {
			response.Error(c, "不可以给系统管理者办理离职")
			return
		}
	} else {
		rdb := common.GetRDB()
		ctx := context.Background()
		val, err := rdb.Get(ctx, token).Result()
		if err != nil {
			log.Println(err)
		}
		var existEmployee model.Employee
		if err := json.Unmarshal([]byte(val), &existEmployee); err != nil {
			response.Error(c, "token已过期需要重新登录")
			return
		}
		employee.CreateBy = existEmployee.NickName
	}

	//查询是否存在相同的用户名或身份证号
	count := userDao.SelectCountByPhoneOrIdCard(employee.ID, employee.Phone, employee.IDCard)
	if count > 0 {
		response.Error(c, "系统已存在相同的用户名或身份证号")
		return
	}
	userDao.UpdateEmployee(employee)
}

func (employeeService *EmployeeServiceImpl) GetEmpById(c *gin.Context, uid string) {
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		panic(err)
	}
	byId := userDao.SelectById(id)
	byId.UserName = byId.Phone
	response.Success(c, byId, "操作成功")
}

func (employeeService *EmployeeServiceImpl) DeactivateEmp(c *gin.Context, id int64) {
	byId := userDao.SelectById(id)
	if byId.IsAdmin {
		response.Error(c, "不可以给系统管理者办理离职")
		return
	}
	if common.STATE_DEL == byId.State {
		response.Error(c, "已是离职状态")
		return
	}
	var employee = model.Employee{
		ID:        id,
		State:     common.STATE_DEL,
		LeaveTime: time.Now(),
	}
	userDao.UpdateEmployee(employee)
}

func (employeeService *EmployeeServiceImpl) ResetEmpPwd(c *gin.Context, eid string, code string) {
	id, err := strconv.ParseInt(eid, 10, 64)
	if err != nil {
		panic(err)
	}
	byId := userDao.SelectById(id)
	if byId.ID == 1 {
		response.Error(c, "该账户不可被修改")
		return
	}
	var employee = model.Employee{
		ID:       id,
		Password: common.DEFAULT_PWD,
	}
	if byId.IsAdmin {
		if code == "123456" {
			userDao.UpdateEmployee(employee)
		} else {
			response.Error(c, "密钥错误")
			return
		}
	} else {
		if code == "456789" {
			userDao.UpdateEmployee(employee)
		} else {
			response.Error(c, "密钥错误")
			return
		}
	}
	employee.UserName = employee.Phone
	rdb := common.GetRDB()
	ctx := context.Background()
	errorPassKey := fmt.Sprintf("%s%s", common.LOGIN_ERRO_PWDNUM, employee.UserName)
	disableKey := fmt.Sprintf("%s%s", common.DISABLEUSER, employee.UserName)
	loginUserKey := fmt.Sprintf("%s%s", common.LOGIN_USER, employee.UserName)
	rdb.Del(ctx, errorPassKey)
	rdb.Del(ctx, disableKey)
	rdb.Del(ctx, loginUserKey)
	response.Success(c, nil, "操作成功")
}

func (employeeService *EmployeeServiceImpl) EditUserPwd(c *gin.Context, pwdDTO dto.QueryEditPwdDTO, token string) {
	//获取缓存中的登录员工信息
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println(err)
	}
	var existEmployee model.Employee
	if err := json.Unmarshal([]byte(val), &existEmployee); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}
	if existEmployee.ID == 1 {
		response.Error(c, "该账户不能被修改")
		return
	}

	//比对用户名是否一致
	if existEmployee.Phone != pwdDTO.UserName {
		response.Error(c, "没有权限修改其他人的密码")
		return
	}

	//比对旧密码是否输入正确
	if existEmployee.Password != pwdDTO.OldPwd {
		response.Error(c, "原密码输入有误")
		return
	}
	//比对新密码和旧密码是否一致
	if pwdDTO.NewPwd == pwdDTO.OldPwd {
		response.Error(c, "新密码和旧密码一致")
		return
	}
	userDao.UpdateEmployeePwd(pwdDTO.NewPwd, existEmployee.Phone)

	//发送邮件通知
	utils.SendMessage("修改密码", fmt.Sprintf("亲爱的%s,你的密码已被修改成%s", existEmployee.NickName, pwdDTO.NewPwd), existEmployee.Email)
	response.Success(c, nil, "操作成功")
}
