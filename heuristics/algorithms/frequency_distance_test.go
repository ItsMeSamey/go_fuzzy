package algorithms

import "testing"

func TestFrequencyDistanceNoCrash(t *testing.T) {
  testCases := []struct {
    a string
    b string
  }{
    {"", ""},
    {"a", "a"},
    {"abb", "bba"},
    {"abc", "acb"},
    {"abc", "bac"},
    {"aabb", "abab"},
    {"aaaa", "bbbb"},
    {"abc", "abcd"},
    {"abcd", "abc"},
    {"apple", "apxle"},
    {"apple", "apxpl"},
    {"apple", "axple"},
    {"apple", "bpple"},
    {"hello", "world"},
    {"testing", "test"},
    {"test", "testing"},
    {"aaaaa", "aaaba"},
    {"aaaba", "aaaaa"},
    {"aaaaa", "aabba"},
    {"aabba", "aaaaa"},
    {"abcde", "edcba"},
    {"microsoft", "mitsubishi"},
    {"intention", "execution"},
    {"aaaa", "aaa"},
    {"aaa", "aaaa"},
    {"cat", "act"},
    {"dog", "god"},
    {"listen", "silent"},
  }

  for _, tc := range testCases {
    t.Run("", func(t *testing.T) {
      _ = FrequencyDistance[float64](tc.a, tc.b)
    })
  }
}

