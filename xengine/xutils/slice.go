package xutils

import "sort"

// SliceMaxInt 求slice的最大int
func SliceMaxInt(list []int) int {
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

// SliceMinInt 求slice的最小int
func SliceMinInt(list []int) int {
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
