package xutils

import "sort"

// HasInt 判断int数是否存在slice之中
func HasInt(sl []int, s int) bool {
	for k := range sl {
		if sl[k] == s {
			return true
		}
	}
	return false
}

// HasString 判断字符串是否存在slice之中
func HasString(sl []string, s string) bool {
	for k := range sl {
		if sl[k] == s {
			return true
		}
	}
	return false
}

// MaxInt 求slice的最大int值
func MaxInt(list []int) int {
	l := len(list)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return list[0]
	}
	sort.Ints(list)
	return list[l-1]
}

// MinInt 求slice的最小int值
func MinInt(list []int) int {
	l := len(list)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return list[0]
	}
	sort.Ints(list)
	return list[0]
}
