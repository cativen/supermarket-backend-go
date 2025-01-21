package model

// GoodsStore 代表数据库中的t_goods_store表
type GoodsStore struct {
	GoodsID    uint   `gorm:"column:goods_id;primaryKey" json:"goodsId"`                     // 商品编号
	StoreID    uint   `gorm:"column:store_id;not null" json:"storeId"`                       // 仓库编号
	InNum      int64  `gorm:"column:in_num;not null" json:"inNum"`                           // 入库数数量
	ResidueNum int64  `gorm:"column:residue_num;not null" json:"residueNum"`                 // 剩余数量
	StoreName  string `gorm:"column:store_name;type:varchar(255);not null" json:"storeName"` // 仓库名
}

func (GoodsStore) TableName() string {
	return "t_goods_store"
}
