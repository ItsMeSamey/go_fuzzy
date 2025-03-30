package fuzzy

import (
  "sort"

  "github.com/ItsMeSamey/go_fuzzy/common"
  "github.com/ItsMeSamey/go_fuzzy/heuristics"
  "github.com/ItsMeSamey/go_fuzzy/transformers"

  "golang.org/x/text/transform"
)

type AccessorInterface[A common.StringLike] interface {
  // Len is the number of elements in the collection.
  Len() int
  // Get i'th element in the collection.
  Get(i int) A
}

type SwapperInterface[A common.StringLike] interface {
  AccessorInterface[A]
  Swap(i, j int)
}

type SwapperArrayInterface[T any, A common.StringLike] interface {
  SwapperInterface[A]
  Array() []T
}

type sortableArray[A common.StringLike] []A
func (s sortableArray[A]) Len() int { return len(s) }
func (s sortableArray[A]) Get(i int) A { return s[i] }
func (s sortableArray[A]) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortableArray[A]) Array() []A { return s }

func ToSwapperArray[A common.StringLike](array []A) SwapperArrayInterface[A, A] {
  return sortableArray[A](array)
}

type sortableType[T any, A common.StringLike] struct {
  array []T
  get   func (t T) A
}
func (s sortableType[T, A]) Len() int { return len(s.array) }
func (s sortableType[T, A]) Get(i int) A { return s.get(s.array[i]) }
func (s sortableType[T, A]) Swap(i, j int) { s.array[i], s.array[j] = s.array[j], s.array[i] }
func (s sortableType[T, A]) Array() []T { return s.array }

func ToSwapper[T any, A common.StringLike](array []T, get func (t T) A) SwapperArrayInterface[T, A] {
  if get == nil { panic("get must be set") }
  return sortableType[T, A]{array, get}
}

func transformStringLike[A common.StringLike](t transform.Transformer, a A) []byte {
  bytea := []byte(a)
  if t == nil { return bytea }
  out, _, err := transform.Bytes(t, bytea)
  if err != nil { return bytea }
  return out
}

type Scorer[F common.FloatType, A common.StringLike, B common.StringLike] struct {
  // Returns a score Between 0 and 1 for the given pair of A and B.
  ScoreFn func(a, b []byte) F

  // Transformer
  Transformer transform.Transformer
}
// Give an array of scores for all the elements in the `accessor` w.r.t. the `target`.
func (sorter Scorer[F, A, B]) ScoreAny(accessor AccessorInterface[A], target B) (out []F) {
  out = make([]F, accessor.Len())
  byteTarget := transformStringLike(sorter.Transformer, target)
  if sorter.ScoreFn == nil {
    sorter.ScoreFn = heuristics.Wrap[F](heuristics.FrequencySimilarity)
    sorter.Transformer = transformers.Lowercase()
  }
  for i := range accessor.Len() { out[i] = sorter.ScoreFn(transformStringLike(sorter.Transformer, accessor.Get(i)), byteTarget) }
  return
}
// Give an array of scores for all the elements in the `array` w.r.t. the `target`.
func (sorter Scorer[F, A, B]) Score(array []A, target B) (out []F) {
  return sorter.ScoreAny(ToSwapperArray(array), target)
}

// A Struct used to sort a collection of elements.
type Sorter[F common.FloatType, A common.StringLike, B common.StringLike] struct {
  Scorer[F, A, B]

  // A value Between 0 and 1 that determines the threshold for the sort.
  // When this is 0, no threshold is applied
  Threshold F
}
type sortAnyType[F common.FloatType, A common.StringLike] struct {
  len     int
  swapper SwapperInterface[A]
  scores  []F
}
func (s *sortAnyType[F, A]) Len() int { return s.len }
func (s *sortAnyType[F, A]) SetLen(i int) { s.len = i }
func (s *sortAnyType[F, A]) Less(i, j int) bool { return s.scores[i] > s.scores[j] }
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
    scores:  sorter.ScoreAny(swapper, target),
  })
}

// Sorts the array in place, returns the number of elements that are in the output
func (sorter Sorter[F, A, B]) Sort(array []A, target B) int {
  return sorter.sort(&sortAnyType[F, A]{
    len:    len(array),
    swapper: ToSwapperArray(array),
    scores: sorter.Score(array, target),
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

  if sorter.Threshold != 0 {
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

