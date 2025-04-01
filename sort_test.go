package fuzzy

import (
  "fmt"
  "testing"

  "github.com/ItsMeSamey/go_fuzzy/heuristics"
  "github.com/ItsMeSamey/go_fuzzy/transformers"

  "golang.org/x/text/transform"
)

func TestReadmeSort(t *testing.T) {
  target := "apple"
  candidates := []string{"aple", "application", "orange", "banana", "appel"}

  sorter := Sorter[float64, string, string]{
    Threshold: 0.6, // Only include strings with similarity >= 0.6
  }

  fmt.Println("Unsorted:", candidates)
  count := sorter.Sort(candidates, target)
  fmt.Println("Sorted (and filtered):", candidates[:count]) // Only the first 'count' elements are sorted, rest are still shuffled
  fmt.Println("Score: ", sorter.Score(candidates, target))
}

func TestReadmeScorer(t *testing.T) {
  strs := []string{"hello world", "Hello fuzzy world", "Hello World 2"}
  query := "Hello World"

  scorer := Scorer[float64, string, string]{
    ScoreFn: heuristics.LevenshteinSimilarityPercentage[float64, string, string],
    Transformer: transform.Chain(transformers.UnicodeNormalize(), transformers.Lowercase()), // Should always UnicodeNormalize before Lowercase
  }

  var scores []float64 = scorer.Score(strs, query)
  fmt.Println(scores)
}

func TestReadmeStringLike(t *testing.T) {
  string1 := "hello world"
  string2 := "hello world 2"
  byteArray1 := []byte("hello byte world")

  var score32 float32 = heuristics.DiceSorensenCoefficient[float32](string1, string2)
  fmt.Printf("Dice-Sorensen Similarity: %f\n", score32)

  var score64 float64 = heuristics.DiceSorensenCoefficient[float64](byteArray1, string2)
  fmt.Printf("Dice-Sorensen Similarity: %f\n", score64)
}

func TestReadmeSortAny(t *testing.T) {
  type Product struct {
    Name string
    Price float32
  }

  target := "apple"
  candidates := []Product{
    {Name: "aple", Price: 10},
    {Name: "application", Price: 20},
    {Name: "orange", Price: 30},
    {Name: "banana", Price: 40},
    {Name: "appel", Price: 50},
  }

  sorter := Sorter[float32, string, string]{} 

  fmt.Println("Unsorted:", candidates)
  // Only the first 'count' elements are sorted, rest are still shuffled
  count := sorter.SortAny(ToSwapper(candidates, func(p Product) string { return p.Name }), target)
  fmt.Println("Sorted (and filtered):", candidates[:count]) // output: [{appel 50} {aple 10} {application 20} {orange 30} {banana 40}]
}

