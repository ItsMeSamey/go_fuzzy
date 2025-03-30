package algorithms

import "github.com/ItsMeSamey/go_fuzzy/common"

// Calculates the Morisitas Overlap Coefficient for the given strings using MultiSet.
// This May? not follow triangle inequality
// Implementation from https://en.wikipedia.org/wiki/Morisita%27s_overlap_index
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
func MorisitasOverlapCoefficient[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  var fa [256]int
  for i := range len(a) { fa[a[i]]++ }
  var fb [256]int
  for i := range len(b) { fb[b[i]]++ }

  d_a := F(0)
  for i := range 256 { d_a += F(fa[i]) * F(fa[i] - 1) }
  d_a /= F(len(a)) * F(len(a) - 1)

  d_b := F(0)
  for i := range 256 { d_b += F(fb[i]) * F(fb[i] - 1) }
  d_b /= F(len(b)) * F(len(b) - 1)

  numerator := F(0)
  for i := range 256 { numerator += F(fa[i]) * F(fb[i]) }
  numerator /= F(len(a)) * F(len(b))

  return numerator / (d_b + d_b)
}

// Calculates the Horns modification of the Morisitas Overlap Coefficient for the given strings using MultiSet.
// This May? not follow triangle inequality
// Implementation from https://en.wikipedia.org/wiki/Morisita%27s_overlap_index
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) = 256 * sizeof(int)
func HornsMorisitasOverlapCoefficient[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  var fa [256]int
  for i := range len(a) { fa[a[i]]++ }
  var fb [256]int
  for i := range len(b) { fb[b[i]]++ }

  d_a := F(0)
  for i := range 256 { d_a += F(fa[i]) * F(fa[i]) }
  d_a /= F(len(a)) * F(len(a))

  d_b := F(0)
  for i := range 256 { d_b += F(fb[i]) * F(fb[i]) }
  d_b /= F(len(b)) * F(len(b))

  numerator := F(0)
  for i := range 256 { numerator += F(fa[i]) * F(fb[i]) }
  numerator /= F(len(a)) * F(len(b))

  return numerator / (d_a + d_b)
}


