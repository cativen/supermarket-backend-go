package utils

import "time"

// 时间转换字符串格式
func FormatTimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05") // Example format: 2006-01-02 15:04:05
}

// 字符串格式解析为时间格式
func ParseStringToTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}
