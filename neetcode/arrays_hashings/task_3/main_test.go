package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
		name   string
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
			name:   "case_1",
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
			name:   "case_2",
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
			name:   "case_3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TwoSum(test.nums, test.target)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}
