package heuristics

import "fuzzy/common"

func JaccardCoefficientCharacter[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  intersection := common.IntersectionCharacterCount(a, b)
  return F(intersection) / F(uint(len(a)) + uint(len(b)) - intersection)
}

func JaccardCoefficientBigram[F common.FloatType, A common.StringLike, B common.StringLike](a A, b B) F {
  // len(a) + len(b) is ok as the sum cant exceed the native integer type anyways
  // 2 * IntersectionCount(a, b) is ok for the same reason
  intersection := common.IntersectionBigramOccurrence(a, b)
  return F(intersection) / F(uint(len(a)) + uint(len(b)) - intersection)
}

