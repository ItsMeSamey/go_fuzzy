package algorithms

import "fuzzy/common"

// Calculates the Levenshtein(/edit) distance between two strings using a space-optimized approach.
// Implementation from https://wikipedia.org/wiki/Levenshtein_distance
//
// Time Complexity: O(n*m)
// Space Complexity: O(2 * min(n,m))
func LevenshteinDistance[A common.StringLike, B common.StringLike](a A, b B) int {
  // Ensure b is shortest, so length of v0 and v1 are minimized
  if len(a) < len(b) { return LevenshteinDistance(b, a) }

  if len(b) == 0 { return len(a) }

  // For ensuring single allocation
  buf := make([]int, 2 * (len(b)+1))

  // create two work vectors of integer distances
  v0 := buf[0: len(b)+1]
  v1 := buf[len(b)+1: 2*(len(b)+1)]

  // initialize v0 (the previous row of distances)
  // this row is A[0][i]: edit distance from an empty s to t;
  // that distance is the number of characters to append to s to make t.
  for i := range len(b)+1 { v0[i] = i }

  for i := range len(a) {
    // calculate v1 (current row distances) from the previous row v0

    // first element of v1 is A[i + 1][0]
    // edit distance is delete (i + 1) chars from s to match empty t
    v1[0] = i + 1

    // fill in the rest of the row
    for j := range len(b) {
      increment := 0
      if a[i] != b[j] { increment = 1 }

      v1[j+1] = min(
        v0[j+1] + 1, // deletion cost
        v1[j] + 1, // insertion cost
        v0[j] + increment, // substitution cost
      )
    }

    // copy v1 (current row) to v0 (previous row) for next iteration
    v0, v1 = v1, v0
  }

  // after the last swap, the results of v1 are now in v0
  return v0[len(b)]
}

