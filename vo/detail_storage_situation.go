package vo

type DetailStorageSituationVo struct {
	GoodsId    uint    `json:"goodsId"`    // 仓库编号
	GoodsName  string  `json:"goodsName"`  // 商品ID
	ResidueNum int64   `json:"residueNum"` // 商品数量
	Percentage float64 `json:"percentage"` // 商品名称
}
