package dto

type QueryDetailStoreGoodsOutDTO struct {
	BaseQuery
	Cn              string `json:"cn"`              // 仓库编号
	GoodsName       string `json:"goodsName"`       // 商品名称
	State1          string `json:"state1"`          // 状态1
	State           string `json:"state"`           // 状态
	StartCreateTime string `json:"startCreateTime"` // 开始创建时间
	EndCreateTime   string `json:"endCreateTime"`   // 结束创建时间
}
