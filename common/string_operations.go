package common


// Len of Intersection of MultiSet of all characters in a and b
func IntersectionCharacterCount[A StringLike, B StringLike](a A, b B) uint {
  var f [256]int

  for i := range len(a) { f[a[i]] += 1 }

  intersection := uint(0)
  for i := range len(b) {
    if f[b[i]] > 0 {
      intersection += 1
      f[b[i]] -= 1
    }
  }

  return intersection
}

// Len of Intersection of Set of all bigrams in a and b
func IntersectionBigramOccurrence[A StringLike, B StringLike](a A, b B) uint {
  // 8 kb only, so not too big of an issue
  var f [256 * (256 >> 6)]uint64

  for i := range len(a) - 1 {
    bigram := uint16(a[i]) << 8 | uint16(a[i + 1])
    f[bigram >> 6] |= 1 << (bigram & 63)
  }
  
  intersection := uint(0)
  for i := range len(b) - 1 {
    bigram := uint16(b[i]) << 8 | uint16(b[i + 1])
    if f[bigram >> 6] & (1 << (bigram & 63)) != 0 { intersection += 1 }
  }

  return intersection
}

