package transformers

import (
  "golang.org/x/text/transform"
  "golang.org/x/text/unicode/norm"
  "golang.org/x/text/encoding/charmap"
)

// Normalize Unicode to NFKD (decomposed form).
func UnicodeNormalize() transform.Transformer {
  return norm.NFKD
}

// Convert Unicode to ASCII, removing any characters that cannot be represented in ASCII.
func UnicodeToAscii() transform.Transformer {
  return transform.Chain(charmap.ISO8859_1.NewEncoder(), transform.RemoveFunc(func(r rune) bool { return r > 127 }))
}

