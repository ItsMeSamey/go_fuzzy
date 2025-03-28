package heuristics

import "fuzzy/common"

// Calculates the Optimal String Alignment (OSA) distance between two strings using a space-optimized approach.
// Implementation adapted from https://wikipedia.org/wiki/Damerau-Levenshtein_distance
//
// Time Complexity: O(n*m)
// Space Complexity: O(3 * min(n,m))
func LevenshteinOSADistance[A common.StringLike, B common.StringLike](a A, b B) int {
  // Ensure b is shortest, so length of v0 and v1 are minimized
  if len(a) < len(b) { return LevenshteinOSADistance(b, a) }

  // Trim common prefix
  for len(b) > 0 && a[0] == b[0] {
    a = a[1:]
    b = b[1:]
  }

  // Trim common suffix
  for len(b) > 0 && a[len(a)-1] == b[len(b)-1] {
    a = a[:len(a)-1]
    b = b[:len(b)-1]
  }

  return simpleLevenshteinOSADistance(a, b)
}

// Calculates the Optimal String Alignment (OSA) distance between two strings using a space-optimized approach.
// Use `LevenshteinOSADistance` unless you have a specific reason to use this version.
// Implementation adapted from https://wikipedia.org/wiki/Damerau-Levenshtein_distance
//
// Time Complexity: O(n*m)
// Space Complexity: O(3 * min(n,m)) 
func LevenshteinOSADistanceNoTrim[A common.StringLike, B common.StringLike](a A, b B) int {
  // Ensure b is shortest, so length of v0 and v1 are minimized
  if len(a) < len(b) { return LevenshteinOSADistanceNoTrim(b, a) }

  return simpleLevenshteinOSADistance(a, b)
}

func simpleLevenshteinOSADistance[A common.StringLike, B common.StringLike](a A, b B) int {
  // No exchanges can take place if smallest string is shorter than 2 characters
  if len(b) < 2 { return simpleLevenshteinDistance(a, b) }

  // To ensure single allocation
  buf := make([]int, 3 * (len(b)+1))

  // create two work vectors of integer distances
  v0 := buf[0 : len(b)+1]
  v1 := buf[len(b)+1: 2*(len(b)+1)]
  v2 := buf[2*(len(b)+1): 3*(len(b)+1)]

  // Initialize v0 same as in LevenshteinDistance
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

    { // so we dont do j >= 1 in the inner loop
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

    // Rotate slices
    v0, v1, v2 = v1, v2, v0
  }

  // v1 is actually v2, (due to rotation of slices)
  return v1[len(b)]
}

