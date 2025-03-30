package gofuzzy

import (
  "sort"

  "fuzzy/common"

  "golang.org/x/text/transform"
)

type Scorer[F common.FloatType, A common.StringLike, B common.StringLike] struct {
  // Returns a score Between 0 and 1 for the given pair of A and B.
  ScoreFn func(a A, b B) F

  // Transformer
  Transformer transform.Transformer
}

type AccessorInterface[A common.StringLike] interface {
  // Len is the number of elements in the collection.
  Len() int
  // Get i'th element in the collection.
  Get(i int) A
}

// Give an array of scores for all the elements in the `accessor` w.r.t. the `target`.
func (sorter Scorer[F, A, B]) ScoreAny(accessor AccessorInterface[A], target B, ) (out []F) {
  out = make([]F, accessor.Len())
  for i := range accessor.Len() { out[i] = sorter.ScoreFn(accessor.Get(i), target) }
  return
}

// Give an array of scores for all the elements in the `array` w.r.t. the `target`.
func (sorter Scorer[F, A, B]) Score(array []A, target B) (out []F) {
  out = make([]F, len(array))
  for i, a := range array { out[i] = sorter.ScoreFn(a, target) }
  return
}

// A Struct used to sort a collection of elements.
type Sorter[F common.FloatType, A common.StringLike, B common.StringLike] struct {
  Scorer[F, A, B]

  // A value Between 0 and 1 that determines the threshold for the sort.
  Threshold F
}

type SwapperInterface[A common.StringLike] interface {
  AccessorInterface[A]
  Swap(i, j int)
}

type sortAnyType[F common.FloatType, A common.StringLike] struct {
  len     int
  swapper SwapperInterface[A]
  scores  []F
}
func (s *sortAnyType[F, A]) Len() int { return s.len }
func (s *sortAnyType[F, A]) SetLen(i int) { s.len = i }
func (s *sortAnyType[F, A]) Less(i, j int) bool { return s.scores[i] < s.scores[j] }
func (s *sortAnyType[F, A]) Swap(i, j int) {
  s.swapper.Swap(i, j)
  s.scores[i], s.scores[j] = s.scores[j], s.scores[i]
}
func (s *sortAnyType[F, A]) Score(i int) F { return s.scores[i] }

// Sorts the `swapper` in place.
func (sorter Sorter[F, A, B]) SortAny(swapper SwapperInterface[A], target B) int {
  return sorter.sort(&sortAnyType[F, A]{
    len:     swapper.Len(),
    swapper: swapper,
    scores:  sorter.ScoreAny(swapper, target, ),
  })
}

type sortType[F common.FloatType, A common.StringLike] struct {
  len    int
  array  []A
  scores []F
}
func (s *sortType[F, A]) Len() int { return s.len }
func (s *sortType[F, A]) SetLen(i int) { s.len = i }
func (s *sortType[F, A]) Less(i, j int) bool { return s.scores[i] < s.scores[j] }
func (s *sortType[F, A]) Swap(i, j int) {
  s.array[i], s.array[j] = s.array[j], s.array[i]
  s.scores[i], s.scores[j] = s.scores[j], s.scores[i]
}
func (s *sortType[F, A]) Score(i int) F { return s.scores[i] }

// Sorts the array in place, returns the number of elements that are in the output
func (sorter Sorter[F, A, B]) Sort(array []A, target B) int {
  return sorter.sort(&sortType[F, A]{
    len:    len(array),
    array:  array,
    scores: sorter.Score(array, target, ),
  })
}

type sortInterface[F common.FloatType, A common.StringLike] interface {
  Len() int
  SetLen(i int)
  Less(i, j int) bool
  Swap(i, j int)
  Score(i int) F
}
func (sorter Sorter[F, A, B]) sort(data sortInterface[F, A]) int {
  below := 0

  if sorter.Threshold == 0 {
    for below < data.Len() && data.Score(below) >= sorter.Threshold { below += 1 }
    for i := below; i < data.Len(); i += 1 {
      if data.Score(i) < sorter.Threshold { continue }
      data.Swap(i, below)
      below += 1
    }
    data.SetLen(below)
  } else {
    below = data.Len()
  }

  sort.Sort(data)
  return below
}

