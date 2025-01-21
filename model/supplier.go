package model

// Supplier 代表供应商表的结构
type Supplier struct {
	Cn      int64  `gorm:"column:cn;primaryKey;autoIncrement;comment:编号" json:"cn"` // 使用uint64因为bigint通常是正数
	Name    string `gorm:"column:name;type:varchar(255);comment:名称" json:"name"`
	Address string `gorm:"column:address;type:varchar(255);comment:地址" json:"address"`
	Tel     string `gorm:"column:tel;type:varchar(255);comment:联系电话" json:"tel"`
	Info    string `gorm:"column:info;type:varchar(255);comment:备注" json:"info"`
	State   string `gorm:"column:state;type:char(2);comment:状态" json:"state"`
}

// TableName 指定了模型对应的数据库表名
func (Supplier) TableName() string {
	return "supplier"
}
