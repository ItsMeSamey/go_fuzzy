package heuristics

import "testing"

func TestLCSLength(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{"Empty strings", "", "", 0},
		{"One empty string (string)", "AGGTAB", "", 0},
		{"One empty string ([]rune)", "AGGTAB", "", 0},
		{"Identical strings (string)", "AGGTAB", "AGGTAB", 6},
		{"Identical strings ([]rune)", "AGGTAB", "AGGTAB", 6},
		{"No common subsequence", "ABC", "DEF", 0},
		{"Common subsequence at the beginning", "ABCDEF", "ABXYZ", 2},
		{"Common subsequence at the end", "XYZABC", "UVWABC", 3},
		{"Common subsequence in the middle", "AXBYCZ", "PBYQCR", 3},
		{"Different lengths (string)", "AGGTAB", "GXTXAYB", 4},
		{"Different lengths ([]rune)", "AGGTAB", "GXTXAYB", 4},
		{"Another different lengths (string)", "ABCDGH", "AEDFHR", 3},
		{"Another different lengths ([]rune)", "ABCDGH", "AEDFHR", 3},
		{"Common subsequence with interleaving characters (string)", "ABCDE", "ACE", 3},
		{"Common subsequence with interleaving characters ([]rune)", "ABCDE", "ACE", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := LCSLength(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("LCSLength(%v, %v) = %d, expected %d", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}

