package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		friendMaxCards []int
		maxCardValue   int
		want           []int
		name           string
	}{
		{
			friendMaxCards: []int{3, 3, 2, 6, 5},
			maxCardValue:   7,
			want:           []int{4, 5, 3, 7, 6},
			name:           "case 1",
		},
		{
			friendMaxCards: []int{2, 1, 2, 2},
			maxCardValue:   4,
			want:           nil,
			name:           "case 2",
		},
		{
			friendMaxCards: []int{18, 6, 14, 7, 4, 11, 11, 5, 17, 9},
			maxCardValue:   20,
			want:           []int{19, 7, 15, 8, 5, 12, 13, 6, 18, 10},
			name:           "case 3",
		},
		{
			friendMaxCards: []int{11, 12, 2, 14, 1, 9, 8, 11, 18, 12},
			maxCardValue:   19,
			want:           []int{12, 13, 3, 15, 2, 10, 9, 14, 19, 16},
			name:           "case 4",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Run(test.friendMaxCards, test.maxCardValue)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Expect: %v, got: %v", test.want, got)
			}
		})
	}
}
