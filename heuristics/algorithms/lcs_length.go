package algorithms

import "fuzzy/common"

// Calculates the Length of Longest_Common_Subsequence between two strings using a space-optimized approach.
// Implementation adapted from https://wikipedia.org/wiki/Longest_common_subsequence
//
// Time Complexity: O(n*m)
// Space Complexity: O(2 * min(n,m))
func LCSLength[A common.StringLike, B common.StringLike](a A, b B) int {
  // We ensure that b is shorter, minimizing size of v0 and v1
  if len(a) < len(b) { return LCSLength(b, a) }

  if len(b) == 0 { return 0 }

  // To ensure that only one allocation is made
  buf := make([]int, 2 * (len(b)+1))

  // create two work vectors of integer distances
  v0 := buf[0 : len(b)+1]
  v1 := buf[len(b)+1: 2*(len(b)+1)]

  // Initialization is not needed as v0 is [0, ...] from initialization

  // Main loop
  for i := range len(a) {
    // v1[0] is already 0 from initialization
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

