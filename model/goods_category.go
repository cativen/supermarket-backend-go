package model

type GoodsCategory struct {
	ID    uint   `gorm:"primaryKey;column:id;autoIncrement;comment:主键" json:"id"` // 使用uint64因为bigint通常是正数
	Name  string `gorm:"type:varchar(255);column:name;comment:分类名" json:"name"`
	Info  string `gorm:"type:varchar(255);column:info;comment:备注" json:"info"`
	State string `gorm:"type:char(2);column:state;comment:状态" json:"state"`
}

// TableName 指定了模型对应的数据库表名
func (GoodsCategory) TableName() string {
	return "goods_category"
}
