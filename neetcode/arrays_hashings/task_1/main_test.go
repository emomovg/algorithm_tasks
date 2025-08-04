package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		wont bool
		name string
	}{
		{
			nums: []int{1, 5, 0, 15},
			wont: false,
			name: "case_1",
		},
		{
			nums: []int{1, 1, 3, 10},
			wont: true,
			name: "case_2",
		},
		{
			nums: []int{1, 0, 4, 24, 443, -5, 1},
			wont: true,
			name: "case_3",
		},
		{
			nums: []int{1, -1, 0, 4},
			wont: false,
			name: "case_4",
		},
		{
			nums: []int{1, 0, 5, 3, 2, 0, 15},
			wont: true,
			name: "case_5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ContainsDuplicate(test.nums)

			if got != test.wont {
				t.Errorf("Expected %v, got %v\n", test.wont, got)
			}
		})
	}
}
