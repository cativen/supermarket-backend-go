package model

// EmpRole 定义员工角色关联信息结构体
type EmpRole struct {
	ID  int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // 主键
	EID int64 `gorm:"column:eid" json:"eid"`                        // 用户id
	RID int64 `gorm:"column:rid" json:"rid"`                        // 角色id
}

// TableName 指定 EmpRole 结构体对应的数据库表名
func (EmpRole) TableName() string {
	return "t_emp_role"
}
