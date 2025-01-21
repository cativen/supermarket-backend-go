package dto

type QueryPointProductsDTO struct {
	BaseQuery
	Name string `json:"name"` // 商品名称
}
