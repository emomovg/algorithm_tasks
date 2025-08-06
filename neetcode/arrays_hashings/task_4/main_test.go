package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
		name string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}},
			name: "case_1",
		},
		{
			strs: []string{""},
			want: [][]string{{""}},
			name: "case_2",
		},
		{
			strs: []string{"a"},
			want: [][]string{{"a"}},
			name: "case_3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := GroupAnagrams(test.strs)

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}
