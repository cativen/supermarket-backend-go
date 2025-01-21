package model

import "time"

type SaleRecord struct {
	CN                int64              `gorm:"column:cn;primaryKey;size:255" json:"cn"`                  // 销售编号
	EID               int64              `gorm:"column:eid;type:bigint" json:"eid"`                        // 员工ID
	Sellway           string             `gorm:"column:sellway;size:2" json:"sellway"`                     // 销售方式
	SellTime          time.Time          `gorm:"column:sell_time;type:datetime" json:"sellTimeDo"`         // 销售时间
	State             string             `gorm:"column:state;size:2" json:"state"`                         // 状态 0:正常 1：删除
	Info              string             `gorm:"column:info;size:255" json:"info"`                         // 备注
	SellBy            string             `gorm:"column:sellby;size:255" json:"sellBy"`                     // 销售人员
	SellTotal         int64              `gorm:"column:sell_total;type:bigint" json:"sellTotal"`           // 销售总数量
	SellTotalMoney    float64            `gorm:"column:sell_totalmoney;type:double" json:"sellTotalmoney"` // 销售总金额
	Type              string             `gorm:"column:type;size:1" json:"type"`                           // 0:非会员消费 1：会员消费
	MemberPhone       string             `gorm:"column:member_phone;size:255" json:"memberPhone"`          // 顾客会员号码
	DetailSaleRecords []DetailSaleRecord `json:"detailSaleRecords" gorm:"-"`                               //销售详情
	SellTimeResp      string             `json:"sellTime" gorm:"-"`                                        // 销售时间
}

func (SaleRecord) TableName() string {
	return "t_sale_records"
}
