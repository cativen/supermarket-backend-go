package dto

type QueryGoodsCategoryDTO struct {
	BaseQuery
	Name  string `json:"name"`  // 商品名称
	State string `json:"state"` // 状态
}
