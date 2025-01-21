package vo

// DetailStoreGoodsOutVo 代表商品出库详情的视图对象
type DetailStoreGoodsOutVo struct {
	Cn         string `json:"cn"`         // 仓库编号
	GoodsID    int64  `json:"goodsId"`    // 商品ID
	GoodsNum   int64  `json:"goodsNum"`   // 商品数量
	GoodsName  string `json:"goodsName"`  // 商品名称
	CreateID   int64  `json:"createid"`   // 创建者ID
	CreateBy   string `json:"createby"`   // 创建者
	CreateTime string `json:"createTime"` // 创建时间
	State      string `json:"state"`      // 状态
	Info       string `json:"info"`       // 描述信息
	StoreID    int64  `json:"storeId"`    // 仓库编号
	StoreName  string `json:"storeName"`  // 仓库名称
	State1     string `json:"state1"`     // 状态1
}
