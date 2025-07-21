package main

import "testing"

func TestIsValidPass(t *testing.T) {
	tests := []struct {
		symbols map[string]int
		count   int
		pass    string
		want    string
		name    string
	}{
		{
			symbols: map[string]int{"a": 2, "l": 1, "w": 1, "k": 1, "o": 1, "f": 1, "p": 1, "z": 1}, //l w k o a f p z a
			count:   9,
			pass:    "kpzaalfow",
			want:    "YES",
			name:    "case 1",
		},
		{
			symbols: map[string]int{"a": 2, "l": 1, "w": 1, "k": 1, "o": 1, "f": 1, "p": 1, "z": 1}, //l w k o a f p z a
			count:   9,
			pass:    "awowfaaaw",
			want:    "NO",
			name:    "case 2",
		},
		{
			symbols: map[string]int{"a": 2, "l": 1, "w": 1, "k": 1, "o": 1, "f": 1, "p": 1, "z": 1}, //l w k o a f p z a
			count:   9,
			pass:    "opfzwaakl",
			want:    "YES",
			name:    "case 3",
		},
		{
			symbols: map[string]int{"a": 2, "b": 1, "c": 1},
			count:   4,
			pass:    "aabc",
			want:    "YES",
			name:    "case 4",
		},
		{
			symbols: map[string]int{"a": 2, "b": 1, "c": 1},
			count:   4,
			pass:    "bcca",
			want:    "NO",
			name:    "case 5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsValidPass(test.symbols, test.pass, test.count)

			if got != test.want {
				t.Errorf("Expected:  %v, got: %v", test.want, got)
			}
		})
	}

}
