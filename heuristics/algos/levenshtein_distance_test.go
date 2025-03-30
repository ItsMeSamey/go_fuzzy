package heuristics

import "testing"

func TestLevenshteinDistance(t *testing.T) {
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
    {"Two substitutions and addition to end", "kitten", "sitting", 3},
    {"Longer strings", "intention", "execution", 5},
    {"Different lengths", "abcdef", "azced", 3},
    {"Transposition (should be 2 with simple implementation)", "ca", "ac", 2},
    {"Another transposition", "abcd", "badc", 3}, // abcd (initial) -> bcd (1) -> bad (2) -> badc (3)
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      actual := LevenshteinDistance(tt.a, tt.b)
      if actual != tt.expected {
        t.Errorf("LevenshteinDistance(%q, %q) = %d, expected %d", tt.a, tt.b, actual, tt.expected)
      }
    })
  }
}

