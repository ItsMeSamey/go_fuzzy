package heuristics

import "fuzzy/common"

// Calculates the Damerau-Levenshtein distance between two strings.
// Implementation from https://wikipedia.org/wiki/Damerau-Levenshtein_distance
//
// Time Complexity: O(n*m)
// Space Complexity: O(n*m)
func DamerauLevenshteinDistance[A common.StringLike, B common.StringLike](a A, b B) int {
  // Ensure b is shortest, so length of v0 and v1 are minimized
  if len(a) < len(b) { return DamerauLevenshteinDistance(b, a) }

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

  return simpleDamerauLevenshteinDistance(a, b)
}

// Calculates the Damerau-Levenshtein distance between two strings.
// Use `DamerauLevenshteinDistance` unless you have a specific reason to use this version.
// Implementation from https://wikipedia.org/wiki/Damerau-Levenshtein_distance
//
// Time Complexity: O(n*m)
// Space Complexity: O(n*m)
func DamerauLevenshteinDistanceNoTrim[A common.StringLike, B common.StringLike](a A, b B) int {
  // Ensure b is shortest, so length of v0 and v1 are minimized
  if len(a) < len(b) { return DamerauLevenshteinDistanceNoTrim(b, a) }

  return simpleDamerauLevenshteinDistance(a, b)
}

func simpleDamerauLevenshteinDistance[A common.StringLike, B common.StringLike](a A, b B) int {
  var da [256]int // Ascii character set

  if (len(b) < 3) { return simpleLevenshteinOSADistance(a, b) }

  // the 2D matrix
  d := make([]int, (len(a)+2)*(len(b)+2))

  // maximum possible distance, used for initialization.
  maxdist := len(a) + len(b)

  // Initialize the distance matrix.
  d[0] = maxdist // d[-1, -1]
  for i := 1; i <= len(a)+1; i++ { // i from 0 to len(a)
    d[i*(len(b)+2)] = maxdist // d[i, -1]
    d[i*(len(b)+2)+1] = i - 1 // d[i, 0]
  }
  for j := 1; j <= len(b)+1; j++ { // j from 0 to len(b)
    d[j] = maxdist // d[-1, j]
    d[j+len(b)+2] = j - 1// d[0, j]
  }

  // Calculate the Damerau-Levenshtein distance.
  for i := 1; i <= len(a); i++ {
    db := 0
    for j := 1; j <= len(b); j++ {
      k := da[b[j-1]] // Go is 0-indexed, pseudocode 1-indexed
      l := db
      cost := 0
      if a[i-1] == b[j-1] { // Go is 0-indexed
        cost = 0
        db = j
      } else {
        cost = 1
      }
      substitutionCost := d[(i-1+1)*(len(b)+2)+(j-1+1)] + cost
      insertionCost := d[(i+1)*(len(b)+2)+(j-1+1)] + 1
      deletionCost := d[(i-1+1)*(len(b)+2)+(j+1)] + 1
      transpositionCost := 0
      if k > 0 && l > 0 {
        transpositionCost = d[(k-1+1)*(len(b)+2)+(l-1+1)] + (i - k - 1) + 1 + (j - l - 1)
      } else {
        transpositionCost = maxdist // Or any value > substitutionCost, insertionCost, deletionCost
      }

      d[(i+1)*(len(b)+2)+(j+1)] = min(substitutionCost, insertionCost, deletionCost, transpositionCost)
    }
    da[a[i-1]] = i
  }

  return d[(len(a)+1)*(len(b)+2)+(len(b)+1)] // d[len(a), len(b)]
}

