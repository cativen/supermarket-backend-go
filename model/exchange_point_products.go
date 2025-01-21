package model

import "time"

type ExchangePointProductsRecord struct {
	CN            string    `gorm:"column:cn;size:255" json:"cn"`                     // 订单号
	GoodsID       int64     `gorm:"column:goods_id;" json:"goodsId"`                  // 商品编号
	MemberID      int64     `gorm:"column:member_id;" json:"memberId"`                // 会员编号
	Integral      int64     `gorm:"column:integral;" json:"integral"`                 // 积分
	UpdateTime    time.Time `gorm:"column:update_time;" json:"UpdateTime"`            // 最近操作时间
	UpdateBy      string    `gorm:"column:updateBy;" gorm:"size:255" json:"updateBy"` // 操作者
	UpdateID      int64     `gorm:"column:update_id;" json:"updateId"`                // 操作者编号
	State         string    `gorm:"column:state;" json:"state"`                       // 状态
	MemberPhone   string    `gorm:"-" json:"memberPhone"`                             // 会员手机号码
	GoodsName     string    `gorm:"-" json:"goodsName"`                               // 商品名称
	GoodsCoverUrl string    `gorm:"-" json:"goodsCoverUrl"`                           // 商品封面图
	UpdateTimeRes string    `gorm:"-" json:"updateTime"`
}

// TableName 指定 Member 结构体对应的数据库表名
func (ExchangePointProductsRecord) TableName() string {
	return "exchange_point_products_records"
}
