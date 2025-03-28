package heuristics

import (
  "fuzzy/common"
  "sort"
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
  if len(b) == 0 { return F(len(a)) }
  if len(a) == 1 {
    if a[0] == b[0] { return 0 }
    return F(1)
  }

  var fa [256][]uint32
  for i := range len(a) { fa[a[i]] = append(fa[a[i]], uint32(i)) }

  var fb [256][]uint32
  for i := range len(b) { fb[b[i]] = append(fb[b[i]], uint32(i)) }

  distance := F(0)
  for i := range 256 {
    ia := fa[i]
    ib := fb[i]

    if len(ia) == len(ib) {
      for j := range len(ia) {
        distance += F(common.Abs(int(ia[j]) - int(ib[j]))) / F(len(a) - 1)
      }
      continue
    }

    if len(ia) < len(ib) { ia, ib = ib, ia }
    if len(ib) == 0 {
      distance += F(len(ia))
      continue
    } else if len(ib) == 1 {
      distance += F(len(ia) - 1)
      idx := sort.Search(len(ia), func(i int) bool { return ia[i] >= ib[0] })

      if idx == len(ia) || idx == len(ia) - 1 {
        distance += F(common.Abs(int(ia[len(ia) - 1]) - int(ib[0]))) / F(len(a) - 1)
      } else {
        distance += F(min(
          common.Abs(int(ia[idx]) - int(ib[0])),
          common.Abs(int(ia[idx+1]) - int(ib[0])),
        )) / F(len(a) - 1)
      }
    } else {
      distance += F(len(ia) - len(ib))
      start := sort.Search(len(ia), func(i int) bool { return ia[i] >= ib[0] })
      end := sort.Search(len(ia), func(i int) bool { return ia[i] >= ib[len(ib)-1] })

      c_start := 0
      c_end := 0
      if end >= len(ia) - 1 {
        end = len(ia) - 1
        start = end - len(ib)
      } else if start == 0 {
        end = len(ib)
      } else {
        c_start = min(common.Abs(int(ia[start]) - int(ib[0])), common.Abs(int(ia[start + 1]) - int(ib[0])))
        c_end = min(common.Abs(int(ia[end]) - int(ib[len(ib)-1])), common.Abs(int(ia[end + 1]) - int(ib[len(ib)-1])))
      }

      for end - start < len(ib) {
        if c_start > c_end {
          end += 1
          if end == len(ia) - 1 {
            start = end - len(ib)
            break
          }
          c_end = min(common.Abs(int(ia[end]) - int(ib[len(ib)-1])), common.Abs(int(ia[end + 1]) - int(ib[len(ib)-1])))
        } else {
          start -= 1
          if start == 0 {
            end = len(ib)
            break
          }
          c_start = min(common.Abs(int(ia[start]) - int(ib[0])), common.Abs(int(ia[start + 1]) - int(ib[0])))
        }
      }

      c_start = common.Abs(int(ia[start]) - int(ib[0]))
      c_end = common.Abs(int(ia[end]) - int(ib[len(ib)-1]))
      partial_distance := F(0)
      for j := 1; j < len(ib) - 1; j += 1 {
        partial_distance += F(min(common.Abs(int(ia[start + j]) - int(ib[j]) - c_start), common.Abs(int(ia[start + j]) - int(ib[j]) - c_end)))
      }
      partial_distance += F(c_start + c_end)
      distance += partial_distance / F(len(a) - 1)
    }
  }

  return distance / F(len(a) + len(b))
}

