package algorithms

import "testing"

func TestJaroDistance(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected float64
	}{
		{"Empty strings", "", "", 1.0},
		{"One empty string", "kitten", "", 0.0},
		{"Another empty string", "", "sitting", 0.0},
		{"Identical strings", "kitten", "kitten", 1.0},
		{"Simple substitution", "kitten", "sitten", 0.888888888888889},
		{"Transposition", "MARTHA", "MARHTA", 0.9444444444444445},
		{"Different lengths", "CRATE", "TRACE", 0.7333333333333333},
		{"No match", "foo", "bar", 0.0},
		{"Partial match 1", "aaa", "aab", 0.777777777777778},
		{"Partial match 2", "very", "vary", 0.8333333333333333},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := JaroDistance[float64](tt.a, tt.b)
			if !floatEquals(actual, tt.expected, 0.0000000000001) {
				t.Errorf("JaroDistance(%q, %q) = %f, expected %f", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}

func TestJaroWinklerLikeDistance(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		l        float64
		expected float64
	}{
		{"Empty strings", "", "", 0.1, 1.0},
		{"One empty string", "kitten", "", 0.1, 0.0},
		{"Another empty string", "", "sitting", 0.1, 0.0},
		{"Identical strings", "kitten", "kitten", 0.1, 1.0},
		{"No match", "foo", "bar", 0.1, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := JaroWinklerDistance(tt.a, tt.b, tt.l, -1)
			if !floatEquals(actual, tt.expected, 0.0000000000001) {
				t.Errorf("JaroWinklerLikeDistance(%q, %q, %f) = %f, expected %f", tt.a, tt.b, tt.l, actual, tt.expected)
			}
		})
	}
}

// floatEquals compares two float64 values for approximate equality within a given epsilon.
func floatEquals(a, b, epsilon float64) bool {
	return (a-b) < epsilon && (b-a) < epsilon
}

