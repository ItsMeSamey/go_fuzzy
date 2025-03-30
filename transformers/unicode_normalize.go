package transformers

import (
  "golang.org/x/text/runes"
  "golang.org/x/text/transform"
  "golang.org/x/text/unicode/norm"
)

// Normalize Unicode to NFKD (decomposed form).
func UnicodeNormalize() transform.Transformer {
  return norm.NFKD
}

type asciiFilter struct {}
func (asciiFilter) Contains(r rune) bool { return r >= 128 }
// Removes any characters that cannot be represented in ASCII.
func AsciiFilter() transform.Transformer {
  return runes.Remove(asciiFilter{})
}

