package dto

// QueryGoods 代表商品查询条件的结构体
type QueryGoodsDTO struct {
	BaseQuery
	ID               uint64  `json:"id"`               // 使用uint64因为Java中的Long通常是正数
	Name             string  `json:"name"`             // 商品名称
	SellPrice        float64 `json:"sellPrice"`        // 销售价格
	CategoryID       uint64  `json:"categoryId"`       // 商品分类ID
	State            string  `json:"state"`            // 商品状态
	OperateStartTime string  `json:"operateStartTime"` // 操作开始时间
	OperateEndTime   string  `json:"operateEndTime"`   // 操作结束时间
}
