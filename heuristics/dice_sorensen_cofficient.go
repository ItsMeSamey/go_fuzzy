package heuristics

import "fuzzy/common"

func IntersectionCount[A common.StringLike, B common.StringLike](a A, b B) uint {
  var f [256]int

  for i := range len(a) { f[a[i]] += 1 }

  intersection := 0
  for i := range len(b) {
    if f[b[i]] > 0 {
      intersection += 1
      f[b[i]] -= 1
    }
  }

  return uint(intersection)
}

func DiceSorensenCoefficient[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  return F(2 * IntersectionCount(a, b)) / F(uint(len(a)) + uint(len(b)))
}

