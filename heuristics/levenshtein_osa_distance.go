package heuristics

import "fuzzy/common"

// Calculates the Optimal String Alignment (OSA) distance between two strings using a space-optimized approach.
// Implementation adapted from https://wikipedia.org/wiki/Damerau-Levenshtein_distance
func LevenshteinOSADistance[A common.StringLike, B common.StringLike](a A, b B) int {
  // So the size of v0 and v1 are minimized
  if len(a) < len(b) { return LevenshteinOSADistance(b, a) }

  if len(b) < 2 { return LevenshteinDistance(a, b) }

  // To ensure single allocation
  buf := make([]int, 3 * (len(b)+1))
  // create two work vectors of integer distances
  v0 := buf[0 : len(b)+1]
  v1 := buf[len(b)+1: 2*(len(b)+1)]
  v2 := buf[2*(len(b)+1): 3*(len(b)+1)]

  for i := range len(b)+1 { v0[i] = i }

  // Fill v1 with first pass of `LevenshteinDistance`
  v1[0] = 1
  for j := range len(b) {
    increment := 0
    if a[0] != b[j] { increment = 1 }

    v1[j+1] = min(
      v0[j+1] + 1, // deletion cost
      v1[j] + 1, // insertion cost
      v0[j] + increment, // substitution cost
    )
  }

  for i := 1; i < len(a); i += 1 {
    v2[0] = i + 1

    {
      increment := 0
      if a[i] != b[0] { increment = 1 }

      v2[0+1] = min(
        v1[1] + 1, // deletion cost
        v2[0] + 1, // insertion cost
        v1[0] + increment, // substitution cost
      )
    }

    for j := 1; j < len(b); j += 1 {
      increment := 0
      if a[i] != b[j] { increment = 1 }

      v2[j+1] = min(
        v1[j+1] + 1, // deletion cost
        v2[j] + 1, // insertion cost
        v1[j] + increment, // substitution cost
      )
      if a[i] == b[j-1] && a[i-1] == b[j] {
        v2[j+1] = min(v2[j+1], v0[j-1] + 1) // transposition
      }
    }

    v0, v1, v2 = v1, v2, v0
  }

  return v1[len(b)]
}

