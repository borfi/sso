package xutils

import (
	"testing"
)

// TestMaxInt .
func TestMaxInt(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"MaxInt",
			args{
				list: []int{},
			},
			0,
		},
		{
			"MaxInt",
			args{
				list: []int{0},
			},
			0,
		},
		{
			"MaxInt",
			args{
				list: []int{1, 3},
			},
			3,
		},
		{
			"MaxInt",
			args{
				list: []int{4, 2},
			},
			4,
		},
		{
			"MaxInt",
			args{
				list: []int{1, 3, 2, 7, 5},
			},
			7,
		},
		{
			"MaxInt",
			args{
				list: []int{6, 3, 1, 3, 2},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.list); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMinInt .
func TestMinInt(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"MinInt",
			args{
				list: []int{},
			},
			0,
		},
		{
			"MinInt",
			args{
				list: []int{0},
			},
			0,
		},
		{
			"MinInt",
			args{
				list: []int{1, 3},
			},
			1,
		},
		{
			"MinInt",
			args{
				list: []int{4, 2},
			},
			2,
		},
		{
			"MinInt",
			args{
				list: []int{1, 3, 2, 7, 5},
			},
			1,
		},
		{
			"MinInt",
			args{
				list: []int{6, 3, 1, 3, 2},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.list); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
