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
  swapper SwapperInterface[A]
  scores  []F
}
func (s *sortAnyType[F, A]) Len() int { return s.swapper.Len() }
func (s *sortAnyType[F, A]) Less(i, j int) bool { return s.scores[i] < s.scores[j] }
func (s *sortAnyType[F, A]) Swap(i, j int) {
  s.swapper.Swap(i, j)
  s.scores[i], s.scores[j] = s.scores[j], s.scores[i]
}

// Sorts the `swapper` in place.
func (sorter Sorter[F, A, B]) SortAny(swapper SwapperInterface[A], target B) int {
  return sorter.sort(&sortAnyType[F, A]{
    swapper: swapper,
    scores:  sorter.ScoreAny(swapper, target, ),
  })
}

type sortType[F common.FloatType, A common.StringLike] struct {
  array    []A
  scores   []F
}
func (s *sortType[F, A]) Len() int { return len(s.array) }
func (s *sortType[F, A]) Less(i, j int) bool { return s.scores[i] < s.scores[j] }
func (s *sortType[F, A]) Swap(i, j int) {
  s.array[i], s.array[j] = s.array[j], s.array[i]
  s.scores[i], s.scores[j] = s.scores[j], s.scores[i]
}

// Sorts the array in place, returns the number of elements that are in the output
func (sorter Sorter[F, A, B]) Sort(array []A, target B) int {
  return sorter.sort(&sortType[F, A]{
    array:  array,
    scores: sorter.Score(array, target, ),
  })
}

func (sorter Sorter[F, A, B]) sort(data sort.Interface) int {
  return 0
}

