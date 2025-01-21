package model

import "time"

type Employee struct {
	ID             int64     `gorm:"column:id;" json:"id" db:"id"`                           // 主键
	Phone          string    `gorm:"column:phone;" json:"phone" db:"phone"`                  // 用户名
	Email          string    `gorm:"column:email;" json:"email" db:"email"`                  // 邮箱
	Address        string    `gorm:"column:address;" json:"address" db:"address"`            // 住址
	Sex            string    `gorm:"column:sex;" json:"sex" db:"sex"`                        // 性别
	Password       string    `gorm:"column:password;" json:"password" db:"password"`         // 密码
	NickName       string    `gorm:"column:nick_name;" json:"nickName" db:"nick_name"`       // 昵称
	HeadImg        string    `gorm:"column:head_img;" json:"headImg" db:"head_img"`          // 头像
	State          string    `gorm:"column:state;" json:"state" db:"state"`                  // 状态 0：在职 1：离职
	IsAdmin        bool      `gorm:"column:is_admin;" json:"isAdmin" db:"is_admin"`          // 是否是超管 1:是 0:不是
	Info           string    `gorm:"column:info;" json:"info" db:"info"`                     // 描述
	CreateBy       string    `gorm:"column:createby;" json:"createBy" db:"createby"`         // 创建者
	CreateTime     time.Time `gorm:"column:create_time;" json:"CreateTime" db:"create_time"` // 创建时间
	Age            int       `gorm:"column:age;" json:"age" db:"age"`                        // 年龄
	DeptID         uint      `gorm:"column:deptId;" json:"deptId" db:"deptId"`               // 部门主键
	IDCard         string    `gorm:"column:id_card;" json:"idCard" db:"id_card"`             // 身份证号
	LeaveTime      time.Time `gorm:"column:leave_time;" json:"LeaveTime" db:"leave_time"`    // 离职时间
	Menus          []Menu    `json:"menus" gorm:"-"`                                         // 菜单列表
	UserName       string    `json:"username" gorm:"-"`                                      // 菜单列表
	DepartName     string    `json:"deptName" gorm:"-"`
	CreateTimeResp string    `gorm:"-" json:"createTime" `
	LeaveTimeResp  string    `gorm:"-" json:"leaveTime"`
}

// TableName 指定 Member 结构体对应的数据库表名
func (Employee) TableName() string {
	return "employee"
}
