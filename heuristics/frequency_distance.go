package heuristics

import (
  "fuzzy/common"
)

func FrequencyDistanceTrim[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // Ensure b is shortest, for trimming
  if len(a) < len(b) { return FrequencyDistanceTrim[F](b, a) }

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

  return FrequencyDistance[F](a, b)
}

func FrequencyDistance[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  if len(a) < len(b) { return FrequencyDistance[F](b, a) }

  var fa [256][]uint32
  for i := range len(a) { fa[a[i]] = append(fa[a[i]], uint32(i)) }

  var fb [256][]uint32
  for i := range len(b) { fb[b[i]] = append(fb[b[i]], uint32(i)) }

  distance := F(0)

  return distance
}

