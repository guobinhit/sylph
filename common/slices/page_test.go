package slices

import (
	"fmt"
	"github.com/guobinhit/sylph/common/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPage(t *testing.T) {
	type args struct {
		arr       interface{}
		pageNum   int
		pageLimit int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Page int slice",
			args: args{
				arr:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				pageNum:   1,
				pageLimit: 3,
			},
			want:  `[1,2,3]`,
			want1: true,
		},
		{
			name: "Page int32 slice",
			args: args{
				arr:       []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				pageNum:   1,
				pageLimit: 3,
			},
			want:  `[1,2,3]`,
			want1: true,
		},
		{
			name: "Page string slice 1",
			args: args{
				arr:       []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				pageNum:   4,
				pageLimit: 3,
			},
			want:  `["10"]`,
			want1: false,
		},
		{
			name: "Page string slice 2",
			args: args{
				arr:       []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				pageNum:   5,
				pageLimit: 3,
			},
			want:  `null`,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Page(tt.args.arr, tt.args.pageNum, tt.args.pageLimit)
			assert.Equalf(t, tt.want, utils.Json(got), "Page(%v, %v, %v)", tt.args.arr, tt.args.pageNum, tt.args.pageLimit)
			assert.Equalf(t, tt.want1, got1, "Page(%v, %v, %v)", tt.args.arr, tt.args.pageNum, tt.args.pageLimit)
		})
	}
}

func TestPage_Usage(t *testing.T) {
	type args struct {
		arr       interface{}
		pageNum   int
		pageLimit int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Page Usage",
			args: args{
				arr:       []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				pageNum:   3,
				pageLimit: 3,
			},
			want:  `["7","8","9"]`,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Unit Test
			got, got1 := Page(tt.args.arr, tt.args.pageNum, tt.args.pageLimit)
			assert.Equalf(t, tt.want, utils.Json(got), "Page(%v, %v, %v)", tt.args.arr, tt.args.pageNum, tt.args.pageLimit)
			assert.Equalf(t, tt.want1, got1, "Page(%v, %v, %v)", tt.args.arr, tt.args.pageNum, tt.args.pageLimit)

			// Usage of Page
			targetArr := got.([]string)
			if targetArr == nil || len(targetArr) == 0 {
				fmt.Println(fmt.Sprintf("got is nil: Page(%v, %v, %v)", tt.args.arr, tt.args.pageNum, tt.args.pageLimit))
			} else {
				for _, a := range targetArr {
					fmt.Println("targetArr element of: ", a)
				}
			}
		})
	}
}
