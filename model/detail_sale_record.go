package model

type DetailSaleRecord struct {
	SellCN       int64   `gorm:"primaryKey;size:255"  json:"sellCn" db:"sell_cn"` // 销售订单号
	GoodsID      int64   `json:"goodsId" db:"goods_id"`                           // 商品编号
	GoodsNum     int64   `json:"goodsNumNumber" db:"goods_num"`                   // 商品数量
	GoodsPrice   float64 `json:"goodsPrice" db:"goods_price"`                     // 销售单价
	GoodsName    string  `gorm:"size:255" json:"goodsName" db:"goods_name"`       // 商品名
	GoodsNumJson string  `json:"goodsNum" db:"-" gorm:"-"`                        // 商品数量
}

// TableName 指定 Member 结构体对应的数据库表名
func (DetailSaleRecord) TableName() string {
	return "detail_sale_records"
}
