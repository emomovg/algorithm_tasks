package main

import "testing"

func TestInAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
		name string
	}{
		{
			s:    "anagram",
			t:    "anagram",
			want: true,
			name: "case_1",
		},
		{
			s:    "rat",
			t:    "car",
			want: false,
			name: "case_2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsAnagram(test.s, test.t)

			if got != test.want {
				t.Errorf("Expected %v, got: %v", test.want, got)
			}
		})
	}
}
