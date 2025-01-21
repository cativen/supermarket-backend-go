package vo

type StorageSituationVo struct {
	StoreId    uint   `json:"storeId"`    // 仓库ID
	StoreName  string `json:"storeName"`  // 仓库名称
	ResidueNum int64  `json:"residueNum"` // 库存数量
}
