package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
		name string
		hasTie bool // indicates if there's a tie in frequencies
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
			name: "example_1",
			hasTie: false,
		},
		{
			nums: []int{1},
			k:    1,
			want: []int{1},
			name: "example_2",
			hasTie: false,
		},
		{
			nums: []int{1, 2, 3, 1, 2, 3, 1},
			k:    2,
			want: []int{1, 2, 3}, // 1 appears 3 times, 2 and 3 both appear 2 times
			name: "tie_frequency",
			hasTie: true,
		},
		{
			nums: []int{4, 1, 4, 2, 3, 4, 1, 2, 4, 3, 4, 1, 4, 2, 3, 4, 1, 2, 4, 3},
			k:    3,
			want: []int{4, 1, 2, 3}, // 4 appears 8 times, 1, 2, 3 each appear 4 times
			name: "larger_input",
			hasTie: true,
		},
		{
			nums: []int{-1, -1, 2, 2, 3},
			k:    2,
			want: []int{-1, 2},
			name: "negative_numbers",
			hasTie: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TopKFrequent(test.nums, test.k)

			if test.hasTie {
				// For cases with ties, just check that all elements in 'got' are in 'want'
				for _, num := range got {
					found := false
					for _, wantNum := range test.want {
						if num == wantNum {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("TopKFrequent(%v, %d) = %v contains %d which is not in expected values %v", 
							test.nums, test.k, got, num, test.want)
					}
				}

				// Also check that we got the right number of elements
				if len(got) != test.k {
					t.Errorf("TopKFrequent(%v, %d) returned %d elements, want %d", 
						test.nums, test.k, len(got), test.k)
				}
			} else {
				// For cases without ties, we can do an exact comparison
				sort.Ints(got)
				sort.Ints(test.want)

				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("TopKFrequent(%v, %d) = %v, want %v", test.nums, test.k, got, test.want)
				}
			}
		})
	}
}
