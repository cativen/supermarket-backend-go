package dto

type MenuQueryDTO struct {
	BaseQuery
	Name string `json:"name"` // 用户名
}
