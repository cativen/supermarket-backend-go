package dto

type QueryExchangePointProductsRecordsDTO struct {
	BaseQuery        // 嵌入基础查询结构体
	Cn        string `json:"cn"`        // 兑换记录编号
	MemberId  int64  `json:"memberId"`  // 会员ID
	StartTime string `json:"startTime"` // 开始时间
	EndTime   string `json:"endTime"`   // 结束时间
}
