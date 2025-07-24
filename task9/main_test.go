package main

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		question string
		facts    []string
		want     int
		name     string
	}{
		{
			question: "How old is Kl?",
			facts:    []string{"Kl is 65 years old", "Paaassat is the same age as Kl", "Otbajvmtgmtxpprs is 12 years old"},
			want:     65,
			name:     "case 1",
		},
		{
			question: "How old is Wcydinmm?",
			facts:    []string{"Tvpg is 100 years old", "Bupylflqdmybvabg is 83 years old", "Wcydinmm is the same age as Tvpg"},
			want:     100,
			name:     "case 1",
		},
		{
			question: "How old is Tzmahtvfaxwyp?",
			facts:    []string{"Aigwydwhmotispzk is 18 years old", "Xeqv is the same age as Tzmahtvfaxwyp", "Tzmahtvfaxwyp is 68 years old"},
			want:     68,
			name:     "case 1",
		},
		{
			question: "How old is A?",
			facts:    []string{"B is 2 years younger than Caac", "A is 44 years older than B", "B is 55 years old"},
			want:     99,
			name:     "case 1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Run(test.question, test.facts)

			if got != test.want {
				t.Errorf("Expected:  %v, got: %v", test.want, got)
			}
		})
	}
}
