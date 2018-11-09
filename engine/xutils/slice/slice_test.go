package xutils

import (
	"testing"
)

func TestHasString(t *testing.T) {
	type args struct {
		list []string
		find string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"HasString",
			args{
				list: []string{"%+v", "%v", "%s", "%d"},
				find: "%d",
			},
			true,
		},
		{
			"HasString",
			args{
				list: []string{"%+v", "%v", "%s", "%d"},
				find: "%+v",
			},
			true,
		},
		{
			"HasString",
			args{
				list: []string{"%+v", "%v", "%s", "%d"},
				find: "%dd",
			},
			false,
		},
		{
			"HasString",
			args{
				list: []string{"%+v", "false", "中国", "国"},
				find: "中国",
			},
			true,
		},
		{
			"HasString",
			args{
				list: []string{"%+v", "false", "中国", "国"},
				find: "国",
			},
			true,
		},
		{
			"HasString",
			args{
				list: []string{"%+v", "false", "中国", "国"},
				find: "为",
			},
			false,
		},
		{
			"HasString",
			args{
				list: []string{"%+v", "false", "中国", "国"},
				find: "kk",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasString(tt.args.list, tt.args.find); got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

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
			"MaxInt0",
			args{
				list: []int{},
			},
			0,
		},
		{
			"MaxInt1",
			args{
				list: []int{0},
			},
			0,
		},
		{
			"MaxInt2",
			args{
				list: []int{1, 3},
			},
			3,
		},
		{
			"MaxInt3",
			args{
				list: []int{4, 2},
			},
			4,
		},
		{
			"MaxInt4",
			args{
				list: []int{1, 3, 2, 7, 5},
			},
			7,
		},
		{
			"MaxInt5",
			args{
				list: []int{6, 3, 1, 3, 2},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.list); got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

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
			"MinInt0",
			args{
				list: []int{},
			},
			0,
		},
		{
			"MinInt1",
			args{
				list: []int{0},
			},
			0,
		},
		{
			"MinInt2",
			args{
				list: []int{1, 3},
			},
			1,
		},
		{
			"MinInt3",
			args{
				list: []int{4, 2},
			},
			2,
		},
		{
			"MinInt4",
			args{
				list: []int{1, 3, 2, 7, 5},
			},
			1,
		},
		{
			"MinInt5",
			args{
				list: []int{6, 3, 1, 3, 2},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.list); got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
