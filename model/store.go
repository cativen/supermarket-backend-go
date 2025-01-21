package model

type Store struct {
	ID      uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                   // 主键
	Name    string `gorm:"column:name;type:varchar(255);uniqueIndex;not null" json:"name"` // 仓库名称
	Address string `gorm:"column:address;type:varchar(255);" json:"address"`               // 仓库地址
	State   string `gorm:"column:state;type:char(2);not null" json:"state"`                // 状态
	Info    string `gorm:"column:info;type:varchar(255);" json:"info"`                     // 描述
}

func (Store) TableName() string {
	return "store"
}
