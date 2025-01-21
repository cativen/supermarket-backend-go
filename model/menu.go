package model

type Menu struct {
	ID          int64  `json:"id" db:"id"`                    // 主键
	Label       string `json:"label" db:"label"`              // 名称
	Purl        string `json:"purl" db:"purl"`                // 地址
	Type        string `json:"type" db:"type"`                // 类型 0:目录 1:菜单 2:按钮
	ParentID    int64  `json:"parentId" db:"parent_id"`       // 父id
	ParentLabel string `json:"parentLabel" db:"parent_label"` // 父名称
	Info        string `json:"info" db:"info"`                // 描述
	State       string `json:"state" db:"state"`              // 状态
	Flag        string `json:"flag" db:"flag"`                // 权限的唯一标识
	Icon        string `json:"icon" db:"icon"`                // 图标
	Component   string `json:"component" db:"component"`      // 组件路径
	Children    []Menu `json:"children" gorm:"-"`             // 子菜单列表
}

// TableName 指定 Member 结构体对应的数据库表名
func (Menu) TableName() string {
	return "t_menu"
}
