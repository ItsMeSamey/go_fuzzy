package heuristics

import "fuzzy/common"

// Calculates the Length of Longest_Common_Subsequence between two strings using a space-optimized approach.
// Implementation adapted from https://wikipedia.org/wiki/Longest_common_subsequence
func LCSLength[A common.StringLike, B common.StringLike](a A, b B) int {
  // So the size of v0 and v1 are minimized
  if len(a) < len(b) { return LCSLength(b, a) }

  if len(b) == 0 { return 0 }

  buf := make([]int, 2 * (len(b)+1))
  // create two work vectors of integer distances
  v0 := buf[0 : len(b)+1]
  v1 := buf[len(b)+1: 2*(len(b)+1)]

  for i := range len(b)+1 { v0[i] = 0 }

  for i := range len(a) {
    for j := range len(b) {
      if a[i] == b[j] {
        v1[j+1] = v0[j] + 1
      } else {
        v1[j+1] = max(v0[j+1], v1[j])
      }
    }

    v0, v1 = v1, v0
  }
  return v0[len(b)]
}

