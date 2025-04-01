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
    Company string
    Description string
  }
  target := "apple"
  candidates := []Product{
    {Name: "aple", Company: "Misspelling Corp", Description: "A misspelling of apple"},
    {Name: "appel", Company: "Misspelling Corp", Description: "Another misspelling of apple"},
    {Name: "application", Company: "Light Corp", Description: "A software application"},
    {Name: "orange", Company: "Fruit Corp", Description: "A fruit"},
    {Name: "banana", Company: "Fruit Corp", Description: "A fruit"},
    {Name: "iphone", Company: "Apple", Description: "A smartphone"},
  }

  sorter := Sorter[float32, string, string]{Threshold: 0.8} 

  fmt.Println("Unsorted:", candidates)
  // Note: `SortAny` for single key
  count := sorter.SortAny(ToSwapper(candidates, func(p Product) string { return p.Name }), target)
  fmt.Println("Sorted (and filtered):", candidates[:count]) // output: 

  // Note: `SortAnyArr` for multiple keys (best match is used)
  count = sorter.SortAnyArr(ToSwapper(candidates, func(p Product) []string { return []string{p.Name, p.Company, p.Description} }), target)
  fmt.Println("Sorted (and filtered):", candidates[:count]) // output: 
}

