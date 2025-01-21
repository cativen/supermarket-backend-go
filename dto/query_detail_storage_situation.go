package dto

type QueryDetailStorageSituationDTO struct {
	BaseQuery
	StoreId uint `json:"storeId"` // 仓库ID
	Id      uint `json:"id"`      // ID
}
