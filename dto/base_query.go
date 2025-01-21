package dto

type BaseQuery struct {
	CurrentPage int                    `json:"currentPage"` // 页码
	PageSize    int                    `json:"pageSize"`    // 每页大小
	Params      map[string]interface{} `json:"params"`      // 参数
}

func (b *BaseQuery) SetDefaultPageSize() {
	if b.PageSize == 0 {
		b.PageSize = 5
	}
}
