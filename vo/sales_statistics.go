package vo

type SalesStatisticsVo struct {
	Total int64             `json:"total"`
	VOS   Page[SaleGoodsVo] `json:"vos"`
}
