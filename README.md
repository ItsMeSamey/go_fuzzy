# Go Fuzzy Search

This Go library provides a collection of functions for performing fuzzy string matching and comparison.

## Features

* **String Similarity Algorithms:**
    * Edit Distance Similarity Percentage
    * Longest Common Subsequence (LCS) Percentage
    * Sørensen–Dice Coefficient
    * And many more...
* Support for `golang.org/x/text/transform` with inbuilt transformers for: Lowercasing, ASCII filtering, Unicode normalization.
* Sorting of string collections based on similarity scores, with threshold cut-off.

## Installation

```bash
go get github.com/ItsMeSamey/go_fuzzy
```

##   Usage

### fuzzy.Sorter
* A struct that uses a `Scorer` (see next section) to sort a collection of strings based on their similarity to a target string, with an optional threshold.

Simple Example: Sorting []string

```go
import (
  "fmt"
  "strings"

  "github.com/ItsMeSamey/go_fuzzy"
)

func main() {
  target := "apple"
  candidates := []string{"aple", "application", "orange", "banana", "appel"}

  sorter := fuzzy.Sorter[float32, string, string]{
    // Scorer: Scorer[float64, string, string]{
      // ScoreFn: Defaults to FrequencySimilarity and Lowercase Transformer
      // if no ScoreFn is specified, transformer defaults to transformers.Lowercase()
      // to prevent this, explicitlly specify a ScoreFn (non null), and explicitlly set the transformer to nil
    //}, 
    Threshold: 0.6, // Only include strings with similarity >= 0.6
  }

  fmt.Println("Unsorted:", candidates)
  // Only the first 'count' elements are sorted, rest are still shuffled
  count := sorter.Sort(candidates, target)
  fmt.Println("Sorted (and filtered):", candidates[:count]) // output: ["appel", "aple"]
}
```

Advanced Example: Sorting array of user data types

```go
import (
  "fmt"
  "strings"

  "github.com/ItsMeSamey/go_fuzzy"
)

type Product struct {
  Name string
  Company string
  Description string
}

func main() {
  target := "apple"
  candidates := []Product{
    {Name: "aple", Company: "Misspelling Corp", Description: "A misspelling of apple"},
    {Name: "appel", Company: "Misspelling Corp", Description: "Another misspelling of apple"},
    {Name: "application", Company: "Light Corp", Description: "A software application"},
    {Name: "orange", Company: "Fruit Corp", Description: "A fruit"},
    {Name: "banana", Company: "Fruit Corp", Description: "A fruit"},
    {Name: "iphone", Company: "Apple", Description: "A smartphone"},
  }

  sorter := fuzzy.Sorter[float32, string, string]{Threshold: 0.8} 

  fmt.Println("Unsorted:", candidates)
  // Note: `SortAny` for single key (using only name here)
  count := sorter.SortAny(fuzzy.ToSwapper(candidates, func(p Product) string { return p.Name }), target)
  fmt.Println("Sorted (and filtered):", candidates[:count])
  // output: [{appel Misspelling Corp Another misspelling of apple} {aple Misspelling Corp A misspelling of apple}]

  // Note: `SortAnyArr` for multiple keys (Name, Company and Description are all used)
  count = sorter.SortAnyArr(fuzzy.ToSwapper(candidates, func(p Product) []string { return []string{p.Name, p.Company, p.Description} }), target)
  fmt.Println("Sorted (and filtered):", candidates[:count])
  // output: [{iphone Apple A smartphone} {appel Misspelling Corp Another misspelling of apple} {aple Misspelling Corp A misspelling of apple}]
}
```

### fuzzy.Scorer
* A struct that holds a scoring function (`ScoreFn`) and a `transform.Transformer`.
    The scoring function calculates the similarity score between two strings.
    The transformer is used to transform strings, before calculating the score.

> [!NOTE]
> Although all the provided functions output values from the range `[0, 1]`, you can create your own implementation that goes beyond this range.
> EG: You can use "github.com/ItsMeSamey/go_fuzzy/heuristics/algorithms".LCSLength as a score, and Threshold as cut-off.

Example: Calculating LevenshteinSimilarityPercentage for an array of strings

```go
import (
  "fmt"

  "golang.org/x/text/transform"
  "github.com/ItsMeSamey/go_fuzzy"
  "github.com/ItsMeSamey/go_fuzzy/heuristics"
  "github.com/ItsMeSamey/go_fuzzy/transformers"
)

func main() {
  strs := []string{"hello world", "Hello fuzzy world", "Hello World 2"}
  query := "Hello World"

  scorer := fuzzy.Scorer[float64, string, string]{
    ScoreFn: heuristics.LevenshteinSimilarityPercentage[float64, string, string],
    Transformer: transform.Chain(transformers.UnicodeNormalize(), transformers.Lowercase()), // Should always UnicodeNormalize before Lowercase
  }

  var scores []float64 = scorer.Score(strs, query)
  fmt.Println(scores)
}
```

### common.StringLike
* This interface represents types that can be treated as strings, currently `string` and `[]byte`, so atgument can be any of these types.

Example: Calculating Sørensen–Dice Coefficient

```go
import (
  "fmt"
  "github.com/ItsMeSamey/go_fuzzy/heuristics"
)

func main() {
  string1 := "hello world"
  string2 := "hello world 2"
  byteArray1 := []byte("hello byte world")

  var score32 float32 = heuristics.DiceSorensenCoefficient[float32](string1, string2)
  fmt.Printf("Dice-Sorensen Similarity: %f\n", score32)

  var score64 float64 = heuristics.DiceSorensenCoefficient[float64](byteArray1, string2)
  fmt.Printf("Dice-Sorensen Similarity: %f\n", score64)
}
```

