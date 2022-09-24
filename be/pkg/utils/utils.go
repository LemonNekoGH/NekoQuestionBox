package utils

import "reflect"

// IndexOf. 获取数组或切片中要查找的元素首次出现的位置
func IndexOf[T any](array []T, search T) int {
	for index, element := range array {
		if reflect.DeepEqual(element, search) {
			return index
		}
	}
	return -1
}

// IsArrayContains. 查询数组或切片中是否存在某一元素
func IsArrayContains[T any](array []T, search T) bool {
	return IndexOf(array, search) != -1
}
