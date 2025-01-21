package vo

type Page[T any] struct {
	Records     []T   `json:"records"` // 存储分页结果的切片
	Total       int64 `json:"total"`
	Size        int   `json:"size"`
	Current     int   `json:"current"`
	OptimizeSql bool  `json:"optimizeCountSql"`
	SearchCount bool  `json:"searchCount"`
}
