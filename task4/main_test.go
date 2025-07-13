package main

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		queue      []int
		queueCount int
		want       string
		name       string
	}{
		{
			queue:      []int{4, 7, 2, 1, 5, 4, 3, 5},
			queueCount: 8,
			want:       "x",
			name:       "case 1",
		},
		{
			queue:      []int{6, 5, 5, 1, 6, 7, 7, 4},
			queueCount: 8,
			want:       "x",
			name:       "case 2",
		},
		{
			queue:      []int{5, 2, 2, 3, 4},
			queueCount: 5,
			want:       "0-000",
			name:       "case 3",
		},
		{
			queue:      []int{8, 3, 2, 4, 3, 1},
			queueCount: 8,
			want:       "-00++0",
			name:       "case 4",
		},
		{
			queue:      []int{1, 4, 5, 1, 4},
			queueCount: 5,
			want:       "0-0+0",
			name:       "case 5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Run(test.queue, test.queueCount)
			if test.want != got {
				t.Errorf("Expected: %v, got: %v", test.want, got)
			}
		})
	}
}
