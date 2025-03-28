package heuristics

import "fuzzy/common"

func DiceSorensenCoefficientCharacter[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  return F(2 * common.IntersectionCharacterCount(a, b)) / F(uint(len(a)) + uint(len(b)))
}

func DiceSorensenCoefficientBigram[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  return F(2 * common.IntersectionBigramOccurrence(a, b)) / F(uint(len(a)) + uint(len(b)))
}

