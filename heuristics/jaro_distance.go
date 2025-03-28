package heuristics

import "fuzzy/common"

// JaroDistance calculates the Jaro distance between two strings.
//
// The Jaro distance is a measure of similarity between two strings. It is
// defined as:
//
// jaro_distance = 1/3 * (m/|s1| + m/|s2| + (m - t)/m)
func JaroDistance[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  if len(a) < len(b) { return JaroDistance[F](b, a) }

  if len(b) == 0 {
    if len(a) == 0 { return 1 }
    return 0
  }

  matchDistance := max(len(a), len(b))/2 - 1
  matches := 0
  aMatches := make([]bool, len(a))
  bMatches := make([]bool, len(b))

  // Find the number of matching characters.
  for i := range len(a) {
    start := max(0, i-matchDistance)
    end := min(len(b)-1, i+matchDistance)

    for j := start; j <= end; j++ {
      if a[i] == b[j] && !bMatches[j] {
        aMatches[i] = true
        bMatches[j] = true
        matches++
        break
      }
    }
  }

  if matches == 0 { return 0 }

  // Calculate the number of transpositions.
  transpositions := 0
  k := 0
  for i := range len(a) {
    if aMatches[i] {
      for !bMatches[k] { k++ }
      if a[i] != b[k] { transpositions++ }
      k++
    }
  }
  transpositions /= 2

  // Calculate the Jaro distance.
  return (F(matches)/F(len(a)) + F(matches)/F(len(b)) + (F(matches) - F(transpositions))/F(matches)) / 3
}

// JaroWinklerDistance calculates the Jaro-Winkler distance between two strings by
// giving more favorable ratings to strings that match from the beginning and the end.
func JaroWinklerLikeDistance[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B, l F) F {
  jaro := JaroDistance[F](a, b)

  if jaro < 0.7 { return jaro }

  // Calculate the length of the matching prefix (up to max 4 characters).
  prefix := 0
  for len(a) > 0 && len(b) > 0 && a[0] == b[0] {
    prefix++
    a = a[1:]
    b = b[1:]
  }

  for len(a) > 0 && len(b) > 0 && a[len(a)-1] == b[len(b)-1] {
    prefix++
    a = a[:len(a)-1]
    b = b[:len(b)-1]
  }

  // Calculate the Jaro-Winkler distance.
  return jaro + F(prefix) * l * (1-jaro)
}

