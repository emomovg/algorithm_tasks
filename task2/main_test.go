package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		numArrays []int
		want      []int
		name      string
	}{
		{
			numArrays: []int{20, 10, 20, 30},
			want:      []int{10, 20, 20, 30},
			name:      "case 1",
		},
		{
			numArrays: []int{6, 3, 4, 3, 1},
			want:      []int{1, 3, 3, 4, 6},
			name:      "case 2",
		},
		{
			numArrays: []int{200, 10, 100, 11, 200},
			want:      []int{10, 11, 100, 200, 200},
			name:      "case 3",
		},
		{
			numArrays: []int{13, 8, 12, 1, 7, 10, 1, 8, 10, 2, 17},
			want:      []int{1, 1, 2, 7, 8, 8, 10, 10, 12, 13, 17},
			name:      "case 4",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			got := QuickSort(test.numArrays)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Expected: %v\n  Got:      %v",
					test.want, got)
			}
		})
	}
}

func TestGetMappedPositions(t *testing.T) {
	var tests = []struct {
		sortArray     []int
		originalArray []int
		want          []int
		name          string
	}{
		{
			sortArray:     []int{10, 20, 20, 30},
			originalArray: []int{20, 10, 20, 30},
			want:          []int{2, 1, 2, 4},
			name:          "case 1",
		},
		{
			sortArray:     []int{1, 3, 3, 4, 6},
			originalArray: []int{6, 3, 4, 3, 1},
			want:          []int{5, 2, 2, 2, 1},
			name:          "case 2",
		},
		{
			sortArray:     []int{10, 11, 100, 200, 200},
			originalArray: []int{200, 10, 100, 11, 200},
			want:          []int{4, 1, 3, 1, 4},
			name:          "case 3",
		},
		{
			sortArray:     []int{1, 1, 2, 7, 8, 8, 10, 10, 12, 13, 17},
			originalArray: []int{13, 8, 12, 1, 7, 10, 1, 8, 10, 2, 17},
			want:          []int{9, 4, 9, 1, 4, 7, 1, 4, 7, 1, 11},
			name:          "case 4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := GetMappedPositions(test.sortArray, test.originalArray)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Expected: %v\n  Got:      %v", test.want, got)
			}
		})
	}
}
