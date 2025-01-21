package utils

import "strconv"

func ConvertStringToInt64(input string) int64 {
	int64Val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return int64Val
}

func ConvertInt64ToString(input int64) string {
	str := strconv.FormatInt(input, 10)
	return str
}
