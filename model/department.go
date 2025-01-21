package model

type Department struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"` // 主键
	Name  string `gorm:"size:20" json:"name"`                // 部门名称
	Info  string `gorm:"size:255" json:"info"`               // 描述
	State string `gorm:"size:2" json:"state"`                // 状态
}

// TableName 指定 Member 结构体对应的数据库表名
func (Department) TableName() string {
	return "department"
}
