package utils

import "strconv"

func ConvertStringToInt8(str string, defaultResult int) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return defaultResult
	}
	return result
}

func ConvertStringToInt64(str string, defaultResult int64) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultResult
	}
	return result
}
