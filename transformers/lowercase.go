package transformers

import "golang.org/x/text/transform"

type lowercaseTransformer struct{}

func (lowercaseTransformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
  nSrc = min(len(src), len(dst))
  nDst = nSrc

  for i := range nSrc {
    if 'A' <= src[i] && src[i] <= 'Z' {
      dst[i] = src[i] + 'a' - 'A'
    } else {
      dst[i] = src[i]
    }
  }

  if len(src) > len(dst) {
    err = transform.ErrShortDst
  } else if len(src) < len(dst) {
    err = transform.ErrShortSrc
  }

  return
}

func (lowercaseTransformer) Reset() {}

func Lowercase() transform.Transformer {
  return lowercaseTransformer{}
}

