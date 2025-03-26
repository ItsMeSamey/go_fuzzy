package common

import (
  "bytes"
  "strings"
)

type Getter[T StringLike] interface {
  Len() int
  Get(idx int) T
}

// Getter form an array of strings
type StringArrayGetter []string
func (g StringArrayGetter) Len() int { return len(g) }
func (g StringArrayGetter) Get(idx int) string { return g[idx] }

// Getter form an array strings, but changes the case to lower case
type NormalizedStringArrayGetter []string
func (g NormalizedStringArrayGetter) Len() int { return len(g) }
func (g NormalizedStringArrayGetter) Get(idx int) string { return strings.ToLower(g[idx]) }

// Getter form an array of bytes (uses unsafe, so be aware of implications)
type ByteArrayGetter [][]byte
func (g ByteArrayGetter) Len() int { return len(g) }
func (g ByteArrayGetter) Get(idx int) []byte { return g[idx] }

// Getter form an array of strings, but changes the case to lower case (uses unsafe, so be aware of implications)
type NormalizedByteArrayGetter [][]byte
func (g NormalizedByteArrayGetter) Len() int { return len(g) }
func (g NormalizedByteArrayGetter) Get(idx int) []byte { return bytes.ToLower(g[idx]) }

