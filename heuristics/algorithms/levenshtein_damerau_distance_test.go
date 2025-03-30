package algorithms

import "testing"

func TestDamerauLevenshteinDistance(t *testing.T) {
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
    {"Transposition", "ca", "ac", 1},
    {"Longer Transposition", "cxa", "axc", 2},
    {"Another transposition", "abcd", "badc", 2},
    {"Transposition at start", "abcde", "bacde", 1},
    {"Two transpositions", "abdcfe", "adbcef", 2},
    {"Transposition and deletion", "abcd", "ac", 2},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      actual := DamerauLevenshteinDistance(tt.a, tt.b)
      if actual != tt.expected {
        t.Errorf("DamerauLevenshteinDistance(%q, %q) = %d, expected %d", tt.a, tt.b, actual, tt.expected)
      }
    })
  }
}

