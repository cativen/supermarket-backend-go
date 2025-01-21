package model

type Role struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id" db:"id"` // 主键
	Name  string `gorm:"size:255;index" json:"name" db:"name"`       // 角色名
	Info  string `gorm:"size:255" json:"info" db:"info"`             // 描述
	State string `gorm:"size:2;index" json:"state" db:"state"`       // 状态 0：正常 -1：停用
}

// TableName 指定 Member 结构体对应的数据库表名
func (Role) TableName() string {
	return "t_role"
}
