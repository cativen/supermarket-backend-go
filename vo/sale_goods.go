package vo

type SaleGoodsVo struct {
	GoodsId     uint64 `json:"goodsId"`     // 商品ID
	GoodsName   string `json:"goodsName"`   // 商品名称
	CoverUrl    string `json:"coverUrl"`    // 封面URL
	SalesVolume int64  `json:"salesVolume"` // 销量
	Percentage  int64  `json:"percentage"`  // 百分比
}
