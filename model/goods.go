package model

import "time"

// Goods 定义商品信息结构体
type Goods struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                 // 主键
	Name          string    `gorm:"size:255;column:name" json:"name"`                             // 商品名
	CreateBy      string    `gorm:"size:20;column:createby" json:"createBy"`                      // 创建者
	CreateTime    time.Time `gorm:"size:20;column:create_time" json:"createTime"`                 // 创建时间
	CategoryID    uint      `gorm:"category_id" json:"categoryId"`                                // 商品分类id
	SellPrice     float64   `gorm:"type:double(10,2);column:sell_price" json:"sellPrice"`         // 销售价格
	PurchasePrice float64   `gorm:"type:double(10,2);column:purchash_price" json:"purchashPrice"` // 进货价格
	UpdateTime    time.Time `gorm:"update_time" json:"updateTime"`                                // 更改时间
	UpdateBy      string    `gorm:"size:255;column:updateby" json:"updateBy"`                     // 更改者
	CategoryName  string    `gorm:"size:255;column:category_name" json:"categoryName"`            // 分类名
	CoverURL      string    `gorm:"size:255;column:cover_url" json:"coverUrl"`                    // 商品封面
	State         string    `gorm:"size:2;column:state" json:"state"`                             // 状态
	ResidueNum    int64     `gorm:"residue_num" json:"residueNum"`                                // 剩余数量
	Info          string    `gorm:"size:255;column:info" json:"info"`                             // 备注
	SalesVolume   int64     `gorm:"sales_volume" json:"salesVolume"`                              // 销量
	Inventory     int64     `gorm:"inventory" json:"inventory"`                                   // 需库存量
	Shelves       int64     `gorm:"shelves" json:"shelves"`                                       // 货架上需摆放的数量
}

// TableName 指定 Member 结构体对应的数据库表名
func (Goods) TableName() string {
	return "goods"
}
