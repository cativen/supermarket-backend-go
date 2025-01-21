package model

import "time"

// PointProduct 映射数据库中的 point_products 表
type PointProduct struct {
	GoodsID    uint      `gorm:"index;comment:'商品ID';column:goods_id" json:"goodsId"`
	GoodsName  string    `gorm:"type:varchar(255);character set utf8mb3;collate utf8mb3_bin;comment:'商品名称';column:goods_name" json:"goodsName"`
	Integral   uint      `gorm:"comment:'积分';column:integral" json:"integral"`
	UpdateBy   string    `gorm:"type:varchar(255);character set utf8mb3;collate utf8mb3_bin;comment:'更新者';column:updateby" json:"updateby"`
	UpdateTime time.Time `gorm:"type:datetime;comment:'更新时间';column:update_time" json:"updateTime"`
	UpdateID   uint      `gorm:"comment:'更新者ID';column:update_id" json:"updateId"`
	CoverURL   string    `gorm:"type:varchar(255);character set utf8mb3;collate utf8mb3_bin;comment:'封面URL';column:cover_url" json:"coverUrl"`
	State      string    `gorm:"type:char(2);character set utf8mb3;collate utf8mb3_bin;not null;comment:'状态';column:state" json:"state"`
}

func (PointProduct) TableName() string {
	return "point_products"
}
