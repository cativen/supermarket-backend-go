package dto

type QuerySupplierDTO struct {
	BaseQuery
	Name    string `json:"name"`    // 名字
	Address string `json:"address"` // 名字
	Info    string `json:"info"`    // 名字
}
