package algorithms

import "testing"

func TestLevenshteinOSADistance(t *testing.T) {
  tests := []struct {
    name     string
    a        string
    b        string
    expected int
  }{
    {"Empty strings", "", "", 0},
    {"One empty string", "kitten", "", 6},
    {"Another empty string", "", "sitting", 7},
    {"Identical strings", "kitten", "kitten", 0},
    {"Simple substitution", "kitten", "sitten", 1},
    {"Simple insertion", "kitten", "kittens", 1},
    {"Simple deletion", "kitten", "kitte", 1},
    {"Substitution at the beginning", "kitten", "vitten", 1},
    {"Substitution at the end", "kitten", "kitteo", 1},
    {"Two substitutions", "kitten", "sitting", 3},
    {"Longer strings", "intention", "execution", 5},
    {"Different lengths", "abcdef", "azced", 3},
    {"Simple transposition", "ca", "ac", 1},
    {"Another transposition", "abcd", "badc", 2},
    {"Transposition with other edits", "kitten", "sitting", 3}, // Example where OSA might differ from standard Levenshtein
    {"More complex transposition", "mart", "tram", 3}, // Expected OSA distance for this example is 3
    {"Consecutive transpositions (OSA restriction)", "abdc", "acbd", 2}, // One transposition each
    {"Consecutive transpositions (OSA restriction) - 2", "abcd", "cadb", 4},
  }

  // Validate the tests with nonoptimal implementation
  for _, tt := range tests {
    actual := LevenshteinOSADistanceNonoptimal(tt.a, tt.b)
    if actual != tt.expected {
      t.Errorf("Validation Failed For: %s\nOSADistance(%q, %q) = %d, expected %d", tt.name, tt.a, tt.b, actual, tt.expected)
    }
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      actual := LevenshteinOSADistance(tt.a, tt.b) // Assuming your OSA implementation is named OSADistance
      if actual != tt.expected {
        t.Errorf("OSADistance(%q, %q) = %d, expected %d", tt.a, tt.b, actual, tt.expected)
      }
    })
  }
}

func LevenshteinOSADistanceNonoptimal(a, b string) int {
  d := make([][]int, len(a)+1)
  for i := 0; i <= len(a); i++ { d[i] = make([]int, len(b)+1) }

  for i := 0; i <= len(a); i++ { d[i][0] = i }
  for j := 0; j <= len(b); j++ { d[0][j] = j }

  for i := 1; i <= len(a); i++ {
    for j := 1; j <= len(b); j++ {
      cost := 0
      if a[i-1] == b[j-1] {
        cost = 0
      } else {
        cost = 1
      }
      d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+cost)

      if i > 1 && j > 1 && a[i-1] == b[j-2] && a[i-2] == b[j-1] {
        d[i][j] = min(d[i][j], d[i-2][j-2]+1)
      }
    }
  }
  return d[len(a)][len(b)]
}

