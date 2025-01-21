package dto

type SaleRecordPageDTO struct {
	BaseQuery            // 嵌入基础查询结构体
	Cn            string `json:"cn"`            // 销售编号
	StartSellTime string `json:"startSellTime"` // 开始销售时间
	EndSellTime   string `json:"endSellTime"`   // 结束销售时间
	Type          string `json:"type"`          // 类型
	Sellway       string `json:"sellway"`       // 销售方式
}
