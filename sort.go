package gofuzzy

import (
  "fuzzy/common"
  "sort"
)


type AccessorInterface[A common.StringLike] interface {
  // Len is the number of elements in the collection.
  Len() int

  Get(i int) A
}

// Give an array of scores for all the elements in the `accessor` w.r.t. the `target`.
func ScoreAny[F common.FloatType, A common.StringLike, B common.StringLike](accessor AccessorInterface[A], target B, scoreFn func(a A, b B) F) (out []F) {
  out = make([]F, accessor.Len())
  for i := range accessor.Len() { out[i] = scoreFn(accessor.Get(i), target) }
  return
}

// Give an array of scores for all the elements in the `array` w.r.t. the `target`.
func Score[F common.FloatType, A common.StringLike, B common.StringLike](array []A, target B, scoreFn func(a A, b B) F) (out []F) {
  out = make([]F, len(array))
  for i, a := range array { out[i] = scoreFn(a, target) }
  return
}

type SwapperInterface[A common.StringLike] interface {
  AccessorInterface[A]
  Swap(i, j int)
}

type sortAnyType[F common.FloatType, A common.StringLike] struct {
  swapper SwapperInterface[A]
  scores  []F
}
func (s *sortAnyType[F, A]) Len() int {
  return s.swapper.Len()
}
func (s *sortAnyType[F, A]) Less(i, j int) bool {
  return s.scores[i] < s.scores[j]
}
func (s *sortAnyType[F, A]) Swap(i, j int) {
  s.swapper.Swap(i, j)
  s.scores[i], s.scores[j] = s.scores[j], s.scores[i]
}

// Sorts the `swapper` in place.
func SortAny[F common.FloatType, A common.StringLike, B common.StringLike](swapper SwapperInterface[A], target B, scoreFn func(a A, b B) F) {
  sorter := &sortAnyType[F, A]{
    swapper: swapper,
    scores:  ScoreAny(swapper, target, scoreFn),
  }
  sort.Sort(sorter)
}

type sortType[F common.FloatType, A common.StringLike] struct {
  array    []A
  scores   []F
}
func (s *sortType[F, A]) Len() int {
  return len(s.array)
}
func (s *sortType[F, A]) Less(i, j int) bool {
  return s.scores[i] < s.scores[j]
}
func (s *sortType[F, A]) Swap(i, j int) {
  s.array[i], s.array[j] = s.array[j], s.array[i]
  s.scores[i], s.scores[j] = s.scores[j], s.scores[i]
}

// The sort options with a 
type SortOptions[F common.FloatType, A common.StringLike, B common.StringLike] struct {
  // Returns a score Between 0 and 1 for the given pair of A and B.
  ScoreFn func(a A, b B) F

  // A value Between 0 and 1 that determines the treshold for the sort.
  Treshold F
}

// Sorts the array in place
func Sort[F common.FloatType, A common.StringLike, B common.StringLike](array []A, target B, options SortOptions[F, A, B]) {
  sorter := &sortType[F, A]{
    array:    array,
    scores:   Score(array, target, options.ScoreFn),
  }
  sort.Sort(sorter)
}

