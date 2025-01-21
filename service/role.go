package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
)

var roleDao dao.RoleDao = dao.RoleDaoImpl{}

type IRoleService interface {
	ClearEmpPermission(uid int64)
	RoleList(c *gin.Context, dto dto.RoleQueryDTO)
	ForbiddenRole(c *gin.Context, id int64)
	EditRole(c *gin.Context, role model.Role)
	SaveRole(c *gin.Context, role model.Role)
	AllRole(c *gin.Context)
	QueryRoleIdsByEid(c *gin.Context, eid int64)
	SaveRoleEmp(c *gin.Context, eid string, ids []string, token string)
	ExportRoleExcel(c *gin.Context)
}

type RoleServiceImpl struct {
}

func (roleService RoleServiceImpl) ClearEmpPermission(uid int64) {
	roleDao.ClearRelationByEid(uid)
}

func (roleService RoleServiceImpl) RoleList(c *gin.Context, dto dto.RoleQueryDTO) {
	list := roleDao.ListRoleByQo(dto)
	response.Success(c, list, "操作成功")
}

func (roleService RoleServiceImpl) ForbiddenRole(c *gin.Context, id int64) {
	byId := roleDao.SelectRoleById(id)
	if byId.ID == 1 || byId.ID == 2 {
		response.Error(c, "不能停用系统拥有者")
		return
	}
	roleDao.UpdateStateById(id)
	response.Success(c, "success", "操作成功")
}

func (roleService RoleServiceImpl) EditRole(c *gin.Context, role model.Role) {
	if role.ID == 1 || role.ID == 2 {
		response.Error(c, "不能停用系统拥有者")
		return
	}
	roleDao.UpdateRole(role)
	response.Success(c, "success", "操作成功")
}

func (roleService RoleServiceImpl) SaveRole(c *gin.Context, role model.Role) {
	if role == (model.Role{}) {
		response.Error(c, "操作失败")
		return
	}

	if role.Name != "" && len(role.Name) > 0 {
		existRole := roleDao.SelectRoleByName(role.Name)
		if existRole != (model.Role{}) {
			response.Error(c, "操作失败，角色名重复")
			return
		} else {
			role.State = common.STATE_NORMAL
			roleDao.SaveRole(role)
			response.Success(c, "success", "操作成功")
			return
		}
	} else {
		response.Error(c, "角色名称格式有误")
		return
	}

}

func (roleService RoleServiceImpl) AllRole(c *gin.Context) {
	roles := roleDao.SelectAllRole()
	var vos = make([]map[string]interface{}, len(roles))
	for index, val := range roles {
		var roleMap = make(map[string]interface{})
		roleMap["id"] = val.ID
		roleMap["label"] = val.Name
		vos[index] = roleMap
	}
	response.Success(c, vos, "操作成功")
}

func (roleService RoleServiceImpl) QueryRoleIdsByEid(c *gin.Context, eid int64) {
	var userDao dao.UserDao = &dao.UserDaoImpl{}
	employee := userDao.SelectById(eid)
	if employee.IsAdmin {
		//查询出所有的ID
		ids := roleDao.SelectAllRoleIds()
		response.Success(c, ids, "操作成功")
		return
	} else {
		//查询出当前员工的角色ID
		ids := roleDao.SelectRoleIdsByEid(eid)
		response.Success(c, ids, "操作成功")
		return
	}
}

func (roleService RoleServiceImpl) SaveRoleEmp(c *gin.Context, eid string, roleIds []string, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println(err)
	}
	var emp model.Employee
	if err := json.Unmarshal([]byte(val), &emp); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}
	eidVal := utils.ConvertStringToInt64(eid)
	if emp.ID == eidVal {
		response.Error(c, "无法为自己赋予职务")
		return
	}
	//查询用户的信息，判断是否是系统管理员
	employee := userDao.SelectById(eidVal)
	if employee.IsAdmin {
		response.Error(c, "无法操作系统管理员的职务")
		return
	}
	//根据员工编号清除关系
	roleDao.ClearRelationByEid(eidVal)

	//重新保存关系
	if len(roleIds) > 0 {
		roleDao.ReSaveRelation(roleIds, eidVal)
	}
}

func (roleService RoleServiceImpl) ExportRoleExcel(c *gin.Context) {
	list := roleDao.SelectAllRole()

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("角色信息")
	titles := []string{"ID", "角色名", "描述", "状态"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, order := range list {
		values := []string{
			strconv.Itoa(int(order.ID)),
			order.Name,
			order.Info,
		}

		// 根据 order.ID 的值添加 "停用" 或 "正常"
		if order.State == "0" {
			values = append(values, "正常")
		} else {
			values = append(values, "停用")
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	filename := "角色信息.xlsx"

	// 使用 URL 编码处理文件名
	encodedFilename := url.QueryEscape(filename)

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+encodedFilename)
	c.Writer.Header().Set("Content-Transfer-Encoding", "binary")

	// 回写到 web 流媒体 形成下载
	if err := file.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "导出失败"})
		return
	}
}
