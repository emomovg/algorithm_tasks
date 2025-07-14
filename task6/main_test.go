package main

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		letters []string
		word    string
		pin     string
		want    string
		name    string
	}{
		{
			letters: []string{"l", "w", "k", "o", "a", "f", "p", "z", "a"},
			pin:     "kpzaalfow",
			want:    "YES",
			name:    "case 1",
		},
		{
			letters: []string{"a", "a", "b", "c"},
			pin:     "abc",
			want:    "NO",
			name:    "case 2",
		},
		{
			letters: []string{"h", "w", "w", "g", "l", "i", "c", "s"},
			pin:     "ghsgwclg",
			want:    "NO",
			name:    "case 3",
		},
		{
			letters: []string{"h", "w", "w", "g", "l", "i", "c", "s"},
			pin:     "wligscwh",
			want:    "YES",
			name:    "case 4",
		},
		{
			letters: []string{"h", "w", "w", "g", "l", "i", "c", "s"},
			pin:     "slglglig",
			want:    "NO",
			name:    "case 5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Run(test.letters, test.pin)

			if got != test.want {
				t.Errorf("Expected: %v, got: %v", test.want, got)
			}
		})
	}

}
