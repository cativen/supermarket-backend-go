package utils

import (
	"gorm.io/gorm"
)

func Paginate[T any](db *gorm.DB, currentPage int, pageSize int, models []T) (count int64, records []T) {
	// 分页参数
	offset := (currentPage - 1) * pageSize
	//总条数
	db.Count(&count)
	// 查询销售记录
	db.Limit(pageSize).Offset(offset).Find(&models)
	return count, models
}
