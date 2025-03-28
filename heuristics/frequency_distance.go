package heuristics

import (
	"fuzzy/common"

	"github.com/blizzy78/varnamelen"
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


// type indexTaker struct {
//   index uint32
//   taker uint32
// }
//
// type indexValue struct {
//   filledTill int
//   array []indexTaker
// }

func FrequencyDistance[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  if len(a) < len(b) { return FrequencyDistance[F](b, a) }

  var fa [256][]uint32
  for i := range len(a) { fa[a[i]] = append(fa[a[i]], uint32(i)) }

  var fb [256][]uint32
  for i := range len(b) { fb[b[i]] = append(fb[b[i]], uint32(i)) }

  distance := F(0)
  for i := range len(a) {
    if fb[a[i]].filledTill == len(fb[a[i]].array) {
      distance += 1
      continue
    }
    idx := fb[a[i]].filledTill
    fb[a[i]].filledTill += 1
    fb[a[i]][idx].taker = uint32(i)
    distance += F(common.Abs(int(fb[a[i]][0].index) - i)) / F(len(a) - 1)
  }

  return distance
}

