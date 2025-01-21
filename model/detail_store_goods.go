package model

import "time"

// DetailStoreGoods 代表数据库中的t_detail_store_goods表
type DetailStoreGoods struct {
	Cn           string     `gorm:"column:cn;primaryKey" json:"cn"`           // 主键
	GoodsID      int64      `gorm:"column:goods_id" json:"goodsId"`           // 商品ID
	GoodsNum     int64      `gorm:"column:goods_num" json:"goodsNum"`         // 商品数量
	GoodsName    string     `gorm:"column:goods_name" json:"goodsName"`       // 商品名称
	GoodsPrice   float64    `gorm:"column:goods_price" json:"goodsPrice"`     // 商品价格
	Type         string     `gorm:"column:type" json:"type"`                  // 类型
	CreateID     int64      `gorm:"column:createid" json:"createId"`          // 创建者ID
	CreateTime   time.Time  `gorm:"column:create_time" json:"createTime"`     // 创建时间
	State        string     `gorm:"column:state" json:"state"`                // 状态
	Info         string     `gorm:"column:info" json:"info"`                  // 描述信息
	ExpiryTime   *time.Time `gorm:"column:expiry_time" json:"expiryTime"`     // 过期时间
	CreateBy     string     `gorm:"column:createby" json:"createBy"`          // 创建者
	BirthTime    *time.Time `gorm:"column:birth_time" json:"birthTime"`       // 生产时间
	State1       string     `gorm:"column:state1" json:"state1"`              // 状态1
	StoreID      uint       `gorm:"column:store_id" json:"storeId"`           // 仓库编号
	SupplierID   int64      `gorm:"column:supplier_id" json:"supplierId"`     // 供货商编号
	SupplierName *string    `gorm:"column:supplier_name" json:"supplierName"` // 供货商名称
	UntreatedNum int64      `gorm:"column:untreated_num" json:"untreatedNum"` // 待处理数量
}

func (DetailStoreGoods) TableName() string {
	return "t_detail_store_goods"
}
