package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
		name string
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
			name: "case_1",
		},
		{
			nums: []int{0, 0},
			want: []int{0, 0},
			name: "case_2",
		},
		{
			nums: []int{-1, 1, 0, 3, -3},
			want: []int{0, 0, 9, 0, 0},
			name: "case_3",
		},
		{
			nums: []int{1, 0, 5, 0, 9},
			want: []int{0, 0, 0, 0, 0},
			name: "case_4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ProductExceptSelf(test.nums)

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}
