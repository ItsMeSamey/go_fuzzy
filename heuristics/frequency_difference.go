package heuristics

import "fuzzy/common"

func FrequencyDifferenceTrim[A common.StringLike, B common.StringLike](a A, b B) int {
  // Ensure b is shortest, for trimming
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

  return FrequencyDifference(a, b)
}

func FrequencyDifference[A common.StringLike, B common.StringLike](a A, b B) int {
  if len(a) < len(b) { return FrequencyDifference(b, a) }

  if len(b) == 0 { return len(a) }

  var f [256]int

  for i := range len(a) { f[a[i]] += 1 }
  for i := range len(b) { f[b[i]] -= 1 }

  sum := 0
  for i := range f { sum += common.Abs(f[i]) }

  return sum
}

