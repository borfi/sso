package xutils

import "sort"

// MaxInt 求slice的最大int
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

// MinInt 求slice的最小int
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
